git describe --tags --dirty="-wip" > describe.txt || true
git rev-parse HEAD > revision.txt
git status --porcelain > status.txt
