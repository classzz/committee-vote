package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/classzz/classzz/btcjson"
	"github.com/classzz/classzz/cross"
	"github.com/classzz/committee-vote/chains"
	common3 "github.com/classzz/committee-vote/chains/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/net/context"
	"math/big"
)

var (
	contractAddress = common.HexToAddress("0x9AC88c5136240312f8817dBB99497aCe62b03F12")
	swaprouter      = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	weth            = common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	eczz            = common.HexToAddress("0x150BbCfF6b1B4D528b48f1A300585Dea0b6490B6")
)

type EthClient struct {
	Client     *ethclient.Client
	PrivateKey string
}

func NewClient(c *chains.Config, private_key string) *EthClient {
	client, err := ethclient.Dial(c.EthRpc)
	if err != nil {
		log.Error("err", err)
	}

	ec := &EthClient{
		Client:     client,
		PrivateKey: private_key,
	}

	return ec
}

// casting
func (ec *EthClient) Casting(items *btcjson.ConvertItemsResult) (string, error) {

	instance, err := common3.NewCommon(contractAddress, ec.Client)
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

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(3))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(800000) // in units
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

	amountIn := int64(uint64(800000)) * gasPrice.Int64()
	paths := []common.Address{weth, eczz}

	if items.AssetType == cross.ExpandedTxConvert_Czz {
		log.Info("ETH mint", "toaddress", toaddress)
		tx, err := instance.Mint(auth, items.MID, toaddress, Amount)
		if err != nil {
			return "", err
		}

		log.Warn("tx sent", "txhash", tx.Hash().Hex(), "toaddress", toaddress, "fromaddress", fromAddress)
		return tx.Hash().Hex(), nil
	}

	ethlist, err := instance.SwapBurnGetAmount(nil, big.NewInt(amountIn), paths, swaprouter)
	if err != nil {
		return "", err
	}

	log.Info("paths amount", "ethlist", ethlist)
	if items.ToToken == "0x0000000000000000000000000000000000000000" {
		log.Info("ETH SwapTokenForEth", "toaddress", toaddress)
		tx, err := instance.SwapTokenForEth(auth, toaddress, Amount, items.MID, ethlist[1], swaprouter, weth, big.NewInt(10000000000000000))
		if err != nil {
			return "", err
		}
		log.Warn("tx sent", "txhash", tx.Hash().Hex(), "toaddress", toaddress, "fromaddress", fromAddress)
		return tx.Hash().Hex(), nil
	}

	log.Info("ETH SwapToken", "toaddress", toaddress)
	tx, err := instance.SwapToken(auth, toaddress, Amount, items.MID, toToken, ethlist[1], swaprouter, weth, big.NewInt(10000000000000000))
	if err != nil {
		return "", err
	}

	log.Warn("tx sent", "txhash", tx.Hash().Hex(), "toaddress", toaddress, "fromaddress", fromAddress)
	return tx.Hash().Hex(), nil
}
