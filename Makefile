watch:
	go tool air -c .air.toml

dev:
	go run ./cmd/workflou

build: templ
	go build -o bin/workflou ./cmd/workflou

.PHONY: test
test:
	go test ./...

templ:
	go tool templ generate