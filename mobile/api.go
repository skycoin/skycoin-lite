package mobile

import (
	"encoding/json"
	"encoding/hex"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin-lite/liteclient"
	"github.com/skycoin/skycoin/src/cipher/go-bip39"
)

// GetAddresses returns a series of addresses based on a seed and the number of addresses
// Seed should be the mnemonic word list
func GetAddresses(seed string, num int) (string, error) {
	addresses := make([]liteclient.Address, num)

	// To maintain compatibility with the desktop wallet seed import
	// the first keypair should be generated directly from the word seed
	// with no decode into bytes etc.
	decodedSeed := []byte(seed)
	for i := 0; i < num; i++ {
		next, keys := cipher.GenerateDeterministicKeyPairsSeed([]byte(decodedSeed), 1)
		nextSeed := hex.EncodeToString(next)
		pub := cipher.PubKeyFromSecKey(keys[0])
		address := liteclient.Address{
			NextSeed: nextSeed,
			Secret:   keys[0].Hex(),
			Public:   pub.Hex(),
			Address:  cipher.AddressFromPubKey(pub).String(),
		}
		addresses[i] = address
		var err error
		// this is thrown away on last loop but its nicer than having special cases for first iteration
		decodedSeed, err = hex.DecodeString(nextSeed) 
		if err != nil {
			return "", err
		}
	}

	res, err := json.Marshal(addresses)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

// PrepareTransaction receives inputs and outputs and returns a signed transaction
// inputsBody and outputsBody are JSONified arrays of TransactionInput and TransactionOutput, respectively.
func PrepareTransaction(inputsBody string, outputsBody string) string {
	tx := liteclient.PrepareTransaction(inputsBody, outputsBody)
	return tx
}


func NewWordSeed() (string, error) {
	seed, err := bip39.NewDefaultMnemonic()
	if err != nil {
		panic(err)
	}

	return seed, nil
}

