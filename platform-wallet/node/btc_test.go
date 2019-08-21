package node

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
)

var btcdTestconfig map[string]interface{}
var walletTestconfig map[string]interface{}

func init() {
	certs, err := ioutil.ReadFile(filepath.Join("/Users/chenminxiang/Library/Application Support/Btcd", "rpc.cert"))
	if err != nil {
		log.Fatal(err)
	}
	btcdTestconfig = map[string]interface{}{
		"rpc_host": "localhost:18334",
		"user":     "bitcoinrpc",
		"password": "TestBitCoin001",
		"certs":    certs,
	}
	certs, err = ioutil.ReadFile(filepath.Join("/Users/chenminxiang/Library/Application Support/Btcwallet", "rpc.cert"))
	if err != nil {
		log.Fatal(err)
	}
	walletTestconfig = map[string]interface{}{
		"rpc_host": "localhost:18332",
		"user":     "bitcoinrpc",
		"password": "TestBitCoin001",
		"certs":    certs,
	}
}

func TestCreateNewAccount(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	err := btcNode.Unlock("bxkj123456", 5)
	if err != nil {
		t.Error(err)
	}
	err = btcNode.walletClient.CreateNewAccount("test01")
	if err != nil {
		t.Error(err)
	}
	accounts, err := btcNode.ListAccount()
	for k, _ := range accounts {
		if k == "test" {
			fmt.Printf("account [test] created")
			return
		}
	}
	t.Error("account [test] not create")
}

func TestBTCListAccount(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	accounts, err := btcNode.ListAccount()
	if err != nil {
		t.Error(err)
	}
	if len(accounts) <= 0 {
		t.Error("account len never be lte 0")
	}
	fmt.Println(accounts)
}

func TestBTCGetBalanceByAccount(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	balance, err := btcNode.GetBalanceByAccount("test")
	if err != nil {
		t.Error(err)
	}
	if len(balance) < 0 {
		t.Error("balance never be lt 0")
	}
	fmt.Println(balance)
}

func TestBTCGetBestBlock(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	hash, err := btcNode.GetBestBlockHash()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hash)
}

func TestBTCGetAddressesByAccount(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	addresses, err := btcNode.GetAddressesByAccount("")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(addresses)
}

func TestBTCSendToAddress(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	addr, err := btcNode.GetAddressesByAccount("test")
	if err != nil {
		t.Error(err)
	}
	btcNode.Unlock("bxkj123456", 10)
	hash, err := btcNode.SendToAddress(addr[0], "300")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hash)
}

func TestBTCSendFromAccount(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	btcNode.Unlock("bxkj123456", 10)
	hash, err := btcNode.SendFromAccount("test", "mutZbGnoxYLhR2tmm26NjU63i5uxXgjwVj", "0.5")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hash)
}

func TestBTCGetNewAddress(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	addr, err := btcNode.GetNewAddress("test01")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(addr)
}

func TestBTCSendMany(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	txid, err := btcNode.SendMany("test", map[string]string{
		"2N88wYCwGJSScpqixgDY8MD81GahbvC4uMW": "20",
		"2N9E2HdomCaMEejxCkbejzvmGFShQHdgDGL": "20",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(txid)
}

func TestBTCSendRawTransaction(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	tx, err := btcNode.CreateRawTransaction("n3jj4XEuPDGiRFawyA8ZdHSe9oWhTpbJck", "mutZbGnoxYLhR2tmm26NjU63i5uxXgjwVj", "0.1")
	if err != nil {
		log.Fatal(err)
	}
	signedTx, completed, err := btcNode.signRawTransaction(tx)
	if err != nil {
		log.Fatal(err)
	}
	txID, err := btcNode.SendRawTransaction(signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(txID, ",", completed)
}

func TestGetUnspentMinMaxAddress(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	unspents, err := btcNode.GetUnspentMinMaxAddress(1, 999, "n3jj4XEuPDGiRFawyA8ZdHSe9oWhTpbJck")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unspents)
}

func TestGetBTCBalanceByAddressMin(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	txid, err := btcNode.GetBalanceByAddressMin("n3jj4XEuPDGiRFawyA8ZdHSe9oWhTpbJck", 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(txid)
}

func TestGetTransaction(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	tx, err := btcNode.GetTransaction("27a4ed66391cc18143f2f8bd0290b8726709f072963b34d57bce480188f5baad")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx)
}

func TestGetBalanceByAccount(t *testing.T) {
	btcNode := NewBTCNode(btcdTestconfig, walletTestconfig)
	defer btcNode.Shutdown()
	tx, err := btcNode.GetBalanceByAccount("test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx)
}
