package wallet

import (
	"fmt"
	"testing"
)

var ethTestConfig = map[string]interface{}{
	"schema":     "http",
	"host":       "localhost:8545",
	"passphrase": "bxkj123456",
}

func TestETHGetNewAddress(t *testing.T) {
	ethWallet := newEthWallet(ethTestConfig)
	account, address, err := ethWallet.GetNewAccount()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("eth account:%s, ", account)
	fmt.Println("eth address:" + address)
}

func TestETHGetBalance(t *testing.T) {
	ethWallet := newEthWallet(ethTestConfig)
	account, err := ethWallet.GetBalance("0x0fb25efcca91d4bfceb0490b70fcbefb95f6d491", 12)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("eth balance:" + account)
}

func TestETHSendTransaction(t *testing.T) {
	ethWallet := newEthWallet(ethTestConfig)
	txId, err := ethWallet.SendTransaction("0x63b8158ee1da30b37ca0712e5432646d4dbf77e9", "0x0fb25efcca91d4bfceb0490b70fcbefb95f6d491", "0.1")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("eth txid is %v \n", txId)
}
