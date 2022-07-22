NAME=testkube-watch
VERSION=v1.0.0

.PHONY: build run test clean

build:
	@go build -o $(NAME) -ldflags="-s -w -X main.version=$(VERSION)"

run:
	@TKW_HOME=$(PWD)/examples go run main.go

test:
	@go test -v

clean:
	@go clean
	@rm -f $(NAME)

docker: build
	@docker build -t lreimer/$(NAME)-controller .

release:
	@goreleaser --rm-dist
