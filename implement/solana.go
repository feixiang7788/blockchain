package implement

import (
	"encoding/json"
	"errors"
	"github.com/feixiang7788/blockchain/http"
)

type Solana struct {
	client http.SolanaHttpClient
}

type SolanaBlockHeightNumberInfo struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  int    `json:"result"`
	Id      int    `json:"id"`
}

type SolanaBlockInfo struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Blockhash string `json:"blockhash"`
	} `json:"result"`
	Id int `json:"id"`
}

func (this *Solana) CurrentBlockHeight() (int, error) {
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "getSlot",
		"params":  []interface{}{},
	}

	response, err := this.client.R().SetBody(requestBody).Post("")
	defer response.RawResponse.Body.Close()
	if err != nil {
		return 0, err
	}

	solanaBlockHeightInfo := SolanaBlockHeightNumberInfo{}
	err = json.Unmarshal(response.Body(), &solanaBlockHeightInfo)
	if err != nil {
		return 0, err
	}

	if solanaBlockHeightInfo.Result <= 0 {
		return 0, errors.New("get solanaBlockHeightInfo failed")
	}

	return solanaBlockHeightInfo.Result, nil
}

func (this *Solana) BlockInfo(blockHeightNumber int) (string, error) {
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "getBlock",
		"params": []interface{}{
			blockHeightNumber,
			map[string]interface{}{
				"encoding":                       "json",
				"transactionDetails":             "none",
				"rewards":                        false,
				"maxSupportedTransactionVersion": 0,
			},
		},
	}

	response, err := this.client.R().SetBody(requestBody).Post("")
	defer response.RawResponse.Body.Close()
	if err != nil {
		return "", err
	}

	solanaBlockHeightInfo := SolanaBlockInfo{}
	err = json.Unmarshal(response.Body(), &solanaBlockHeightInfo)
	if err != nil {
		return "", err
	}

	if solanaBlockHeightInfo.Result.Blockhash == "" {
		return "", errors.New("get solanaBlockInfo failed")
	}

	return solanaBlockHeightInfo.Result.Blockhash, nil
}

func NewSolana(clinet http.SolanaHttpClient) *Solana {
	return &Solana{
		client: clinet,
	}
}
