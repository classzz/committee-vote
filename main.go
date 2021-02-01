package main

import (
	"github.com/classzz/committee-vote/blocks"
	"github.com/classzz/committee-vote/common"
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

	if cfg.Block.Enabled {
		scanning := blocks.NewScanning(&cfg.Block, mysqlClient)
		go scanning.Start()
		defer scanning.Stop()
	}

}
