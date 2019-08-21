package wallet

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"github.com/chilts/sid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type BTCWallet struct {
	client     *rpcclient.Client
	grpcClient *grpc.ClientConn
	ntfc       *WalletNotificationHandlers
	passphrase string
}

func newBTCWallet(config map[string]interface{}) *BTCWallet {
	certs, err := ioutil.ReadFile(config["cert_file"].(string))
	walletConf := &rpcclient.ConnConfig{
		Host:         config["rpc_host"].(string),
		User:         config["user"].(string),
		Pass:         config["password"].(string),
		Endpoint:     "ws",
		Certificates: certs,
	}
	w, err := rpcclient.New(walletConf, nil)
	if err != nil {
		panic(err)
	}
	passphraseByte := []byte(config["passphrase"].(string))
	sha256Byte := sha256.Sum256(passphraseByte)
	creds, err := credentials.NewClientTLSFromFile(config["cert_file"].(string), "localhost")
	if err != nil {
		panic(err)
	}
	c, err := grpc.Dial("localhost:19332", grpc.WithTransportCredentials(creds))
	node := BTCWallet{
		client:     w,
		passphrase: hex.EncodeToString(sha256Byte[:]),
		grpcClient: c,
	}
	node.ntfc = &WalletNotificationHandlers{super: &node, s: make(chan struct{})}
	if config["enable_ntfc"] != nil && config["enable_ntfc"].(string) == "true" {
		node.ntfc.run()
	}
	return &node
}

func (bw *BTCWallet) GetNewAccount() (string, string, error) {
	bw.client.WalletPassphrase(bw.passphrase, 5)
	accountId := sid.Id()
	err := bw.client.CreateNewAccount(accountId)
	if err != nil {
		return "", "", err
	}
	address, err := bw.client.GetNewAddress(accountId)
	if err != nil {
		return "", "", err
	}
	return accountId, address.String(), nil
}

func (bw *BTCWallet) GetBalance(account string, minConfirms int) (string, error) {
	balance, err := bw.client.GetBalanceMinConf(account, minConfirms)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%f", balance.ToBTC()), nil
}

func (bw *BTCWallet) SendTransaction(fromAccount string, toAddress string, amount string) (string, error) {
	bw.client.WalletPassphrase(bw.passphrase, 5)
	btcAddress, err := btcutil.DecodeAddress(toAddress, &chaincfg.TestNet3Params)
	if err != nil {
		return "", err
	}
	btcAmount, err := getBTCAmount(amount)
	if err != nil {
		return "", err
	}
	txId, err := bw.client.SendFrom(fromAccount, btcAddress, btcAmount)
	return txId.String(), err
}

func (bw *BTCWallet) Shutdown() {
	bw.client.Shutdown()
	bw.ntfc.stop()
	bw.grpcClient.Close()
}

func (bw *BTCWallet) WaitForShutdown() {
	bw.client.WaitForShutdown()
}
