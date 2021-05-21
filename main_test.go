package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestGasPrice(t *testing.T) {

	client, err := ethclient.Dial("https://eth-mainnet.alchemyapi.io/v2/N1MfOz71XmlBniz-OWOZomT2zfgNCQxg")
	if err != nil {

	}

	gas, err := client.SuggestGasPrice(context.Background())
	fmt.Println(gas)
}
