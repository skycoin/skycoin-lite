package main

import (
	"errors"
	"syscall/js"

	"github.com/skycoin/skycoin-lite/liteclient"
)

func recoverFromPanic(response *interface{}) {
	if err := recover(); err != nil {
		finalResponse := make(map[string]interface{})

		if r, ok := err.(error); ok {
			finalResponse["error"] = r.Error()
		} else if r, ok := err.(string); ok {
			finalResponse["error"] = r
		} else {
			finalResponse["error"] = "Error performing cryptographic operation"
		}

		*response = finalResponse
	}
}

func checkParams(params *[]js.Value) {
	for _, element := range *params {
		if element.Type() != js.TypeString {
			panic(errors.New("Invalid argument type"))
		}
	}
}

// Main functions

func generateAddress(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	functionResponse := liteclient.GenerateAddress(inputs[0].String())

	finalResponse := make(map[string]interface{})
	finalResponse["address"] = functionResponse.Address
	finalResponse["nextSeed"] = functionResponse.NextSeed
	finalResponse["public"] = functionResponse.Public
	finalResponse["secret"] = functionResponse.Secret

	return finalResponse
}

func prepareTransaction(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	functionResponse := liteclient.PrepareTransaction(inputs[0].String(), inputs[1].String())

	return functionResponse
}

func prepareTransactionWithSignatures(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	functionResponse := liteclient.PrepareTransactionWithSignatures(inputs[0].String(), inputs[1].String(), inputs[2].String())

	return functionResponse
}

// Extra functions

func verifySignature(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	liteclient.VerifySignature(inputs[0].String(), inputs[1].String(), inputs[2].String())

	return
}

func chkSig(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	liteclient.ChkSig(inputs[0].String(), inputs[1].String(), inputs[2].String())

	return
}

func verifySignedHash(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	liteclient.VerifySignedHash(inputs[0].String(), inputs[1].String())

	return
}

func verifySeckey(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	liteclient.VerifySeckey(inputs[0].String())

	return
}

func verifyPubkey(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	liteclient.VerifyPubkey(inputs[0].String())

	return
}

func addressFromPubKey(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	functionResponse := liteclient.AddressFromPubKey(inputs[0].String())

	return functionResponse
}

func addressFromSecKey(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	functionResponse := liteclient.AddressFromSecKey(inputs[0].String())

	return functionResponse
}

func pubKeyFromSig(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	functionResponse := liteclient.PubKeyFromSig(inputs[0].String(), inputs[1].String())

	return functionResponse
}

func signHash(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	functionResponse := liteclient.SignHash(inputs[0].String(), inputs[1].String())

	return functionResponse
}

func main() {
	c := make(chan bool)
	cipherNamespace := "SkycoinCipher"
	js.Global().Set(cipherNamespace, js.FuncOf(nil))
	js.Global().Get(cipherNamespace).Set("generateAddress", js.FuncOf(generateAddress))
	js.Global().Get(cipherNamespace).Set("prepareTransaction", js.FuncOf(prepareTransaction))
	js.Global().Get(cipherNamespace).Set("prepareTransactionWithSignatures", js.FuncOf(prepareTransactionWithSignatures))

	cipherExtrasNamespace := "SkycoinCipherExtras"
	js.Global().Set(cipherExtrasNamespace, js.FuncOf(nil))
	js.Global().Get(cipherExtrasNamespace).Set("verifySignature", js.FuncOf(verifySignature))
	js.Global().Get(cipherExtrasNamespace).Set("chkSig", js.FuncOf(chkSig))
	js.Global().Get(cipherExtrasNamespace).Set("verifySignedHash", js.FuncOf(verifySignedHash))
	js.Global().Get(cipherExtrasNamespace).Set("verifySeckey", js.FuncOf(verifySeckey))
	js.Global().Get(cipherExtrasNamespace).Set("verifyPubkey", js.FuncOf(verifyPubkey))
	js.Global().Get(cipherExtrasNamespace).Set("addressFromPubKey", js.FuncOf(addressFromPubKey))
	js.Global().Get(cipherExtrasNamespace).Set("addressFromSecKey", js.FuncOf(addressFromSecKey))
	js.Global().Get(cipherExtrasNamespace).Set("pubKeyFromSig", js.FuncOf(pubKeyFromSig))
	js.Global().Get(cipherExtrasNamespace).Set("signHash", js.FuncOf(signHash))
	<-c
}
