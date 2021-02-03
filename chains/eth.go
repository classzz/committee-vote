package chains

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/classzz/classzz/btcjson"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"math/big"
)

type EthClient struct {
	Client     *ethclient.Client
	PrivateKey string
}

//wif: KxXPxci7wiXyKphnQ2VryVhdD8fMYdto8cEUo5NS3gLNkcX59S9g
//priv: 26ef857fb4f3b1f8e88cacea4e35c4d0d4bd3688cbdf6355aa6232e9fc3266fb
//pubk: [4 154 219 255 88 241 223 234 161 110 193 172 8 183 143 154 241 45 169 73 25 25 244 248 143 237 236 113 137 244 186 248 156 160 240 219 103 160 88 98 150 85 200 181 210 83 145 27 184 153 220 141 36 103 187 237 60 224 133 74 91 89 184 54 13]
//pubk compressed : [3 154 219 255 88 241 223 234 161 110 193 172 8 183 143 154 241 45 169 73 25 25 244 248 143 237 236 113 137 244 186 248 156]
//pub: 039adbff58f1dfeaa16ec1ac08b78f9af12da9491919f4f88fedec7189f4baf89c
//addressScript: 3871a56a7cabd0fd1f4610fdbc462af224059c65
//address: cqu8rft20j4aplglgcg0m0zx9tezgpvuv5w8c7d95l
// c4918edc5f8953a1a716324610271e0c6f891815

// casting
func (ec *EthClient) Casting(items *btcjson.ConvertItemsResult) error {

	address := common.HexToAddress("0x9a81F246Fb739387ca57fD06D454786Fe2338b21")
	instance, err := NewChains(address, ec.Client)
	privateKey, err := crypto.HexToECDSA(ec.PrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ec.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := ec.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(8888))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	toaddresspuk, err := crypto.DecompressPubkey(items.PubKey)
	if err != nil || toaddresspuk == nil {
		return err
	}

	toaddress := crypto.PubkeyToAddress(*toaddresspuk)
	tx, err := instance.Mint(auth, toaddress, items.Amount, items.MID)
	if err != nil {
		return err
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
	return nil
}
