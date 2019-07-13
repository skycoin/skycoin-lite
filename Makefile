.DEFAULT_GOAL := help
.PHONY: build-js build-js-min
.PHONY: test lint check
.PHONY: install-linters format
.PHONY: help
.PHONY: test-js
.PHONY: test-suite-ts test-suite-ts-extensive

build-js: ## Build /skycoin/skycoin.go. The result is saved in the repo root
	go build -o gopherjs-tool vendor/github.com/gopherjs/gopherjs/tool.go
	GOOS=linux ./gopherjs-tool build skycoin/skycoin.go -o js/skycoin.js

build-js-min: ## Build /skycoin/skycoin.go. The result is minified and saved in the repo root
	go build -o gopherjs-tool vendor/github.com/gopherjs/gopherjs/tool.go
	GOOS=linux ./gopherjs-tool build skycoin/skycoin.go -m -o js/skycoin.js

build-wasm: ## Build /wasm/skycoin.go. The result is saved in the repo root as skycoin-lite.wasm
	GOOS=js GOARCH=wasm go build -o js/skycoin-lite.wasm ./wasm/skycoin.go

test-js: ## Run the Go tests using JavaScript
	go build -o gopherjs-tool vendor/github.com/gopherjs/gopherjs/tool.go
	./gopherjs-tool test ./skycoin/ -v

test-suite-ts: ## Run the ts version of the cipher test suite for GopherJS. Use a small number of test cases
	cd js && npm run test

test-suite-ts-extensive: ## Run the ts version of the cipher test suite for GopherJS. All the test cases
	cd js && npm run test-extensive

test-suite-ts-wasm: ## Run the ts version of the cipher test suite for wasm and additional tests
	cd vendor/github.com/skycoin/skycoin/src/cipher/secp256k1-go && GOOS=js GOARCH=wasm go test -c -o test.wasm
	cd vendor/github.com/skycoin/skycoin/src/cipher/secp256k1-go/secp256k1-go2 && GOOS=js GOARCH=wasm go test -c -o test.wasm
	cd js && npm run test-wasm
	cd js/tests && node cipher-wasm-internal.js
	cd vendor/github.com/skycoin/skycoin/src/cipher/secp256k1-go && rm test.wasm
	cd vendor/github.com/skycoin/skycoin/src/cipher/secp256k1-go/secp256k1-go2 && rm test.wasm

test:
	go test ./... -timeout=10m -cover

lint: ## Run linters. Use make install-linters first.
	vendorcheck ./...
	golangci-lint run -c ./.golangci.yml ./...
	@# The govet version in golangci-lint is out of date and has spurious warnings, run it separately
	go vet -all ./...

check: lint test ## Run tests and linters

install-linters: ## Install linters
	go get -u github.com/FiloSottile/vendorcheck
	# For some reason this install method is not recommended, see https://github.com/golangci/golangci-lint#install
	# However, they suggest `curl ... | bash` which we should not do
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

format: ## Formats the code. Must have goimports installed (use make install-linters).
	goimports -w -local github.com/skycoin/skycoin-lite ./skycoin
	goimports -w -local github.com/skycoin/skycoin-lite ./liteclient
	goimports -w -local github.com/skycoin/skycoin-lite ./mobile
	goimports -w -local github.com/skycoin/skycoin-lite ./main.go

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
