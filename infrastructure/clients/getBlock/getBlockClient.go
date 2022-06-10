package getBlock

import (
	"bytes"
	"encoding/json"
	"get-block/domain/balance"
	"io"
	"log"
	"net/http"
	"net/url"
)

const jsonRpc = "2.0"
const blockNumberId = "blockNumber"

type Config struct {
	BaseUrl *url.URL
	Key     string
}

type Client struct {
	Config *Config
}

type BlockNumber struct {
	Result string `json:"result"`
}

type Request struct {
	Id      string        `json:"id"`
	JsonRpc string        `json:"jsonRpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

func NewConfig(baseUrl string, key string) *Config {
	url, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return &Config{
		BaseUrl: url,
		Key:     key,
	}
}

func NewClient(config *Config) *Client {
	return &Client{
		Config: config,
	}
}

func (client *Client) GetBlockNumber() (string, error) {
	blockNumber := new(BlockNumber)

	resp, err := client.SendRequest("POST", client.Config.BaseUrl.String(), &Request{
		Id:      blockNumberId,
		JsonRpc: jsonRpc,
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
	})

	if err != nil {
		return blockNumber.Result, err
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return blockNumber.Result, err
		}
		err = json.Unmarshal(bodyBytes, blockNumber)
		if err != nil {
			return blockNumber.Result, err
		}
	}

	return blockNumber.Result, nil
}

func (client *Client) GetBlockByNumber(blockNumber string) (*balance.Block, error) {
	block := new(balance.Block)
	resp, err := client.SendRequest("POST", client.Config.BaseUrl.String(), &Request{
		Id:      blockNumberId,
		JsonRpc: jsonRpc,
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{blockNumber, true},
	})

	if err != nil {
		return block, err
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return block, err
		}
		err = json.Unmarshal(bodyBytes, block)
		if err != nil {
			return block, err
		}
	}

	return block, nil
}

func (client *Client) SendRequest(method, url string, request *Request) (resp *http.Response, err error) {
	data, err := json.Marshal(request)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if req != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("x-api-key", client.Config.Key)
	httpClient := &http.Client{}
	return httpClient.Do(req)
}
