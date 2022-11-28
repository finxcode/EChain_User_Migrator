package app

type Response struct {
	ContractAddress string `json:"contractAddress"`
	PrivateKey      string `json:"prvKey"`
	PubKey          string `json:"pubKey"`
	TransactionHash string `json:"transactionHash"`
	Result          int8   `json:"result"`
}
