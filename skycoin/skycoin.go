package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/skycoin/skycoin-lite/liteclient"
)

func main() {
	js.Global.Set("Cipher", map[string]interface{}{
		"GenerateAddresses":  liteclient.GenerateAddress,
		"PrepareTransaction": liteclient.PrepareTransaction,
	})

	js.Global.Set("CipherExtras", map[string]interface{}{
		"VerifySignature":  liteclient.VerifySignature,
		"ChkSig": liteclient.ChkSig,
		"VerifySignedHash": liteclient.VerifySignedHash,
		"VerifySeckey": liteclient.VerifySeckey,
		"VerifyPubkey": liteclient.VerifyPubkey,
		"AddressFromPubKey": liteclient.AddressFromPubKey,
		"AddressFromSecKey": liteclient.AddressFromSecKey,
		"PubKeyFromSig": liteclient.PubKeyFromSig,
		"SignHash": liteclient.SignHash,
	})
}
