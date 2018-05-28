package main

import (
	"encoding/hex"
	"fmt"

	"github.com/skycoin/skycoin-lite/liteclient"
)

// For testing purposes. This file is not part of the library.
func main() {
	seed := hex.EncodeToString([]byte("nest*"))

	addrs, err := liteclient.GenerateAddresses(seed, 3)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(addrs)
	}

	address1, err := liteclient.GenerateAddress(seed)
	fmt.Println("----")
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(address1)
	}

	address2, err := liteclient.GenerateAddress(address1.NextSeed)
	fmt.Println("----")
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(address2)
	}

	address3, err := liteclient.GenerateAddress(address2.NextSeed)
	fmt.Println("----")
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(address3)
	}
}
