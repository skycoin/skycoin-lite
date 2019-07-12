#!/usr/bin/env bash

set -e -o pipefail

if [ -z "$VERSION" ]; then
	echo "VERSION must be set"
	exit 1
fi

echo "installing golangci-lint v$VERSION to $(go env GOPATH)"

# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v$VERSION

golangci-lint --version
