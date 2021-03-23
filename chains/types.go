package chains

type Config struct {
	EthRpc  string `json:"eth_rpc"`
	HecoRpc string `json:"heco_rpc"`
	BscRpc  string `json:"bsc_rpc"`
}
