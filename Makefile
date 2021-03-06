.PHONY: install
VERSION ?= local-dev

build:
	go build -ldflags="-s -w -X 'github.com/craftcms/nitro/internal/cmd.Version=${VERSION}'" -o nitro ./cmd/cli
local: build
	mv nitro /usr/local/bin/nitro
test:
	go test ./...
releaser:
	goreleaser --skip-publish --rm-dist --skip-validate
