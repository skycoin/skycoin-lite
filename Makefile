.DEFAULT_GOAL := help
.PHONY: build-js build-android

build-js: ## Build js library with gopherjs
	@mkdir -p build/js
	gopherjs build -o build/js/skycoin-lite.js

build-android: ## Build android aar archive with gomobile
	@mkdir -p build/android
	gomobile bind -target=android github.com/skycoin/skycoin-lite/mobile

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

