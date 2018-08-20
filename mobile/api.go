package mobile

import (
	"encoding/json"
	"encoding/hex"
	"github.com/skycoin/skycoin-lite/liteclient"
	"github.com/skycoin/skycoin/src/cipher/go-bip39"
)

func GetAddresses(seed string, num int) (addr string, err error) {
	defer func() {
		r := recover()
		err, _ = r.(error)
	}()
	hexSeed := hex.EncodeToString([]byte(seed))
	addresses := liteclient.GenerateAddresses(hexSeed, num);
	
	byteaddr, err := json.Marshal(addresses)
	if err != nil {
		return "", err
	}
	addr = string(byteaddr)

	return addr, nil
}

// PrepareTransaction receives inputs and outputs and returns a signed transaction
// inputsBody and outputsBody are JSONified arrays of TransactionInput and TransactionOutput.
func PrepareTransaction(inputsBody string, outputsBody string) (tx string, err error) {

	defer func() {
		r := recover()
		err, _ = r.(error)
	}()

	tx = liteclient.PrepareTransaction(inputsBody, outputsBody)

	return
}


func NewWordSeed() (string, error) {
	seed, err := bip39.NewDefaultMnemonic()
	if err != nil {
		panic(err)
	}

	return seed, nil
}

