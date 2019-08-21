package node

import (
	"fmt"
	"testing"
)

var ethTestConfig = map[string]interface{}{
	"schema": "http",
	"host":   "localhost:8545",
	"contracts": []map[string]interface{}{
		{
			"name":              "bixin",
			"constract_address": "0x968Fe2D3bb26E6f7D64351Fe168767B8E7213517",
			"decimals":          18,
			"symbol":            "BXT",
		},
	},
}

func TestETHNewAddress(t *testing.T) {
	ethWallet := NewEthWallet("test_eth", ethTestConfig)
	address, err := ethWallet.GetNewAddress("", "test")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("eth address:" + address)
}

func TestGetBalanceByAddressMin(t *testing.T) {
	ethWallet := NewEthWallet("test_eth", ethTestConfig)
	account, err := ethWallet.GetBalanceByAccount("0x9cc61d218442963f38d656c6091b3b382dfe15b7", 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("eth balance:" + account)
}

func TestUnlockAccount(t *testing.T) {
	ethWallet := NewEthWallet("test_eth", ethTestConfig)
	err := ethWallet.UnlockAccount("0x9cc61d218442963f38d656c6091b3b382dfe15b7", "test", 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("unlock success")
}

func TestGetBlockNum(t *testing.T) {
	ethWallet := NewEthWallet("test_eth", ethTestConfig)
	num, err := ethWallet.GetBlockNum()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("eth block num is %d \n", num)
}

func TestSendTransaction(t *testing.T) {
	ethWallet := NewEthWallet("test_eth", ethTestConfig)
	err := ethWallet.UnlockAccount("0x9cc61d218442963f38d656c6091b3b382dfe15b7", "test", 2)
	txId, err := ethWallet.SendTransaction("0x9cc61d218442963f38d656c6091b3b382dfe15b7", "0x7E46b6d43c8E2a063Da460232AB88F6A8fC1Ab4b", "0.2")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("eth txid is %v \n", txId)
}

func TestGetTokenBalance(t *testing.T) {
	ethWallet := NewEthWallet("test_eth", ethTestConfig)
	amount, err := ethWallet.contracts["BXT"].GetBalanceByAddressMin("0x46892325277c495cd2e7489fb30b5514697abce2", 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("BXT balance is %v \n", amount)
}

func TestSendTokenTransaction(t *testing.T) {
	ethWallet := NewEthWallet("test_eth", ethTestConfig)
	err := ethWallet.UnlockAccount("0x9cc61d218442963f38d656c6091b3b382dfe15b7", "test", 2)
	txId, err := ethWallet.contracts["BXT"].SendTransaction("0x9cc61d218442963f38d656c6091b3b382dfe15b7", "0x7E46b6d43c8E2a063Da460232AB88F6A8fC1Ab4b", "555.5")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("BXT transfer txid is %v \n", txId)
}
