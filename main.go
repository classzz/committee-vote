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
	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
)

var (
	cfg   common.Config
	RawDB *storage.RawDB
)

// 处理跨域请求,支持options访问
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token, UserName")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {

	var err error
	// Load configuration file
	common.LoadConfig(&cfg)

	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(true)))
	glogger.Verbosity(log.Lvl(cfg.DebugLevel))
	log.Root().SetHandler(glogger)

	db, err := leveldb.OpenFile("data/db", nil)
	if err != nil {
		panic(fmt.Sprintf("Leveldb err %s", err))
	}
	defer db.Close()

	eth := ethereum.NewClient(&cfg.Chains.EthClient, cfg.PrivateKey)
	heco := heco.NewClient(&cfg.Chains.HecoClient, cfg.PrivateKey)
	bsc := bsc.NewClient(&cfg.Chains.BscClient, cfg.PrivateKey)
	RawDB = &storage.RawDB{DB: db}

	scanning := scanning.NewScanning(&cfg.Block, RawDB)

	if scanning == nil {
		return
	}
	scanning.EthClient = eth
	scanning.HecoClient = heco
	scanning.BscClient = bsc
	go scanning.Start()
	defer scanning.Stop()

	route := gin.Default()

	// 允许使用跨域请求  全局中间件
	route.Use(CORSMiddleware())
	v1 := route.Group("/v1/")
	{
		v1.POST("/currentheight", CurrentHeight)             //设置访问的路由
		v1.POST("/getconvertitemall", GetConvertItemAll)     //设置访问的路由
		v1.POST("/getconvertitembymid", GetConvertItemByMid) //设置访问的路由
	}

	if err := route.Run(cfg.HttpServer.Server); err != nil {
		panic("err:" + err.Error())
	}
}
