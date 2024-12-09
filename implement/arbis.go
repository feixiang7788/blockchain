package implement

import (
	"encoding/json"
	"errors"
	"game_server/leaf/log"
	"game_server/pkg/http"
	"strconv"
	"strings"
)

type Arbis struct {
	client http.ArbisHttpClient
}

type ArbisBlockHeightInfo struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  string `json:"result"`
}

type ArbisBlockInfo struct {
	Result struct {
		Hash string `json:"hash"`
	} `json:"result"`
}

func (this *Arbis) CurrentBlockHeight() (int, error) {
	params := map[string]string{
		"module": "proxy",
		"action": "eth_blockNumber",
		"apikey": "TYTX5R1IUX2WWA8ZUGYNAK1YIRK2RFJUVG",
	}

	response, err := this.client.R().SetQueryParams(params).Get("/api")
	log.Debug("Arbis获取区块最新高度返回结果：%v", response.String())
	defer response.RawResponse.Body.Close()
	if err != nil {
		return 0, err
	}

	solanaBlockHeightInfo := ArbisBlockHeightInfo{}
	err = json.Unmarshal(response.Body(), &solanaBlockHeightInfo)
	if err != nil {
		return 0, err
	}

	if solanaBlockHeightInfo.Result == "" {
		return 0, errors.New("get arbisBlockHieghtInfo failed1")
	}

	num, err := strconv.ParseInt(strings.Replace(solanaBlockHeightInfo.Result, "0x", "", 1), 16, 0)
	if solanaBlockHeightInfo.Result == "" {
		return 0, errors.New("get arbisBlockHieghtInfo failed2")
	}

	return int(num), nil
}

func (this *Arbis) BlockInfo(blockHeightNumber int) (string, error) {
	hexStr := strconv.FormatInt(int64(blockHeightNumber), 16)
	params := map[string]string{
		"module":  "proxy",
		"action":  "eth_getBlockByNumber",
		"tag":     hexStr,
		"boolean": "true",
		"apikey":  "TYTX5R1IUX2WWA8ZUGYNAK1YIRK2RFJUVG",
	}

	response, err := this.client.R().SetQueryParams(params).Get("https://api.arbiscan.io/api")
	log.Debug("Arbis获取区块高度：%v, 返回结果：%v", blockHeightNumber, response.String())
	defer response.RawResponse.Body.Close()
	if err != nil {
		return "", err
	}

	solanaBlockInfo := ArbisBlockInfo{}
	err = json.Unmarshal(response.Body(), &solanaBlockInfo)
	if err != nil {
		return "", err
	}

	if solanaBlockInfo.Result.Hash == "" {
		return "", errors.New("get arbisBlockInfo failed")
	}
	return solanaBlockInfo.Result.Hash, err
}

func NewArbis(clinet http.ArbisHttpClient) *Arbis {
	return &Arbis{
		client: clinet,
	}
}
