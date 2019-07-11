#!/usr/bin/env bash

set -e -o pipefail

if [ -z "$VERSION" ]; then
	echo "VERSION must be set"
	exit 1
fi

which go

go env

GOPATH=$(go env GOPATH)

echo "installing golangci-lint v$VERSION to $GOPATH"

# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v$VERSION

golangci-lint --version
