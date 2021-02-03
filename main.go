package main

import (
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

	if cfg.Block.Enabled {
		scanning := scanning.NewScanning(&cfg.Block, mysqlClient)
		go scanning.Start()
		defer scanning.Stop()
	}

	select {}
}
