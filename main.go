package main

import (
	"fmt"
	"github.com/classzz/committee-vote/chains/bsc"
	"github.com/classzz/committee-vote/chains/ethereum"
	"github.com/classzz/committee-vote/chains/heco"
	"github.com/classzz/committee-vote/common"
	"github.com/classzz/committee-vote/scanning"
	"github.com/classzz/committee-vote/storage"
	"github.com/ethereum/go-ethereum/log"
	"github.com/syndtr/goleveldb/leveldb"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

var (
	cfg common.Config
	db  *leveldb.DB
)

func main() {

	var err error
	// Load configuration file
	common.LoadConfig(&cfg)

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(true)))
	glogger.Verbosity(log.Lvl(cfg.DebugLevel))
	log.Root().SetHandler(glogger)

	db, err = leveldb.OpenFile("data/db", nil)
	if err != nil {
		panic(fmt.Sprintf("Leveldb err %s", err))
	}
	defer db.Close()

	eth := ethereum.NewClient(&cfg.Chains.EthClient, cfg.PrivateKey)
	heco := heco.NewClient(&cfg.Chains.HecoClient, cfg.PrivateKey)
	bsc := bsc.NewClient(&cfg.Chains.BscClient, cfg.PrivateKey)
	rawdb := &storage.RawDB{DB: db}

	scanning := scanning.NewScanning(&cfg.Block, rawdb)

	if scanning == nil {
		return
	}
	scanning.EthClient = eth
	scanning.HecoClient = heco
	scanning.BscClient = bsc
	go scanning.Start()
	defer scanning.Stop()

	rpc.Register(new(RPCService))
	sock, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error("listen error:", err)
	}

	for {
		conn, err := sock.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
