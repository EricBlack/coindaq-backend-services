package wallet

import (
	"fmt"
	"testing"
)

var ercTestConfig = map[string]interface{}{
	"name":              "bixin",
	"constract_address": "0x968Fe2D3bb26E6f7D64351Fe168767B8E7213517",
	"decimals":          18,
	"symbol":            "BXT",
}

func TestGetTokenBalance(t *testing.T) {
	ethWallet := newEthWallet(ethTestConfig)
	erc20Wallet := newERCToken(ethWallet, ERC20TokenABI, ercTestConfig)
	amount, err := erc20Wallet.GetBalance("0x46892325277c495cd2e7489fb30b5514697abce2", 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("BXT balance is %v \n", amount)
}

func TestSendTokenTransaction(t *testing.T) {
	ethWallet := newEthWallet(ethTestConfig)
	erc20Wallet := newERCToken(ethWallet, ERC20TokenABI, ercTestConfig)
	txId, err := erc20Wallet.SendTransaction("0x9cc61d218442963f38d656c6091b3b382dfe15b7", "0x7E46b6d43c8E2a063Da460232AB88F6A8fC1Ab4b", "555.5")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("BXT transfer txid is %v \n", txId)
}
