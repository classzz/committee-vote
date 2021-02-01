package common

import (
	"encoding/json"
	"github.com/classzz/committee-vote/blocks"
	"github.com/classzz/committee-vote/storage"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Mysql storage.MysqlConfig `json:"mysql"`
	Block blocks.Config       `json:"block"`
}

func LoadConfig(cfg *Config) {
	// Default common.
	configFileName := "config.json"
	if len(os.Args) > 1 {
		configFileName = os.Args[1]
	}
	configFileName, _ = filepath.Abs(configFileName)
	log.Printf("Loading config: %v", configFileName)

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal("File error: ", err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&cfg); err != nil {
		log.Fatal("Config error: ", err.Error())
	}
}

func (cfg *Config) GetConfig() *Config {
	return cfg
}
