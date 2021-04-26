package common

import (
	"encoding/json"
	"github.com/classzz/committee-vote/chains"
	"github.com/classzz/committee-vote/scanning"
	"log"
	"os"
	"path/filepath"
)

type HttpConfig struct {
	Server   string `json:"server"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type Config struct {
	HttpServer HttpConfig      `json:"http_server"`
	Block      scanning.Config `json:"block"`
	Chains     chains.Config   `json:"chains"`
	PrivateKey string          `json:"private_key"`
	DebugLevel int             `json:"debug_level"`
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
