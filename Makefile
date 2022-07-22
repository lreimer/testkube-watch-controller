NAME=testkube-watch
VERSION=v1.0.0

.PHONY: build run test clean

build:
	# omit the symbol table, debug information and the DWARF table
	@go build -o $(NAME) -ldflags="-s -w -X main.version=$(VERSION)"

run:
	@go run main.go

test:
	@go test -v

clean:
	@go clean
	@rm -f $(NAME)

releaser:
	@goreleaser build --snapshot --skip-publish --rm-dist
