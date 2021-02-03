package storage

import (
	"database/sql"
	"fmt"
	"github.com/classzz/classzz/btcjson"
	_ "github.com/go-sql-driver/mysql"
	"math/big"
	"time"
)

const (
	NETWORK = "tcp"
	PORT    = 3306
)

type MysqlClient struct {
	mysqlDB *sql.DB
}

func NewMysqlClient(cfg MysqlConfig) *MysqlClient {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", cfg.UserName, cfg.PassWord, NETWORK, cfg.Server, PORT, cfg.Database1)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return nil
	}

	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数

	conn := &MysqlClient{
		mysqlDB: DB,
	}

	return conn
}

func (conn *MysqlClient) Stop() {
	conn.mysqlDB.Close()
}

func (c *MysqlClient) BlockInstall(block *CzzBlocks) int {

	insertSQL := "insert into blocks (previous_block_hash, `hash`, height, `size`, confirmations, version, version_hex, merkleroot, c_time, nonce, bits, difficulty, is_main) values (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err := c.mysqlDB.Exec(insertSQL, block.PreviousBlockHash, block.Hash, block.Height, block.Size, block.Confirmations, block.Version, block.VersionHex, block.Merkleroot, block.CTime, block.Nonce, block.Bits, block.Difficulty, block.IsMain)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return 1
	}
	return 0
}

func (c *MysqlClient) ConvertItemInstall(info *btcjson.ConvertItemsResult) int {
	insertSQL := "insert into convert_items (mid, asset_type, convert_type, pubkey, `height`, exttxhash, `index`, `amount`) values (?,?,?,?,?,?,?,?)"
	_, err := c.mysqlDB.Exec(insertSQL, info.MID, info.AssetType, info.ConvertType, info.PubKey, 0, info.ExtTxHash, 0, info.Amount)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return 1
	}
	return 0
}

func (conn *MysqlClient) FindConvertItem(mid *big.Int) *int64 {

	rows, err := conn.mysqlDB.Query("select mid from convert_items where `mid` = ? ", mid.Int64())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	var ok *int64

	if rows.Next() {
		rows.Scan(&ok)
	}

	return ok
}

func (c *MysqlClient) BlockFindMaxHeight() int64 {

	rows, err := c.mysqlDB.Query("select max(height)  from blocks")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer rows.Close()
	if rows.Next() {
		var height int64
		rows.Scan(&height)
		return height
	}
	return 0
}
