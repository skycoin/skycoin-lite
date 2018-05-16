package liteclient

import (
	"fmt"
	"github.com/skycoin/skycoin-lite/service"
	"encoding/hex"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/coin"
	"errors"
)

type Wallet struct {
	Seed string
	Addresses int
}

type Address struct {
	Address string
	Secret string
	Coins uint64
	Hours uint64
}

func Send(wlt Wallet, toAddr string, amount uint64) (string, error) {
	rawTransaction, err := PrepareTx(wlt, toAddr, amount)
	res, err := service.InjectTransaction(rawTransaction)

	return res, err
}

func PrepareTx(wlt Wallet, toAddr string, amount uint64) (string, error) {

	addresses, _ := Addresses(wlt.Seed, wlt.Addresses);
	stringifiedAddresses := make([]string, len(addresses))
	for i, address := range addresses {
		stringifiedAddresses[i] = address.Address
	}

	totalUtxos, err := service.GetOutputs(stringifiedAddresses)

	utxos, err := getSufficientOutputs(totalUtxos, amount)

	if err != nil {
		return "", err
	}

	bal, hours := func(utxos []service.Output) (uint64, uint64) {
		var c, h uint64
		for _, u := range utxos {
			c += u.GetCoins()
			h += u.GetHours()
		}
		return c, h
	}(utxos)

	var txOut []coin.TransactionOutput
	chgAmt := bal - amount
	chgHours := hours / 4
	chgAddr := stringifiedAddresses[0]
	if chgAmt > 0 {
		txOut = append(txOut,
			makeTxOut(toAddr, amount, chgHours/2),
			makeTxOut(chgAddr, chgAmt, chgHours/2))
	} else {
		txOut = append(txOut, makeTxOut(toAddr, amount, chgHours/2))
	}

	newTransaction := coin.Transaction{}

	for _, in := range utxos {
		newTransaction.PushInput(cipher.MustSHA256FromHex(*in.Hash))
	}

	for _, out := range txOut {
		newTransaction.PushOutput(out.Address, out.Coins, out.Hours)
	}

	keys := make([]cipher.SecKey, len(utxos))
	for i, in := range utxos {
		s := retrievePrivateKeyForAddress(addresses, *in.Address)
		k, err := cipher.SecKeyFromHex(s)

		if err != nil {
			return "", fmt.Errorf("invalid private key:%v", err)
		}

		keys[i] = k
	}

	newTransaction.SignInputs(keys)
	newTransaction.UpdateHeader()
	d := newTransaction.Serialize()

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(d), nil
}

func getSufficientOutputs(utxos []service.Output, amt uint64) ([]service.Output, error) {
	outMap := make(map[string][]service.Output)
	for _, u := range utxos {
		outMap[u.GetAddress()] = append(outMap[u.GetAddress()], u)
	}

	allUtxos := []service.Output{}
	var allBal uint64
	for _, utxos := range outMap {
		allBal += func(utxos []service.Output) uint64 {
			var bal uint64
			for _, u := range utxos {
				if u.GetCoins() == 0 {
					continue
				}
				bal += u.GetCoins()
			}
			return bal
		}(utxos)

		allUtxos = append(allUtxos, utxos...)
		if allBal >= amt {
			return allUtxos, nil
		}
	}

	return nil, errors.New("insufficient balance")
}

func makeTxOut(addr string, coins uint64, hours uint64) coin.TransactionOutput {
	out := coin.TransactionOutput{}
	out.Address = cipher.MustDecodeBase58Address(addr)
	out.Coins = coins
	out.Hours = hours
	return out
}

func retrievePrivateKeyForAddress(addresses []Address, address string) string {
	for _, a := range addresses {
		if a.Address == address {
			return a.Secret
		}
	}

	return ""
}

func Addresses(seed string, amount int) ([]Address, error) {
	_, secretKeys := cipher.GenerateDeterministicKeyPairsSeed([]byte(seed), amount)
	addresses := make([]Address, amount)
	for i, sec := range secretKeys {
		pub := cipher.PubKeyFromSecKey(sec)
		address := Address{
			cipher.AddressFromPubKey(pub).String(),
			sec.Hex(),
			0,
			0,
		}
		addresses[i] = address
	}

	return addresses, nil
}

func AddressesWithBalance(addresses []Address) ([]Address, error) {
	stringifiedAddresses := make([]string, len(addresses))
	for i, address := range addresses {
		stringifiedAddresses[i] = address.Address
	}

	outputs, _ := service.GetOutputs(stringifiedAddresses)

	for _, output := range outputs {
		for i, address := range addresses {
			if address.Address == output.GetAddress() {
				addresses[i].Coins += output.GetCoins()
				addresses[i].Hours += output.GetHours()
			}
		}
	}

	return addresses, nil
}

func Outputs(addresses []Address) ([]service.Output, error) {
	stringifiedAddresses := make([]string, len(addresses))
	for i, address := range addresses {
		stringifiedAddresses[i] = address.Address
	}

	return service.GetOutputs(stringifiedAddresses)
}
