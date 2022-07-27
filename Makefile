NAME=testkube-watch
VERSION=v1.0.5

.PHONY: run test clean

build:
	@goreleaser build --single-target --snapshot --rm-dist

run:
	@TKW_HOME=$(PWD)/examples go run main.go

test:
	@go test -v

clean:
	@go clean
	@rm -f $(NAME)
	@rm -rf dist/

docker:
	@docker build -t lreimer/$(NAME)-controller --build-arg version=$(VERSION) .

release:
	@goreleaser --rm-dist
