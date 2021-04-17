package heco

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
	"golang.org/x/net/context"
	"math/big"
)

var (
	contractAddress = common.HexToAddress("0x711D839CD1E6E81B971F5b6bBB4a6BD7C4B60Ac6")
	swaprouter      = common.HexToAddress("0xED7d5F38C79115ca12fe6C0041abb22F0A06C300")
	wht             = common.HexToAddress("0x5545153ccfca01fbd7dd11c0b23ba694d9509a6f")
	hczz            = common.HexToAddress("0x112489c758D405874e9Ece0586FD50B315216fcA")
)

type HecoClient struct {
	Client     *ethclient.Client
	PrivateKey string
}

func NewClient(c *chains.Config, private_key string) *HecoClient {
	client, err := ethclient.Dial(c.HecoRpc)
	if err != nil {
		fmt.Println(err)
	}

	ec := &HecoClient{
		Client:     client,
		PrivateKey: private_key,
	}

	return ec
}

// casting
func (ec *HecoClient) Casting(items *btcjson.ConvertItemsResult) (string, error) {

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

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(128))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(0) // in units
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
