git describe --tags > describe.txt || true
git rev-parse HEAD > revision.txt
git status --porcelain > status.txt
