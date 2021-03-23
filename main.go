package main

import (
	"github.com/classzz/committee-vote/chains/bsc"
	"github.com/classzz/committee-vote/chains/ethereum"
	"github.com/classzz/committee-vote/chains/heco"
	"github.com/classzz/committee-vote/common"
	"github.com/classzz/committee-vote/scanning"
	"github.com/classzz/committee-vote/storage"
	"github.com/ethereum/go-ethereum/log"
	"os"
)

var (
	cfg common.Config
)

func main() {

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(true)))
	glogger.Verbosity(log.Lvl(5))
	log.Root().SetHandler(glogger)

	// Load configuration file
	common.LoadConfig(&cfg)
	mysqlClient := storage.NewMysqlClient(cfg.Mysql)
	eth := ethereum.NewClient(&cfg.Chains, cfg.PrivateKey)
	heco := heco.NewClient(&cfg.Chains, cfg.PrivateKey)
	bsc := bsc.NewClient(&cfg.Chains, cfg.PrivateKey)
	scanning := scanning.NewScanning(&cfg.Block, mysqlClient)

	if scanning == nil {
		return
	}
	scanning.EthClient = eth
	scanning.HecoClient = heco
	scanning.BscClient = bsc
	go scanning.Start()
	defer scanning.Stop()

	select {}
}
