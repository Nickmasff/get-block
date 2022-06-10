package getBlock

import (
	"encoding/json"
	"get-block/domain/balance"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

const lastBlock = "0xe3cbae"

type Config struct {
	BaseUrl *url.URL
	Key     string
}

type Client struct {
	Config *Config
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
	return lastBlock, nil
}

func (client *Client) GetBlockByNumber(blockNumber string) (*balance.Block, error) {
	res := balance.Block{}
	path, err := os.Getwd()
	if err != nil {
		return &res, err
	}
	return client.readFromFile(filepath.Join(path, "mocks/getBlockByNumber.json"))
}

func (client *Client) readFromFile(path string) (*balance.Block, error) {
	res := balance.Block{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return &res, err
	}
	_ = json.Unmarshal([]byte(file), &res)
	return &res, nil
}
