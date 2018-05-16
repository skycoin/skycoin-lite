# Skycoin Liteclient

This repository contains a liteclient for Skycoin written in Go. At the moment it is only used to compile an [Android Archive](https://developer.android.com/studio/projects/android-library.html) and a JS library with [gopherjs](https://github.com/gopherjs/gopherjs).

## Compiling Android aar and jar

For the compilation process to Android Archive, we use [Go Mobile](https://github.com/golang/mobile).

```bash
$ gomobile bind -target=android github.com/skycoin/skycoin-lite/mobile
```

## Compile javascript library

For the compilation process to javascript library, we use [gopherjs](https://github.com/gopherjs/gopherjs).

Follow gopherjs's install instructions, then do:

```bash
$ gopherjs build
```
