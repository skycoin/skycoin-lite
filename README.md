[![Build Status](https://travis-ci.com/skycoin/skycoin-lite.svg?branch=master)](https://travis-ci.com/skycoin/skycoin-lite)

# Skycoin Liteclient

This repository contains a small wrapper for Skycoin (written in Go) to provide mobile, JS and wasm bindings.

At the moment it is used to compile
an [Android Archive](https://developer.android.com/studio/projects/android-library.html), an iOS Framework,
a (deprecated) JS library with [gopherjs](https://github.com/gopherjs/gopherjs) and a wasm file.

Supports go1.10+. However, for compiling the wasm file you must use Go v1.12.x (compatibility
with Go v1.13+ is not guaranteed).

## Important note about error handling

Many functions on this library call the `panic()` function of the Go programming language in case of important
errors, however, for various reasons the panics are converted into errors on each supported language. Due to
this, it is important to bear in mind that errors returned by this library may be due to extremely important
problems and continuing the execution of operations with results obtained from the library could cause loss
of coins.

It is recommended to be very careful with errors.

## Compiling Android aar and jar, and iOS Framework

For the compilation process to Android Archive and iOS Framework, we use [Go Mobile](https://github.com/golang/mobile).

```bash
$ gomobile bind -target=android github.com/skycoin/skycoin-lite/mobile
$ gomobile bind -target=ios github.com/skycoin/skycoin-lite/mobile
```

## Compile javascript library

> IMPORTANT: the ability for transpiling the code to a JavaScript library is deprecated due to important
performance issues and should not be used. Please compile to a wasm file instead.

For the compilation process to javascript library, we use [gopherjs](https://github.com/gopherjs/gopherjs).

To compile the library use `make build-js` or `make build-js-min` (if you want the final file to be minified).
After compiling, `the main.js` and `main.js.map` files will be created/updated in the root of the repository.

## Compile wasm file

> IMPORTANT: you need Go v1.12.x to use this function. It is not guaranteed to work with Go v1.13+.

To compile the wasm file use `make build-wasm`. After compiling, the `skycoin-lite.wasm` file will be
created/updated in the root of the repository.

## Development

The javascript library is created starting from [skycoin/skycoin.go](skycoin/skycoin.go). The wasm file is
created starting from [wasm/skycoin.go](wasm/skycoin.go). The Android/iOS library is created starting
from [mobile/api.go](mobile/api.go).

### Running tests

#### Gopherjs tests

gopherjs tests can be run with

```sh
make test-js
```

The tests require node syscall support installed, see install instructions at
https://github.com/gopherjs/gopherjs/blob/master/doc/syscalls.md#nodejs-on-linux-and-macos

Note that you can't use the vendored gopherjs for this, because the gopherjs/node-syscall package
can't be vendored by dep. You'll have to install gopherjs to your `GOPATH` with `go get`.

To enable stacktraces, install source maps:

```sh
cd js
npm install --global source-map-support
```

and make sure `NODE_PATH` is set to the value of `npm root --global` in your environment.

#### TS cipher test suite for GopherJS

> IMPORTANT: the ability for transpiling the code to a JavaScript library is deprecated due to important
performance issues and should not be used. Please compile to a wasm file instead.

The repository includes a TypeScript version of the cipher test suite, originally written in Go in
the Skycoin main repository. It test the GopherJS version of the library, so before
using it you must compile the GopherJS version with `make build-js` or `make build-js-min`.
Because testing the GopherJS version takes a significant amount of time to complete, the test suite
can be run with a limited number of cases with:

```sh
make test-suite-ts
```

The test suite can be run with all test cases using:

```sh
make test-suite-ts-extensive
```

#### TS cipher test suite for wasm

There is test suite for the wasm version of the library, just like there is one for the GopherJS
version. Before using it you must compile the wasm version with `make build-wasm`. The test suite
can be run with:

```sh
make test-suite-ts-wasm
```

### Formatting

All `.go` source files should be formatted `goimports`.  You can do this with:

```sh
make format
```

### Code Linting

Install prerequisites:

```sh
make install-linters
```

Run linters:

```sh
make lint
```
