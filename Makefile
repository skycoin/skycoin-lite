.DEFAULT_GOAL := help

.PHONY: build-js build-js-min test lint check install-linters format fix-skycoin-dep help
.PHONY: test-js

build-js: ## Build /skycoinjs/skycoinjs.go. The result is saved in the repo root
	go build -o gopherjs-tool vendor/github.com/gopherjs/gopherjs/tool.go
	GOOS=linux ./gopherjs-tool build skycoinjs/skycoinjs.go

build-js-min: ## Build /skycoinjs/skycoinjs.go. The result is minified and saved in the repo root
	go build -o gopherjs-tool vendor/github.com/gopherjs/gopherjs/tool.go
	GOOS=linux ./gopherjs-tool build skycoinjs/skycoinjs.go -m

test-js:
	go build -o gopherjs-tool vendor/github.com/gopherjs/gopherjs/tool.go
	./gopherjs-tool test ./skycoinjs/

test:
	go test ./... -timeout=10m -cover

lint: ## Run linters. Use make install-linters first.
	vendorcheck ./...
	gometalinter --deadline=3m -j 2 --disable-all --tests --exclude .. --vendor \
		-E goimports \
		-E unparam \
		-E deadcode \
		-E errcheck \
		-E gas \
		-E goconst \
		-E gofmt \
		-E golint \
		-E ineffassign \
		-E maligned \
		-E megacheck \
		-E misspell \
		-E nakedret \
		-E structcheck \
		-E unconvert \
		-E varcheck \
		-E vet \
		./...

check: lint test ## Run tests and linters

install-linters: ## Install linters
	go get -u github.com/FiloSottile/vendorcheck
	go get -u github.com/alecthomas/gometalinter
	gometalinter --vendored-linters --install

format: ## Formats the code. Must have goimports installed (use make install-linters).
	goimports -w ./gopher
	goimports -w ./liteclient
	goimports -w ./mobile

fix-skycoin-dependency: ## Modify the Skycoin code inside vendor, so that gopherjs can transpile correctly (see readme.md for more info).
	fix-skycoin-dependency.sh

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
