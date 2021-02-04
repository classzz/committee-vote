package main

import (
	"github.com/classzz/committee-vote/chains"
	"github.com/classzz/committee-vote/common"
	"github.com/classzz/committee-vote/scanning"
	"github.com/classzz/committee-vote/storage"
)

var (
	mysqlClient *storage.MysqlClient
	cfg         common.Config
)

func main() {

	// Load configuration file
	common.LoadConfig(&cfg)
	mysqlClient = storage.NewMysqlClient(cfg.Mysql)
	eth := chains.NewEthClient(&cfg.Chains, cfg.PrivateKey)

	if cfg.Block.Enabled {
		scanning := scanning.NewScanning(&cfg.Block, mysqlClient, eth)
		go scanning.Start()
		defer scanning.Stop()
	}

	select {}
}
