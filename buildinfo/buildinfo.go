package buildinfo

import (
	"bytes"
	"embed"
	"fmt"
	"runtime"
	"runtime/debug"
)

var version string

//go:generate bash -e gen.sh
//go:embed *.txt
var files embed.FS

func init() {
	var info struct {
		desc  string
		rev   string
		wip   bool
		build *debug.BuildInfo
	}

	if describe, err := files.ReadFile("describe.txt"); err == nil {
		info.desc = string(bytes.TrimSpace(describe))
	}
	if revision, err := files.ReadFile("revision.txt"); err == nil {
		info.rev = string(bytes.TrimSpace(revision))
	}
	if status, err := files.ReadFile("status.txt"); err == nil {
		info.wip = len(status) != 0
	}
	if build, ok := debug.ReadBuildInfo(); ok {
		info.build = build
	}

	version = compose(info.desc, info.rev, info.wip, info.build)
}

func compose(desc, rev string, wip bool, build *debug.BuildInfo) string {
	if build == nil {
		if desc != "" {
			return fmt.Sprintf("%s %s/%s", desc, runtime.GOOS, runtime.GOARCH)
		}
		if rev != "" {
			if wip {
				rev += "-wip"
			}
			return fmt.Sprintf("%s %s/%s", rev, runtime.GOOS, runtime.GOARCH)
		}
		return ""
	}

	var vcs struct {
		revision string
		modified bool
	}

	for _, item := range build.Settings {
		switch item.Key {
		case "vcs.revision":
			vcs.revision = item.Value
		case "vcs.modified":
			vcs.modified = item.Value == "true"
		}
	}

	if rev != vcs.revision || wip != vcs.modified || desc == "" {
		if vcs.revision == "" {
			return ""
		}
		if vcs.modified {
			desc = vcs.revision + "-wip"
		} else {
			desc = vcs.revision
		}
	}

	return fmt.Sprintf("%s %s %s/%s", desc, build.GoVersion, runtime.GOOS, runtime.GOARCH)
}

func Version() string {
	return version
}
