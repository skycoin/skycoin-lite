package service

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

type InjectTxnReq struct {
	RawTransaction               string `json:"raw,omitempty"`
}

func InjectTransaction(rawtx string) (string, error) {
	url := NodeAddress + "/transaction"

	req := InjectTxnReq{
		RawTransaction: rawtx,
	}

	jsonRequest, _ := json.Marshal(req)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonRequest))

	body, _ := ioutil.ReadAll(resp.Body)
	text := string(body)

	return text, err
}
