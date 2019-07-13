package main

import (
	"errors"
	"syscall/js"

	"github.com/skycoin/skycoin-lite/liteclient"
)

// recoverFromPanic captures the panics and returns an object with the error message.
// It must be used in all the functions that can be called using the compiled wasm
// file, as the Go code contains multiple panics that would completelly stop the
// excecution of the wasm application without returning adequate errors to the JS code.
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

// checkParams checks if all the params are of he type js.TypeString.
func checkParams(params *[]js.Value) {
	for _, element := range *params {
		if element.Type() != js.TypeString {
			panic(errors.New("Invalid argument type"))
		}
	}
}

// Main functions:
// The following functions are simply wrappers to call the functions in
// liteclient/client.go.

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

// Extra functions:
// The following functions are simply wrappers to call the functions in
// liteclient/extras.go.

func verifyPubKeySignedHash(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	liteclient.VerifyPubKeySignedHash(inputs[0].String(), inputs[1].String(), inputs[2].String())

	return
}

func verifyAddressSignedHash(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	liteclient.VerifyAddressSignedHash(inputs[0].String(), inputs[1].String(), inputs[2].String())

	return
}

func verifySignatureRecoverPubKey(this js.Value, inputs []js.Value) (response interface{}) {
	defer recoverFromPanic(&response)
	checkParams(&inputs)

	liteclient.VerifySignatureRecoverPubKey(inputs[0].String(), inputs[1].String())

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
	// Create a channel for keeping the application alive
	c := make(chan bool)

	// Add the main functions to the the "window.SkycoinCipher" object.
	cipherNamespace := "SkycoinCipher"
	js.Global().Set(cipherNamespace, js.FuncOf(nil))
	js.Global().Get(cipherNamespace).Set("generateAddress", js.FuncOf(generateAddress))
	js.Global().Get(cipherNamespace).Set("prepareTransaction", js.FuncOf(prepareTransaction))
	js.Global().Get(cipherNamespace).Set("prepareTransactionWithSignatures", js.FuncOf(prepareTransactionWithSignatures))

	// Add the extra functions to the the "window.SkycoinCipherExtras" object.
	cipherExtrasNamespace := "SkycoinCipherExtras"
	js.Global().Set(cipherExtrasNamespace, js.FuncOf(nil))
	js.Global().Get(cipherExtrasNamespace).Set("verifyPubKeySignedHash", js.FuncOf(verifyPubKeySignedHash))
	js.Global().Get(cipherExtrasNamespace).Set("verifyAddressSignedHash", js.FuncOf(verifyAddressSignedHash))
	js.Global().Get(cipherExtrasNamespace).Set("verifySignatureRecoverPubKey", js.FuncOf(verifySignatureRecoverPubKey))
	js.Global().Get(cipherExtrasNamespace).Set("verifySeckey", js.FuncOf(verifySeckey))
	js.Global().Get(cipherExtrasNamespace).Set("verifyPubkey", js.FuncOf(verifyPubkey))
	js.Global().Get(cipherExtrasNamespace).Set("addressFromPubKey", js.FuncOf(addressFromPubKey))
	js.Global().Get(cipherExtrasNamespace).Set("addressFromSecKey", js.FuncOf(addressFromSecKey))
	js.Global().Get(cipherExtrasNamespace).Set("pubKeyFromSig", js.FuncOf(pubKeyFromSig))
	js.Global().Get(cipherExtrasNamespace).Set("signHash", js.FuncOf(signHash))
	<-c
}
