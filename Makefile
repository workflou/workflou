watch:
	go tool air -build.cmd="make build" \
		-build.bin="bin/workflou" \
		-build.include_ext="go,js,css,html" \
		-build.exclude_regex="min.js|_test.go" \
		-build.exclude_dir="" \
		-tmp_dir="bin"

dev:
	go run ./cmd/workflou

build:
	go build -o bin/workflou ./cmd/workflou