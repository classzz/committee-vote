package blocks

type Config struct {
	Enabled bool `json:"enabled"`

	// rpc
	RpcHost     string `json:"rpc_host"`
	RpcUsername string `json:"rpc_username"`
	RpcPassword string `json:"rpc_password"`
}
