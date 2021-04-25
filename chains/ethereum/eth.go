package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/classzz/classzz/btcjson"
	"github.com/classzz/classzz/cross"
	"github.com/classzz/committee-vote/chains"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/net/context"
	"math/big"
)

type EthClient struct {
	Cfg        *chains.ClientInfo
	Client     *ethclient.Client
	PrivateKey string
}

func NewClient(c *chains.ClientInfo, privateKey string) *EthClient {
	client, err := ethclient.Dial(c.RpcHost)
	if err != nil {
		log.Error("err", err)
	}

	ec := &EthClient{
		Client:     client,
		PrivateKey: privateKey,
	}

	return ec
}

// casting
func (ec *EthClient) Casting(items *btcjson.ConvertItemsResult) (*types.Transaction, error) {
	contractAddress := common.HexToAddress(ec.Cfg.ContractAddress)
	swaprouter := common.HexToAddress(ec.Cfg.SwapRouter)
	weth := common.HexToAddress(ec.Cfg.Eth)
	eczz := common.HexToAddress(ec.Cfg.Czz)

	instance, err := NewCommon(contractAddress, ec.Client)
	privateKey, err := crypto.HexToECDSA(ec.PrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ec.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := ec.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(3))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = gasPrice

	toaddresspuk, err := crypto.DecompressPubkey(items.PubKey)
	if err != nil || toaddresspuk == nil {
		toaddresspuk, err = crypto.UnmarshalPubkey(items.PubKey)
		if err != nil || toaddresspuk == nil {
			return nil, err
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
			return nil, err
		}

		log.Warn("tx sent", "txhash", tx.Hash().Hex(), "toaddress", toaddress, "fromaddress", fromAddress)
		return tx, nil
	}

	ethlist, err := instance.SwapBurnGetAmount(nil, big.NewInt(amountIn), paths, swaprouter)
	if err != nil {
		return nil, err
	}

	log.Info("paths amount", "ethlist", ethlist)
	if items.ToToken == "0x0000000000000000000000000000000000000000" {
		log.Info("ETH SwapTokenForEth", "toaddress", toaddress)
		tx, err := instance.SwapTokenForEth(auth, toaddress, Amount, items.MID, big.NewInt(0), swaprouter, weth, big.NewInt(10000000000000000))
		if err != nil {
			return nil, err
		}
		log.Warn("tx sent", "txhash", tx.Hash().Hex(), "toaddress", toaddress, "fromaddress", fromAddress)
		return tx, nil
	}

	if items.ToToken == ec.Cfg.Czz {
		log.Info("ETH Mint", "toaddress", toaddress)
		tx, err := instance.Mint(auth, items.MID, toaddress, Amount)
		if err != nil {
			return nil, err
		}
		log.Warn("tx sent", "txhash", tx.Hash().Hex(), "toaddress", toaddress, "fromaddress", fromAddress)
		return tx, nil
	}

	log.Info("ETH SwapToken", "toaddress", toaddress)
	tx, err := instance.SwapToken(auth, toaddress, Amount, items.MID, toToken, big.NewInt(0), swaprouter, weth, big.NewInt(10000000000000000))
	if err != nil {
		return nil, err
	}

	log.Warn("tx sent", "txhash", tx.Hash().Hex(), "toaddress", toaddress, "fromaddress", fromAddress)
	return tx, nil
}
