package main

import (
	"github.com/classzz/committee-vote/chains/ethereum"
	"github.com/classzz/committee-vote/chains/heco"
	"github.com/classzz/committee-vote/common"
	"github.com/classzz/committee-vote/scanning"
	"github.com/classzz/committee-vote/storage"
)

var (
	cfg common.Config
)

func main() {

	// Load configuration file
	common.LoadConfig(&cfg)
	mysqlClient := storage.NewMysqlClient(cfg.Mysql)
	eth := ethereum.NewClient(&cfg.Chains, cfg.PrivateKey)
	heco := heco.NewClient(&cfg.Chains, cfg.PrivateKey)
	scanning := scanning.NewScanning(&cfg.Block, mysqlClient)

	if scanning == nil {
		return
	}
	scanning.EthClient = eth
	scanning.HecoClient = heco
	go scanning.Start()
	defer scanning.Stop()

	select {}
}
