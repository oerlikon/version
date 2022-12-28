package buildinfo_test

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
	"testing"

	"version/buildinfo"
)

func TestCompose(t *testing.T) {
	var target = fmt.Sprintf(" %s/%s", runtime.GOOS, runtime.GOARCH)

	var tests = []struct {
		desc     string
		rev      string
		wip      bool
		build    *debug.BuildInfo
		expected string
	}{
		{ // 1
			"",
			"",
			false,
			nil,
			"",
		},
		{ // 2
			"",
			"407c1bad388caa515ba4727588711e508b38af3d",
			false,
			nil,
			"g407c1bad388c" + target,
		},
		{ // 3
			"",
			"407c1bad388caa515ba4727588711e508b38af3d",
			true,
			nil,
			"g407c1bad388c-wip" + target,
		},
		{ // 4
			"v0.0.0",
			"407c1bad388caa515ba4727588711e508b38af3d",
			false,
			nil,
			"v0.0.0" + target,
		},
		{ // 5
			"v0.0.0",
			"407c1bad388caa515ba4727588711e508b38af3d",
			true,
			nil,
			"v0.0.0-wip" + target,
		},
		{ // 6
			"",
			"",
			false,
			&debug.BuildInfo{
				GoVersion: "go1.19.4",
				Settings: []debug.BuildSetting{
					{
						Key:   "vcs.revision",
						Value: "6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
					},
					{
						Key:   "vcs.modified",
						Value: "false",
					},
				},
			},
			"g6e31ad65dcfc go1.19.4" + target,
		},
		{ // 7
			"",
			"",
			false,
			&debug.BuildInfo{
				GoVersion: "go1.19.4",
				Settings: []debug.BuildSetting{
					{
						Key:   "vcs.revision",
						Value: "6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
					},
					{
						Key:   "vcs.modified",
						Value: "true",
					},
				},
			},
			"g6e31ad65dcfc-wip go1.19.4" + target,
		},
		{ // 8
			"v0.0.0",
			"407c1bad388caa515ba4727588711e508b38af3d",
			true,
			&debug.BuildInfo{
				GoVersion: "go1.19.4",
				Settings: []debug.BuildSetting{
					{
						Key:   "vcs.revision",
						Value: "6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
					},
					{
						Key:   "vcs.modified",
						Value: "false",
					},
				},
			},
			"g6e31ad65dcfc go1.19.4" + target,
		},
		{ // 9
			"v0.0.0",
			"407c1bad388caa515ba4727588711e508b38af3d",
			false,
			&debug.BuildInfo{
				GoVersion: "go1.19.4",
				Settings: []debug.BuildSetting{
					{
						Key:   "vcs.revision",
						Value: "6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
					},
					{
						Key:   "vcs.modified",
						Value: "true",
					},
				},
			},
			"g6e31ad65dcfc-wip go1.19.4" + target,
		},
		{ // 10
			"v0.0.0",
			"6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
			false,
			&debug.BuildInfo{
				GoVersion: "go1.19.4",
				Settings: []debug.BuildSetting{
					{
						Key:   "vcs.revision",
						Value: "6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
					},
					{
						Key:   "vcs.modified",
						Value: "true",
					},
				},
			},
			"g6e31ad65dcfc-wip go1.19.4" + target,
		},
		{ // 11
			"",
			"6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
			false,
			&debug.BuildInfo{
				GoVersion: "go1.19.4",
				Settings: []debug.BuildSetting{
					{
						Key:   "vcs.revision",
						Value: "6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
					},
					{
						Key:   "vcs.modified",
						Value: "false",
					},
				},
			},
			"g6e31ad65dcfc go1.19.4" + target,
		},
		{ // 12
			"",
			"6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
			true,
			&debug.BuildInfo{
				GoVersion: "go1.19.4",
				Settings: []debug.BuildSetting{
					{
						Key:   "vcs.revision",
						Value: "6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
					},
					{
						Key:   "vcs.modified",
						Value: "true",
					},
				},
			},
			"g6e31ad65dcfc-wip go1.19.4" + target,
		},
		{ // 13
			"v0.0.0",
			"6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
			false,
			&debug.BuildInfo{
				GoVersion: "go1.19.4",
				Settings: []debug.BuildSetting{
					{
						Key:   "vcs.revision",
						Value: "6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
					},
					{
						Key:   "vcs.modified",
						Value: "false",
					},
				},
			},
			"v0.0.0 go1.19.4" + target,
		},
		{ // 14
			"v0.0.0",
			"6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
			true,
			&debug.BuildInfo{
				GoVersion: "go1.19.4",
				Settings: []debug.BuildSetting{
					{
						Key:   "vcs.revision",
						Value: "6e31ad65dcfcb9f0908cea7d17e17b5d78d1f0a8",
					},
					{
						Key:   "vcs.modified",
						Value: "true",
					},
				},
			},
			"v0.0.0-wip go1.19.4" + target,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			v := buildinfo.Compose(test.desc, test.rev, test.wip, test.build)
			if v != test.expected {
				t.Errorf("%s != %s", v, test.expected)
			}
		})
	}
}
