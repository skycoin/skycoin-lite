package main

import (
	"fmt"
	"github.com/skycoin/skycoin-lite/liteclient"
	"encoding/hex"
)

func main() {
	seed := hex.EncodeToString([]byte("nest*"))
	fmt.Println(liteclient.Addresses(seed, 3))

	address1 := liteclient.GenerateAddress(seed)
	fmt.Println("----")
	fmt.Println(address1)

	address2 := liteclient.GenerateAddress(address1.NextSeed)
	fmt.Println("----")
	fmt.Println(address2)

	address3 := liteclient.GenerateAddress(address2.NextSeed)
	fmt.Println("----")
	fmt.Println(address3)
}
