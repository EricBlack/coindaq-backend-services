package node

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

func NewBTCDNotificationHandlers() rpcclient.NotificationHandlers {
	return rpcclient.NotificationHandlers{
		OnFilteredBlockConnected:    blockConnectedHandler,
		OnFilteredBlockDisconnected: blockDisconnectedHandler,
	}
}

func blockConnectedHandler(height int32, header *wire.BlockHeader, txs []*btcutil.Tx) {
}

func blockDisconnectedHandler(height int32, header *wire.BlockHeader) {
}
