package wallet

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math"
	"math/big"

	web3 "github.com/johnnycx127/go-web3"
	"github.com/johnnycx127/go-web3/dto"
	"github.com/johnnycx127/go-web3/eth/block"
	"github.com/johnnycx127/go-web3/providers"
	"github.com/shopspring/decimal"
)

type EthWallet struct {
	client     *web3.Web3
	passphrase string
}

func newEthWallet(config map[string]interface{}) *EthWallet {
	connection := web3.NewWeb3(providers.NewHTTPProvider(config["host"].(string), 10, false))
	passphraseByte := []byte(config["passphrase"].(string))
	sha256Byte := sha256.Sum256(passphraseByte)
	wallet := EthWallet{
		client:     connection,
		passphrase: hex.EncodeToString(sha256Byte[:]),
	}
	return &wallet
}

func (ethw *EthWallet) GetNewAccount() (string, string, error) {
	account, err := ethw.client.Personal.NewAccount(ethw.passphrase)
	if err != nil {
		return "", "", err
	}
	return account, account, nil
}

func (ethw *EthWallet) GetBalance(account string, minConf int) (string, error) {
	amountWei, err := ethw.client.Eth.GetBalance(account, block.LATEST)
	if err != nil {
		return "", err
	}
	unitEth := decimal.NewFromFloat(math.Pow(10, (-18)))
	amountDecimal := decimal.NewFromBigInt(amountWei.ToBigInt(), 0)
	return unitEth.Mul(amountDecimal).Round(8).String(), nil
}

func (ethw *EthWallet) SendTransaction(fromAccount string, toAddress string, amount string) (string, error) {
	succ, err := ethw.client.Personal.UnlockAccount(fromAccount, ethw.passphrase, uint64(5))
	if err != nil {
		return "", err
	}
	if !succ {
		return "", errors.New("unlock account failure")
	}
	amountDecimal, err := decimal.NewFromString(amount)
	amountWei, _ := big.NewInt(0).SetString(amountDecimal.Mul(decimal.NewFromFloat(math.Pow(10, 18))).String(), 10)
	txId, err := ethw.client.Eth.SendTransaction(&dto.TransactionParameters{
		From:     fromAccount,
		To:       toAddress,
		Gas:      big.NewInt(3000000),
		GasPrice: big.NewInt(18000000000),
		Value:    amountWei,
	})
	return txId, err
}

func (bw *EthWallet) Shutdown() {
}

func (bw *EthWallet) WaitForShutdown() {
}
