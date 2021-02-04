package heco

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"testing"
)

func TestEthClient_Casting(t *testing.T) {

	client, err := ethclient.Dial("http://121.36.88.105:9223")
	if err != nil {
		log.Fatal(err)
	}
	eth := &HecoClient{Client: client}
	eth.Casting(nil)

}
