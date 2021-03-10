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

	address := common.HexToAddress("0xf1e41979540BC776b2Fe8961f4C17E81Bf03894e")
	instance, err := NewHeco(address, ec.Client)
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
	path1 := common.HexToAddress("0xE30d43717DB115D2f205acCfeCedec67aDfDE089")
	path2 := common.HexToAddress("0x11D89c7966db767F2c933E7F1E009CD740b03677")
	paths := []common.Address{path1, path2}

	ethlist, err := instance.SwapBurnGetAmount(nil, big.NewInt(amountIn), paths)
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

	tx, err := instance.SwapToken(auth, toaddress, Amount, big.NewInt(0).Add(items.MID, big.NewInt(100000)), toToken, ethlist[1], big.NewInt(1000000000000000))
	if err != nil {
		return "", err
	}
	fmt.Printf("tx sent: %s toaddress %s fromaddress %s \r\n", tx.Hash().Hex(), toaddress, fromAddress)
	return tx.Hash().Hex(), nil
}
