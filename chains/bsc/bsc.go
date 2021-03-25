package bsc

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
	contractAddress = common.HexToAddress("0xb39E84c6AD0574af30fb5f0185ad2d4f2DBa4262")
	swaprouter      = common.HexToAddress("0xD99D1c33F9fC3444f8101754aBC46c52416550D1")
	wbnb            = common.HexToAddress("0xae13d989dac2f0debff460ac112a837c89baa7cd")
	bczz            = common.HexToAddress("0x507B8283aD724aA06dF09Df3d1D6eb3816EE51d5")
)

type BscClient struct {
	Client     *ethclient.Client
	PrivateKey string
}

func NewClient(c *chains.Config, private_key string) *BscClient {
	client, err := ethclient.Dial(c.BscRpc)
	if err != nil {
		log.Fatal(err)
	}

	ec := &BscClient{
		Client:     client,
		PrivateKey: private_key,
	}

	return ec
}

// casting
func (ec *BscClient) Casting(items *btcjson.ConvertItemsResult) (string, error) {

	instance, err := NewBsc(contractAddress, ec.Client)
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

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(97))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(800000) // in units
	auth.GasPrice = gasPrice

	toaddresspuk, err := crypto.DecompressPubkey(items.PubKey)
	if err != nil || toaddresspuk == nil {
		toaddresspuk, err = crypto.UnmarshalPubkey(items.PubKey)
		if err != nil || toaddresspuk == nil {
			return "", err
		}
	}

	toaddress := crypto.PubkeyToAddress(*toaddresspuk)
	toToken := common.HexToAddress(items.ToToken)
	Amount := big.NewInt(0).Sub(items.Amount, items.FeeAmount)
	amountIn := int64(auth.GasLimit) * gasPrice.Int64()
	paths := []common.Address{wbnb, bczz}

	if items.AssetType == cross.ExpandedTxConvert_Czz {
		fmt.Println("BSC mint toaddress", toaddress)
		tx, err := instance.Mint(auth, toaddress, Amount)
		if err != nil {
			return "", err
		}

		fmt.Printf("tx sent: %s toaddress %s fromaddress %s \r\n", tx.Hash().Hex(), toaddress, fromAddress)
		return tx.Hash().Hex(), nil
	}

	ethlist, err := instance.SwapBurnGetAmount(nil, big.NewInt(amountIn), paths, swaprouter)
	if err != nil {
		return "", err
	}
	fmt.Println("paths amount", ethlist)

	if items.ToToken == "0x0" {
		fmt.Println("BSC SwapTokenForHt toaddress", toaddress)
		tx, err := instance.SwapTokenForBsc(auth, toaddress, Amount, items.MID, big.NewInt(0), swaprouter, wbnb, big.NewInt(10000000000000000))
		if err != nil {
			return "", err
		}
		fmt.Printf("tx sent: %s toaddress %s fromaddress %s \r\n", tx.Hash().Hex(), toaddress, fromAddress)
		return tx.Hash().Hex(), nil
	}

	fmt.Println("BSC SwapToken toaddress", toaddress)
	tx, err := instance.SwapToken(auth, toaddress, Amount, items.MID, toToken, big.NewInt(0), swaprouter, wbnb, big.NewInt(1000000000000000))
	if err != nil {
		return "", err
	}
	fmt.Printf("tx sent: %s toaddress %s fromaddress %s \r\n", tx.Hash().Hex(), toaddress, fromAddress)
	return tx.Hash().Hex(), nil
}
