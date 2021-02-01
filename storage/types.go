package storage

type MysqlConfig struct {
	Server    string `json:"server"`
	UserName  string `json:"user_name"`
	PassWord  string `json:"pass_word"`
	Database1 string `json:"database"`
}

// ============= Mysql
type CzzBlocks struct {
	Id                int64   `json:"id"`
	PreviousBlockHash string  `json:"previous_block_hash"`
	Hash              string  `json:"hash"`
	Height            int64   `json:"height"`
	Size              int64   `json:"size"`
	Confirmations     int64   `json:"confirmations"`
	Version           int32   `json:"version"`
	VersionHex        string  `json:"version_hex"`
	Merkleroot        string  `json:"merkleroot"`
	CTime             int64   `json:"c_time"`
	Nonce             uint64  `json:"nonce"`
	Bits              string  `json:"bits"`
	Difficulty        float64 `json:"difficulty"`
	IsMain            int     `json:"is_main"`
	UpdateTime        int     `json:"update_time"`
	CreateTime        int     `json:"create_time"`
}
