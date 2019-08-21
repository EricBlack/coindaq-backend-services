package model

import (
	"fmt"
	"time"
)

type TxCache struct {
	Id            string    `xorm:"id text pk"`
	CoinId        int64     `xorm:"coin_id bigint"`
	CoinSymbol    string    `xorm:"coin_symbol text"`
	BlockHash     string    `xorm:"block_hash" text`
	Account       string    `xorm:"account text"`
	Address       string    `xorm:"address text"`
	Kind          string    `xorm:"tx_type text"`
	BlockHeight   int32     `xorm:"block_height int"`
	BlockIndex    int64     `xorm:"block_index bigint"`
	Fee           float64   `xorm:"fee numeric"`
	Amount        float64   `xorm:"amount numeric"`
	Confirmations int32     `xorm:"confirmations int"`
	Finished      bool      `xorm:"finished bool"`
	CreatedAt     time.Time `xorm:"created_at datetime"`
	UpdatedAt     time.Time `xorm:"updated_at datetime"`
	FinishedAt    time.Time `xorm:"finished_at datetime"`
}

type TxCacheFilter struct {
	Symbol   string
	MaxConf  int32
	Kind     string
	Finished bool
	FromTime *time.Time
}

func (uw *TxCache) TableName() string {
	return "tx_cache"
}

func CreateUnmiedTxCache(tx *TxCache) error {
	coin, err := GetCoinInfoBySymbol(tx.CoinSymbol)
	if err != nil {
		return err
	}
	tx.Finished = false
	tx.CoinId = coin.Id
	tx.CreatedAt = time.Now()
	tx.UpdatedAt = time.Now()
	_, err = ormEngin.Insert(tx)
	fmt.Println(err)
	return err
}

func GetTxById(id string) (*TxCache, error) {
	var tx *TxCache
	_, err := ormEngin.Id(id).Get(tx)
	return tx, err
}

func ListTx(filter *TxCacheFilter) ([]*TxCache, error) {
	var txList []*TxCache
	session := ormEngin.NewSession()
	defer session.Close()
	session.Where("coin_symbol = ?", filter.Symbol)
	session.And("finished = ?", filter.Finished)
	if filter.MaxConf > 0 {
		session.And("confirmations > 0 AND confirmations < ?", filter.MaxConf)
	}
	if filter.Kind != "" {
		session.And("tx_type = ?", filter.Kind)
	}
	if filter.FromTime != nil {
		session.And("created_at > ?", filter.FromTime)
	}
	err := session.Find(&txList)
	if err != nil {
		return nil, err
	}
	return txList, nil
}

func InsertMinedTxCache(id string, blockHash string, height int32, blockIndex int64) error {
	tx := TxCache{
		BlockHash:     blockHash,
		BlockHeight:   height,
		BlockIndex:    blockIndex,
		Confirmations: 1,
		UpdatedAt:     time.Now(),
	}
	_, err := ormEngin.ID(id).Cols("block_hash", "height", "block_index", "confirmations", "updated_at").Update(&tx)
	return err
}

func (txc *TxCache) IncConf() error {
	txc.Confirmations += 1
	_, err := ormEngin.ID(txc.Id).Cols("confirmations", "updated_at").Update(&TxCache{Confirmations: txc.Confirmations, UpdatedAt: time.Now()})
	return err
}

func (txc *TxCache) Finish() error {
	_, err := ormEngin.ID(txc.Id).Cols("finished", "updated_at", "finished_at").Update(&TxCache{Finished: true, UpdatedAt: time.Now(), FinishedAt: time.Now()})
	return err
}
