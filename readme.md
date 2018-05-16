# Skycoin Liteclient

This repository contains a liteclient for Skycoin written in Go. At the moment it is only used to compile an [Android Archive](https://developer.android.com/studio/projects/android-library.html).

## Compiling

For the compilation process to Android Archive, we use [Go Mobile](https://github.com/golang/mobile).

```bash
$ gomobile bind -target=android github.com/montycrypto/skycoin-lite/mobile/mobile
```
