package wallet

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"
	"time"
)

var walletTestconfig map[string]interface{}

func init() {
	certPath := filepath.Join("/Users/chenminxiang/Library/Application Support/Btcwallet", "rpc.cert")
	walletTestconfig = map[string]interface{}{
		"rpc_host":   "localhost:18332",
		"user":       "bitcoinrpc",
		"password":   "TestBitCoin001",
		"cert_file":  certPath,
		"passphrase": "bxkj123456",
	}
}

func TestBTCSendTransaction(t *testing.T) {
	btcWallet := newBTCWallet(walletTestconfig)
	defer btcWallet.Shutdown()
	hash, err := btcWallet.SendTransaction("1Kg6gB8AygO-7XYbWyVFmPF", "ms3vkeDk9XoFt9JH2dd4BPCD6XpWYTAYR8", "0.01")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(hash)
}

func TestBTCGetBalance(t *testing.T) {
	btcWallet := newBTCWallet(walletTestconfig)
	defer btcWallet.Shutdown()
	amount, err := btcWallet.GetBalance("test", 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(amount)
}

func TestBTCGetNewAccount(t *testing.T) {
	btcWallet := newBTCWallet(walletTestconfig)
	defer btcWallet.Shutdown()
	account, address, err := btcWallet.GetNewAccount()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(account)
	fmt.Println(address)
}

func TestBTCTxNotification(t *testing.T) {
	walletTestconfig["enable_ntfc"] = "true"
	btcWallet := newBTCWallet(walletTestconfig)
	// hash, err := btcWallet.SendTransaction("test", "mhHwxT5q76QPVNRjvF7SxTxijLHnxR4GEN", "0.1")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(hash)
	log.Println("Client shutdown in 1800 seconds...")
	time.AfterFunc(time.Second*1800, func() {
		log.Println("Client shutting down...")
		btcWallet.Shutdown()
		log.Println("Client shutdown complete.")
	})

	// Wait until the client either shuts down gracefully (or the user
	// terminates the process with Ctrl+C).
	btcWallet.WaitForShutdown()
}
