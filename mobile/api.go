package mobile

import (
	"encoding/json"

	"github.com/skycoin/skycoin-lite/liteclient"
)

// GetAddresses returns a series of addresses based on a seed and the number of addresses
// Seed is be hex-encoded bytes.
func GetAddresses(seed string, amount int) (string, error) {
	addresses := liteclient.GenerateAddresses(seed, amount)

	response, err := json.Marshal(addresses)
	if err != nil {
		panic(err)
	}

	return string(response), nil
}
