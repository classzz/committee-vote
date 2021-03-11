package heco

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/classzz/classzz/btcjson"
	"github.com/classzz/classzz/cross"
	"github.com/classzz/committee-vote/chains"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"log"
	"math/big"
)

var (
	contractAddress = common.HexToAddress("0x034d0162892893e688DC53f3194160f06EBf265E")
	router          = common.HexToAddress("0x539A9Fbb81D1D2DC805c698B55C8DF81cbA6b350")
	wht             = common.HexToAddress("0xA9e7417c676F70E5a13c919e78FB1097166568C5")
	hczz            = common.HexToAddress("0xF8444BF82C634d6F575545dbb6B77748bB1e3e19")
)

type HecoClient struct {
	Client     *ethclient.Client
	PrivateKey string
}

func NewClient(c *chains.Config, private_key string) *HecoClient {
	client, err := ethclient.Dial(c.HecoRpc)
	if err != nil {
		log.Fatal(err)
	}

	ec := &HecoClient{
		Client:     client,
		PrivateKey: private_key,
	}

	return ec
}

// casting
func (ec *HecoClient) Casting(items *btcjson.ConvertItemsResult) (string, error) {

	instance, err := NewHeco(contractAddress, ec.Client)
	privateKey, err := crypto.HexToECDSA(ec.PrivateKey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ec.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := ec.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(256))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(600000) // in units
	auth.GasPrice = gasPrice

	amountIn := int64(auth.GasLimit) * gasPrice.Int64()
	paths := []common.Address{hczz, wht}

	ethlist, err := instance.SwapBurnGetAmount(nil, big.NewInt(amountIn), paths, router)
	if err != nil {
		return "", err
	}

	toaddresspuk, err := crypto.DecompressPubkey(items.PubKey)
	if err != nil || toaddresspuk == nil {
		toaddresspuk, err = crypto.UnmarshalPubkey(items.PubKey)
		if err != nil || toaddresspuk == nil {
			return "", err
		}
	}

	toaddress := crypto.PubkeyToAddress(*toaddresspuk)
	toToken := common.HexToAddress(items.ToToken)
	Amount := big.NewInt(0).Mul(big.NewInt(0).Sub(items.Amount, items.FeeAmount), big.NewInt(10000000000))

	if items.AssetType == cross.ExpandedTxConvert_Czz {
		tx, err := instance.Mint(auth, toaddress, Amount)
		if err != nil {
			return "", err
		}

		fmt.Printf("tx sent: %s toaddress %s fromaddress %s \r\n", tx.Hash().Hex(), toaddress, fromAddress)
		return tx.Hash().Hex(), nil
	}

	tx, err := instance.SwapToken(auth, toaddress, Amount, items.MID, toToken, ethlist[1], router, wht, big.NewInt(1000000000000000))
	if err != nil {
		return "", err
	}
	fmt.Printf("tx sent: %s toaddress %s fromaddress %s \r\n", tx.Hash().Hex(), toaddress, fromAddress)
	return tx.Hash().Hex(), nil
}
