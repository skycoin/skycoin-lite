package main

import (
	"encoding/hex"
	"testing"

	"github.com/skycoin/skycoin-lite/liteclient"
)

func TestGenerateAddress(t *testing.T) {
	stringSeed := "abcdefg"
	seed := hex.EncodeToString([]byte(stringSeed))

	addr, err := liteclient.GenerateAddress(seed)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if addr.Address != "gyJNKKj95bCn6o5mUCQgz8SCem7av2W3CG" {
		t.Fatalf("GenerateAddress address is invalid")
	}

	if addr.Public != "02e1f33d00576ef4d89adcd1c8cb732810de862e8c5c6ecdc0208ccc6faefd1e09" {
		t.Fatalf("GenerateAddress pubkey is invalid")
	}

	nextAddr, err := liteclient.GenerateAddress(addr.NextSeed)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if nextAddr.Address != "2YzQDZFpS64u5ydnE9iKK9WJPXU2EbmSzmW" {
		t.Fatalf("GenerateAddress address is invalid 2")
	}

	if nextAddr.Public != "023ea091a649aef9d8fe5b956bd5cf2a9e47b8a7ba41f8b83997de0d1b1a851e73" {
		t.Fatalf("GenerateAddress pubkey is invalid 2")
	}
}
