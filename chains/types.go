package chains

type Config struct {
	EthClient  ClientInfo `json:"eth_client"`
	HecoClient ClientInfo `json:"heco_client"`
	BscClient  ClientInfo `json:"bsc_client"`
}

type ClientInfo struct {
	RpcHost         string `json:"rpc_host"`
	ContractAddress string `json:"contract_address"`
	SwapRouter      string `json:"swap_router"`
	Eth             string `json:"eth"`
	Czz             string `json:"czz"`
	Current         string `json:"current"`
}
