package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/classzz/classzz/btcjson"
	"github.com/classzz/classzz/cross"
	"github.com/classzz/committee-vote/chains"
	common2 "github.com/classzz/committee-vote/chains/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/net/context"
	"math/big"
	"strings"
	"time"
)

var (
	ChainName = "ETH"
	ChainID   = big.NewInt(1)
)

type EthClient struct {
	Cfg        *chains.ClientInfo
	Client     *ethclient.Client
	Tracker    *common2.GasTracker
	PrivateKey string
}

func (ec *EthClient) SyncTracker() {
	for {
		if Tracker, err := common2.GetGasTracker(); err != nil {
			log.Error("ETh SyncTracker", "err", err)
		} else {
			ec.Tracker = Tracker
		}
		time.Sleep(16 * time.Second)
	}
}

func NewClient(c *chains.ClientInfo, privateKey string) *EthClient {
	client, err := ethclient.Dial(c.RpcHost)
	if err != nil {
		log.Error("err", err)
	}

	ec := &EthClient{
		Client:     client,
		PrivateKey: privateKey,
		Cfg:        c,
	}
	go ec.SyncTracker()
	return ec
}

// casting
func (ec *EthClient) Casting(items *btcjson.ConvertItemsResult) (*types.Transaction, error) {

	countSplit := strings.Split(items.ToToken, "#")
	if items.AssetType != cross.ExpandedTxConvert_Czz {
		if len(countSplit) < 1 {
			return nil, fmt.Errorf("countSplit err %s", countSplit)
		}

		if !strings.Contains(ec.Cfg.SwapRouter, countSplit[1]) {
			return nil, fmt.Errorf("SwapRouter err %s", countSplit)
		}
	} else {
		countSplit[0] = ""
		countSplit = append(countSplit, "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	}
	contractAddress := common.HexToAddress(ec.Cfg.ContractAddress)
	swaprouter := common.HexToAddress(countSplit[1])
	current := common.HexToAddress(ec.Cfg.Current)
	czz := common.HexToAddress(ec.Cfg.Czz)

	instance, err := common2.NewCommon(contractAddress, ec.Client)
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

	//gasPrice, err := ec.Client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	return nil, err
	//}

	gasPrice := big.NewInt(ec.Tracker.Data.Fast)
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, ChainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = gasPrice   // in wei

	toaddresspuk, err := crypto.DecompressPubkey(items.PubKey)
	if err != nil || toaddresspuk == nil {
		toaddresspuk, err = crypto.UnmarshalPubkey(items.PubKey)
		if err != nil || toaddresspuk == nil {
			return nil, err
		}
	}

	toaddress := crypto.PubkeyToAddress(*toaddresspuk)
	toToken := common.HexToAddress(countSplit[0])
	Amount := big.NewInt(0).Sub(items.Amount, items.FeeAmount)

	amountIn := int64(uint64(800000)) * gasPrice.Int64()
	paths := []common.Address{current, czz}
	ethlist, err := instance.SwapBurnGetAmount(nil, big.NewInt(amountIn), paths, swaprouter)
	if err != nil {
		return nil, err
	}

	log.Info(ChainName, "gasprice", gasPrice, "ethlist", ethlist, "mint fromAddress", fromAddress, "toaddress", toaddress)
	gaspaths := []common.Address{czz, current}
	if items.AssetType == cross.ExpandedTxConvert_Czz {
		tx, err := instance.MintWithGas(auth, items.MID, toaddress, Amount, ethlist[len(ethlist)-1], swaprouter, gaspaths, big.NewInt(10000000000000000))
		if err != nil {
			return nil, err
		}
		log.Info(ChainName, "tx hash ", tx.Hash().Hex())
		return tx, nil
	}

	if countSplit[0] == "0x0000000000000000000000000000000000000000" {
		tx, err := instance.SwapTokenForEthWithPath(auth, toaddress, Amount, items.MID, ethlist[len(ethlist)-1], swaprouter, gaspaths, big.NewInt(10000000000000000))
		if err != nil {
			return nil, err
		}
		log.Info(ChainName, "tx hash ", tx.Hash().Hex())
		return tx, nil
	}

	if countSplit[0] == ec.Cfg.Czz {
		tx, err := instance.MintWithGas(auth, items.MID, toaddress, Amount, ethlist[len(ethlist)-1], swaprouter, gaspaths, big.NewInt(10000000000000000))
		if err != nil {
			return nil, err
		}
		log.Info(ChainName, "tx hash ", tx.Hash().Hex())
		return tx, nil
	}

	userPath := []common.Address{czz, current, toToken}
	tx, err := instance.SwapTokenWithPath(auth, toaddress, Amount, items.MID, ethlist[len(ethlist)-1], swaprouter, userPath, gaspaths, big.NewInt(1000000000000000))
	if err != nil {
		return nil, err
	}
	log.Info(ChainName, "tx hash ", tx.Hash().Hex())
	return tx, nil
}
