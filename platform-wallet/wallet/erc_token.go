package wallet

import (
	"math/big"

	"github.com/johnnycx127/go-web3/dto"
	"github.com/johnnycx127/go-web3/eth"
	"github.com/shopspring/decimal"
)

type ERCToken struct {
	super            *EthWallet
	constractCli     *eth.Contract
	constractAddress string
	symbol           string
	decimals         int
}

func newERCToken(ethWallet *EthWallet, abi string, config map[string]interface{}) *ERCToken {
	contract, err := ethWallet.client.Eth.Contract(abi)
	if err != nil {
		panic(err)
	}
	wallet := ERCToken{
		symbol:           config["symbol"].(string),
		decimals:         config["decimals"].(int),
		constractAddress: config["constract_address"].(string),
		constractCli:     contract,
		super:            ethWallet,
	}
	return &wallet
}

func (ethtw *ERCToken) GetBalance(address string, minConf int) (string, error) {
	res, err := ethtw.constractCli.Call(&dto.TransactionParameters{
		From: address,
		To:   ethtw.constractAddress,
	}, "balanceOf", address)
	if err != nil {
		return "", err
	}
	resHex, err := res.ToComplexIntResponse()
	if err != nil {
		return "", err
	}
	amountDecimal := decimal.NewFromBigInt(resHex.ToBigInt(), 0)
	return decimal.NewFromBigInt(big.NewInt(1), -int32(ethtw.decimals)).Mul(amountDecimal).Round(int32(ethtw.decimals)).String(), nil
}

func (ethtw *ERCToken) SendTransaction(fromAccount string, toAddress string, amount string) (string, error) {
	ethtw.super.client.Personal.UnlockAccount(fromAccount, ethtw.super.passphrase, 5)
	amountDecimal, err := decimal.NewFromString(amount)
	amountWeiStr := amountDecimal.Mul(decimal.NewFromBigInt(big.NewInt(1), int32(ethtw.decimals))).String()
	amountBigInt, _ := big.NewInt(0).SetString(amountWeiStr, 10)
	if err != nil {
		return "", err
	}
	txId, err := ethtw.constractCli.Send(&dto.TransactionParameters{
		From:     fromAccount,
		To:       ethtw.constractAddress,
		Gas:      big.NewInt(3000000),
		GasPrice: big.NewInt(18000000000),
	}, "transfer", toAddress, amountBigInt)
	return txId, err
}

func (ethtw *ERCToken) GetNewAccount() (string, string, error) {
	return "", "", nil
}

func (bw *ERCToken) Shutdown() {
}

func (bw *ERCToken) WaitForShutdown() {
}
