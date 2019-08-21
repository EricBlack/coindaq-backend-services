package node

import (
	"strings"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

type BTCNode struct {
	// CoinInfo     CoinInfo
	client *rpcclient.Client
}

func NewBTCNode(btcdConfig map[string]interface{}, walletConfig map[string]interface{}) BTCNode {
	btcdConf := &rpcclient.ConnConfig{
		Host:         btcdConfig["rpc_host"].(string),
		User:         btcdConfig["user"].(string),
		Pass:         btcdConfig["password"].(string),
		Endpoint:     "ws",
		Certificates: btcdConfig["certs"].([]byte),
	}
	btcdNtfnHandlers := NewBTCDNotificationHandlers()
	c, err := rpcclient.New(btcdConf, &btcdNtfnHandlers)
	if err != nil {
		panic(err)
	}
	node := BTCNode{
		client: c,
	}
	return node
}

func (n *BTCNode) GetBestBlockHash() (string, error) {
	hash, err := n.client.GetBestBlockHash()
	if err != nil {
		return "", err
	}
	return hash.String(), nil
}

func (n *BTCNode) SendHexRawTransaction(hex string, allowHighFees bool) (string, error) {
	tx := wire.MsgTx{}
	tx.Deserialize(strings.NewReader(hex))
	txID, err := n.client.SendRawTransaction(&tx, allowHighFees)
	if err != nil {
		return "", err
	}
	return txID.String(), nil
}

func (n *BTCNode) GetRawTransaction(txid string) (*btcutil.Tx, error) {
	hash, err := chainhash.NewHashFromStr(txid)
	if err != nil {
		return nil, err
	}
	return n.client.GetRawTransaction(hash)
}

func (n *BTCNode) Shutdown() {
	n.client.Shutdown()
}

func (n *BTCNode) WaitForShutdown() {
	n.client.WaitForShutdown()
}
