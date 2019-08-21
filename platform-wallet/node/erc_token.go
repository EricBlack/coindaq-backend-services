package node

import (
	"math/big"

	"github.com/johnnycx127/go-web3/dto"
	"github.com/johnnycx127/go-web3/eth"
	"github.com/shopspring/decimal"
)

type ERCToken struct {
	constractCli     *eth.Contract
	constractAddress string
	symbol           string
	name             string
	decimals         int
}

func (ethtw *ERCToken) GetBalanceByAddressMin(address string, minConf int) (string, error) {
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

func (ethtw *ERCToken) SendTransaction(fromAddress string, toAddress string, amount string) (string, error) {
	amountDecimal, err := decimal.NewFromString(amount)
	amountWeiStr := amountDecimal.Mul(decimal.NewFromBigInt(big.NewInt(1), int32(ethtw.decimals))).String()
	amountBigInt, _ := big.NewInt(0).SetString(amountWeiStr, 10)
	if err != nil {
		return "", err
	}
	txId, err := ethtw.constractCli.Send(&dto.TransactionParameters{
		From:     fromAddress,
		To:       ethtw.constractAddress,
		Gas:      big.NewInt(3000000),
		GasPrice: big.NewInt(18000000000),
	}, "transfer", toAddress, amountBigInt)
	return txId, err
}
