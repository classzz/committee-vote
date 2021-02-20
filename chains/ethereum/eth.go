package ethereum

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

type EthClient struct {
	Client     *ethclient.Client
	PrivateKey string
}

func NewClient(c *chains.Config, private_key string) *EthClient {
	client, err := ethclient.Dial(c.EthRpc)
	if err != nil {
		log.Fatal(err)
	}

	ec := &EthClient{
		Client:     client,
		PrivateKey: private_key,
	}

	return ec
}

// casting
func (ec *EthClient) Casting(items *btcjson.ConvertItemsResult) (string, error) {

	address := common.HexToAddress("0x1E2E51B940cFB54E6f8F878Dc06eF7A5386364AE")
	instance, err := NewEthereum(address, ec.Client)
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

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(8888))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(397570) // in units
	auth.GasPrice = gasPrice

	toaddresspuk, err := crypto.DecompressPubkey(items.PubKey)
	if err != nil || toaddresspuk == nil {
		return "", err
	}

	toaddress := crypto.PubkeyToAddress(*toaddresspuk)
	//toToken := common.HexToAddress(items.ToToken)
	Amount := big.NewInt(0).Mul(items.Amount, big.NewInt(10000000000))

	if items.AssetType == cross.ExpandedTxConvert_Czz {
		tx, err := instance.Mint(auth, toaddress, Amount)
		if err != nil {
			return "", err
		}

		fmt.Printf("tx sent: %s toaddress %s fromaddress %s \r\n", tx.Hash().Hex(), toaddress, fromAddress)
		return tx.Hash().Hex(), nil
	}

	tx, err := instance.SwapTokenForEth(auth, toaddress, Amount, items.MID)
	if err != nil {
		return "", err
	}

	fmt.Printf("tx sent: %s toaddress %s fromaddress %s \r\n", tx.Hash().Hex(), toaddress, fromAddress)
	return tx.Hash().Hex(), nil
}
