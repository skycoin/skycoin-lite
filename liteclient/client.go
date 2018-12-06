package liteclient

import (
	"encoding/hex"
	"encoding/json"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/base58"
	"github.com/skycoin/skycoin/src/coin"
)

// Address includes a skycoin address, a public and secret key
// and the next seed to generate the next address from
type Address struct {
	NextSeed string
	Secret   string
	Public   string
	Address  string
}

// TransactionInput represents a transaction input
type TransactionInput struct {
	Hash   string
	Secret string
}

// TransactionOutput represents a transaction output
type TransactionOutput struct {
	Address string
	Coins   uint64
	Hours   uint64
}

// GenerateAddress generates an address from a seed. The seed should be hex-encoded bytes.
func GenerateAddress(seed string) Address {
	return GenerateAddresses(seed, 1)[0]
}

// GenerateAddresses generates addresses from a seed. The seed should be hex-encoded bytes.
func GenerateAddresses(seed string, num int) []Address {
	addresses := make([]Address, num)

	nextSeed := seed
	for i := 0; i < num; i++ {
		decodedSeed, err := hex.DecodeString(nextSeed)
		if err != nil {
			panic(err)
		}

		next, keys := cipher.GenerateDeterministicKeyPairsSeed([]byte(decodedSeed), 1)
		nextSeed = hex.EncodeToString(next)
		pub := cipher.PubKeyFromSecKey(keys[0])
		address := Address{
			NextSeed: nextSeed,
			Secret:   keys[0].Hex(),
			Public:   pub.Hex(),
			Address:  cipher.AddressFromPubKey(pub).String(),
		}

		addresses[i] = address
	}

	return addresses
}

// PrepareTransaction receives inputs and outputs and returns a signed transaction
// inputsBody and outputsBody are JSONified arrays of TransactionInput and TransactionOutput, respectively.
func PrepareTransaction(inputsBody string, outputsBody string) string {
	newTransaction := buildTransaction(inputsBody, outputsBody, nil, false)
	d := newTransaction.Serialize()

	return hex.EncodeToString(d)
}

// PrepareTransactionWithSignatures receives inputs, outputs,and the signatures and returns.
// inputsBody and outputsBody are JSONified arrays of TransactionInput and TransactionOutput, respectively.
// signatureList is a JSONified array of strings.
func PrepareTransactionWithSignatures(inputsBody string, outputsBody string, signatureList string) string {
	var signatures []string
	if err := json.Unmarshal([]byte(signatureList), &signatures); err != nil {
		panic(err)
	}

	newTransaction := buildTransaction(inputsBody, outputsBody, signatures, false)
	d := newTransaction.Serialize()

	return hex.EncodeToString(d)
}

// GetTransactionInnerHash receives the inputs and outputs, creates a transaction and returns its inner hash.
// inputsBody and outputsBody are JSONified arrays of TransactionInput and TransactionOutput, respectively.
func GetTransactionInnerHash(inputsBody string, outputsBody string) string {
	newTransaction := buildTransaction(inputsBody, outputsBody, nil, true)

	return newTransaction.InnerHash.Hex()
}

// Creates a coin.Transaction using the given lists of inputs, outputs and signatures. If signatureList is nil or
// empty the signatures are created using the Secret property of each input.
// inputsBody and outputsBody are JSONified arrays of TransactionInput and TransactionOutput, respectively.
// signatureList is a list of signatures as Base58 strings. If ignoreSigning == true, then the transaction is not signed.
func buildTransaction(inputsBody string, outputsBody string, signatureList []string, ignoreSigning bool) coin.Transaction {
	var inputs []TransactionInput
	var outputs []TransactionOutput

	if err := json.Unmarshal([]byte(inputsBody), &inputs); err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(outputsBody), &outputs); err != nil {
		panic(err)
	}

	newTransaction := coin.Transaction{}

	keys := make([]cipher.SecKey, len(inputs))

	for i, in := range inputs {
		if len(signatureList) == 0 && !ignoreSigning {
			k, err := cipher.SecKeyFromHex(in.Secret)
			if err != nil {
				panic(err)
			}

			keys[i] = k
		}

		inputHash, err := cipher.SHA256FromHex(in.Hash)
		if err != nil {
			panic(err)
		}

		newTransaction.PushInput(inputHash)
	}

	for _, out := range outputs {
		addr, err := cipher.DecodeBase58Address(out.Address)
		if err != nil {
			panic(err)
		}

		if addr.Null() {
			panic("output address is the null address")
		}

		newTransaction.PushOutput(addr, out.Coins, out.Hours)
	}

	if !ignoreSigning {
		if len(signatureList) == 0 {
			newTransaction.SignInputs(keys)
		} else {
			newTransaction.Sigs = make([]cipher.Sig, len(signatureList))
			for i, sig := range signatureList {
				sigBytes, err := base58.Base582Hex(sig)
				if err != nil {
					panic(err)
				}
				newTransaction.Sigs[i] = cipher.NewSig(sigBytes)
			}
		}
	}

	newTransaction.UpdateHeader()

	if !ignoreSigning {
		if err := newTransaction.Verify(); err != nil {
			panic(err)
		}
	}

	return newTransaction
}
