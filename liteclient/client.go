package liteclient

import (
	"encoding/hex"
	"encoding/json"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
)

type Address struct {
	NextSeed string
	Secret   string
	Public   string
	Address  string
}

type TransactionInput struct {
	Hash   string
	Secret string
}

type TransactionOutput struct {
	Address string
	Coins   uint64
	Hours   uint64
}

// Receives a hex encrypted string
func GenerateAddress(seed string) Address {
	decodedString, _ := hex.DecodeString(seed)
	next, keys := cipher.GenerateDeterministicKeyPairsSeed([]byte(decodedString), 1)
	pub := cipher.PubKeyFromSecKey(keys[0])
	address := Address{
		hex.EncodeToString(next),
		keys[0].Hex(),
		pub.Hex(),
		cipher.AddressFromPubKey(pub).String(),
	}

	return address
}

// Receives inputs and outputs and returns a signed transaction
func PrepareTransaction(inputsBody string, outputsBody string) string {
	var inputs []TransactionInput
	var outputs []TransactionOutput

	json.Unmarshal([]byte(inputsBody), &inputs)
	json.Unmarshal([]byte(outputsBody), &outputs)

	newTransaction := coin.Transaction{}

	keys := make([]cipher.SecKey, len(inputs))

	for i, in := range inputs {
		k, _ := cipher.SecKeyFromHex(in.Secret)
		keys[i] = k
		newTransaction.PushInput(cipher.MustSHA256FromHex(in.Hash))
	}

	for _, out := range outputs {
		newTransaction.PushOutput(cipher.MustDecodeBase58Address(out.Address), out.Coins, out.Hours)
	}

	newTransaction.SignInputs(keys)
	newTransaction.UpdateHeader()
	d := newTransaction.Serialize()

	return hex.EncodeToString(d)
}

// Currently not in use
func Addresses(seed string, amount int) ([]Address, error) {
	//For this to return the same addresses as the wallet, this must use "[]byte(hex.DecodeString(seed))" instead of "[]byte(seed)".
	_, secretKeys := cipher.GenerateDeterministicKeyPairsSeed([]byte(seed), amount)
	addresses := make([]Address, amount)
	for i, sec := range secretKeys {
		pub := cipher.PubKeyFromSecKey(sec)
		address := Address{
			"",
			cipher.AddressFromPubKey(pub).String(),
			pub.Hex(),
			sec.Hex(),
		}
		addresses[i] = address
	}

	return addresses, nil
}
