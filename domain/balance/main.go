package balance

type Block struct {
	Result struct {
		Hash         string `json:"hash"`
		Number       string `json:"number"`
		Transactions []struct {
			From  string `json:"from"`
			To    string `json:"to"`
			Value string `json:"value"`
		} `json:"transactions"`
	} `json:"result"`
}

type GetBlockClientInterface interface {
	GetBlockNumber() (string, error)
	GetBlockByNumber(blockNumber string) (*Block, error)
}
