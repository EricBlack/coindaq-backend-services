package wallet

import (
	"context"
	"io"

	"bx.com/platform-wallet/model"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	pb "github.com/btcsuite/btcwallet/rpc/walletrpc"
	log "github.com/sirupsen/logrus"
)

type WalletNotificationHandlers struct {
	super *BTCWallet
	s     chan struct{}
}

func (wnf *WalletNotificationHandlers) handleTxNtf() {
	bindedC := pb.NewWalletServiceClient(wnf.super.grpcClient)
	c, err := bindedC.TransactionNotifications(context.Background(), &pb.TransactionNotificationsRequest{})
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			select {
			case <-wnf.s:
				c.CloseSend()
				return
			default:
				in, err := c.Recv()
				if err == io.EOF {
					// read done.
					close(wnf.s)
					return
				}
				if err != nil {
					log.Fatalf("Failed to receive a note : %v", err)
				}
				updateReciveTxCache(wnf.super.client, in)
			}
		}
	}()
}

func updateReciveTxCache(w *rpcclient.Client, in *pb.TransactionNotificationsResponse) {
	//unmined
	for _, t := range in.UnminedTransactions {
		txHash, err := chainhash.NewHash(t.Hash)
		if err != nil {
			log.Error(err)
		}
		tx, err := w.GetTransaction(txHash)
		if err != nil {
			log.Error(err)
		}
		for _, detail := range tx.Details {
			if detail.Category == "receive" && detail.Account != "" {
				txCache := model.TxCache{
					Id:         tx.TxID,
					CoinSymbol: "BTC",
					Account:    detail.Account,
					Address:    detail.Address,
					Kind:       detail.Category,
					Fee:        tx.Fee,
					Amount:     tx.Amount,
				}
				model.CreateUnmiedTxCache(&txCache)
			}
		}
	}
	//mined
	for _, block := range in.AttachedBlocks {
		for _, t := range block.Transactions {
			txHash, err := chainhash.NewHash(t.Hash)
			if err != nil {
				log.Error(err)
			}
			tx, err := w.GetTransaction(txHash)
			if err != nil {
				log.Error(err)
			}
			for _, detail := range tx.Details {
				if detail.Category == "receive" && detail.Account != "" {
					cHash, err := chainhash.NewHash(block.Hash)
					if err != nil {
						log.Error(err)
					}
					err = model.InsertMinedTxCache(tx.TxID, cHash.String(), block.Height, tx.BlockIndex)
					if err != nil {
						log.Error(err)
					}
				}
			}
		}
		coin, err := model.GetCoinInfoBySymbol("BTC")
		if err != nil {
			log.Error(err)
		}
		filter := model.TxCacheFilter{
			Symbol:   "BTC",
			MaxConf:  int32(coin.MinConfirms),
			Kind:     "receive",
			Finished: false,
			FromTime: nil,
		}
		txList, err := model.ListTx(&filter)
		for _, minedTx := range txList {
			err = minedTx.IncConf()
			if err != nil {
				log.Error(err)
			}
			if minedTx.Confirmations == int32(coin.MinConfirms) {
				err := minedTx.Finish()
				if err != nil {
					log.Error(err)
				}
			}
		}
	}
}

// func (wnf *WalletNotificationHandlers) getTx(hash *chainhash.Hash) (*wire.MsgTx, error) {
//
// 	txByte, err := hex.DecodeString(txRes.Hex)
// 	if err != nil {
// 		return nil, err
// 	}
// 	tx, err := btcutil.NewTxFromBytes(txByte)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return tx.MsgTx(), nil
// }
//
// func (wnf *WalletNotificationHandlers) getAccountByAddr(addr btcutil.Address) (string, error) {
// 	account, err := wnf.super.client.GetAccount(addr)
// 	if err != nil {
// 		return "", err
// 	}
// 	return account, nil
// }
//
// func (wnf *WalletNotificationHandlers) lookupInputAccount(details *wtxmgr.TxDetails, deb wtxmgr.DebitRecord) string {
// 	// TODO: Debits should record which account(s?) they
// 	// debit from so this doesn't need to be looked up.
// 	prevOP := &details.MsgTx.TxIn[deb.Index].PreviousOutPoint
// 	prev, err := wnf.getTx(&prevOP.Hash)
// 	if err != nil {
// 		log.Errorf("Cannot query previous transaction details for %v: %v", prevOP.Hash, err)
// 		return ""
// 	}
// 	if prev == nil {
// 		log.Errorf("Missing previous transaction %v", prevOP.Hash)
// 		return ""
// 	}
// 	prevOut := prev.TxOut[prevOP.Index]
// 	_, addrs, _, err := txscript.ExtractPkScriptAddrs(prevOut.PkScript, &chaincfg.TestNet3Params)
// 	var inputAcct string
// 	if err == nil && len(addrs) > 0 {
// 		inputAcct, err = wnf.getAccountByAddr(addrs[0])
// 	}
// 	if err != nil {
// 		log.Errorf("Cannot fetch account for previous output %v: %v", prevOP, err)
// 		inputAcct = ""
// 	}
// 	return inputAcct
// }
//
// func (wnf *WalletNotificationHandlers) lookupOutputChain(details *wtxmgr.TxDetails, cred wtxmgr.CreditRecord) (account string) {
// 	output := details.MsgTx.TxOut[cred.Index]
// 	_, addrs, _, err := txscript.ExtractPkScriptAddrs(output.PkScript, &chaincfg.TestNet3Params)
// 	var acct string
// 	if err == nil && len(addrs) > 0 {
// 		acct, err = wnf.getAccountByAddr(addrs[0])
// 	}
// 	if err != nil {
// 		log.Errorf("Cannot fetch account for wallet output: %v", err)
// 	} else {
// 		account = acct
// 	}
// 	return
// }

func (wnf *WalletNotificationHandlers) run() {
	wnf.handleTxNtf()
}

func (wnf *WalletNotificationHandlers) stop() {
	close(wnf.s)
}
