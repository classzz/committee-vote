package heco

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/classzz/classzz/btcjson"
	"github.com/classzz/classzz/cross"
	"github.com/classzz/committee-vote/chains"
	"github.com/classzz/committee-vote/chains/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"math/big"
)

type HecoClient struct {
	Cfg        *chains.ClientInfo
	Client     *ethclient.Client
	PrivateKey string
}

func NewClient(c *chains.ClientInfo, privateKey string) *HecoClient {
	client, err := ethclient.Dial(c.RpcHost)
	if err != nil {
		fmt.Println(err)
	}

	ec := &HecoClient{
		Client:     client,
		PrivateKey: privateKey,
	}

	return ec
}

// casting
func (ec *HecoClient) Casting(items *btcjson.ConvertItemsResult) (string, error) {

	instance, err := ethereum.NewCommon(contractAddress, ec.Client)
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

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(128))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(0) // in units
	//auth.GasPrice = gasPrice

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
	paths := []common.Address{wht, hczz}

	if items.AssetType == cross.ExpandedTxConvert_Czz {
		fmt.Println("HECO mint toaddress", toaddress)
		tx, err := instance.Mint(auth, items.MID, toaddress, Amount)
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

	if items.ToToken == "0x0000000000000000000000000000000000000000" {
		fmt.Println("HECO SwapTokenForHt toaddress", toaddress)
		tx, err := instance.SwapTokenForEth(auth, toaddress, Amount, items.MID, ethlist[1], swaprouter, wht, big.NewInt(10000000000000000))
		if err != nil {
			return "", err
		}
		fmt.Printf("tx sent: %s toaddress %s fromaddress %s \r\n", tx.Hash().Hex(), toaddress, fromAddress)
		return tx.Hash().Hex(), nil
	}

	fmt.Println("HECO SwapToken toaddress", toaddress)
	tx, err := instance.SwapToken(auth, toaddress, Amount, items.MID, toToken, ethlist[1], swaprouter, wht, big.NewInt(1000000000000000))
	if err != nil {
		return "", err
	}
	fmt.Printf("tx sent: %s toaddress %s fromaddress %s \r\n", tx.Hash().Hex(), toaddress, fromAddress)
	return tx.Hash().Hex(), nil
}
