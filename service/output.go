package service

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type Output struct {
	Hash             *string `protobuf:"bytes,10,opt,name=hash" json:"hash,omitempty"`
	SrcTx            *string `protobuf:"bytes,11,opt,name=src_tx" json:"src_tx,omitempty"`
	Address          *string `protobuf:"bytes,12,opt,name=address" json:"address,omitempty"`
	Coins            *float64 `protobuf:"varint,13,opt,name=coins" json:"coins,string,omitempty"`
	Hours            *uint64 `protobuf:"varint,14,opt,name=hours" json:"hours,omitempty"`
}

type OutputResponse struct {
	Outputs []Output `json:"head_outputs"`
}

func (m *Output) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *Output) GetCoins() uint64 {
	if m != nil && m.Coins != nil {
		return uint64(*m.Coins * 1000000)
	}
	return 0
}

func (m *Output) GetHours() uint64 {
	if m != nil && m.Hours != nil {
		return *m.Hours
	}
	return 0
}

func GetOutputs(addrs []string) ([]Output, error) {
	if len(addrs) == 0 {
		return []Output{}, nil
	}

	resp, err := http.Get(NodeAddress + "/outputs?addresses=" + strings.Join(addrs, ","))

	v := OutputResponse{}

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body,&v)

	return v.Outputs, err
}