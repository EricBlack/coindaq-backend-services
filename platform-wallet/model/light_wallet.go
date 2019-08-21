package model

import (
	"errors"
	"time"
)

type LightWalletCoin struct {
	Id         int64
	WalletName string    `xorm:"wallet_name text"`
	UserId     int64     `xorm:"user_id bigint"`
	CoinId     int64     `xorm:"coin_id bigint"`
	Address    string    `xorm:"address text"`
	CreatedAt  time.Time `xorm:"created_at datetime"`
	UpdatedAt  time.Time `xorm:"updated_at datetime"`
}

func (lw LightWalletCoin) TableName() string {
	return "light_wallet"
}

func (lwc *LightWalletCoin) validateField() error {
	if lwc.WalletName == "" {
		return errors.New("wallet name can not be empty")
	}
	if lwc.UserId == 0 {
		return errors.New("user id can not be empty")
	}
	if lwc.CoinId == 0 {
		return errors.New("coin id can not be empty")
	}
	if lwc.Address == "" {
		return errors.New("address can not be empty")
	}
	return nil
}

func CreateLightWalletCoin(lwc *LightWalletCoin) error {
	err := lwc.validateField()
	if err != nil {
		return err
	}
	lwc.CreatedAt = time.Now()
	lwc.UpdatedAt = time.Now()
	exists, err := GetLightWalletCoin(lwc.UserId, lwc.WalletName, lwc.CoinId)
	if err != nil {
		return err
	}
	if exists != nil {
		lwc.Id = exists.Id
		return nil
	}
	_, err = ormEngin.Insert(lwc)
	return err
}

func GetLightWalletCoin(userId int64, walletName string, coinId int64) (*LightWalletCoin, error) {
	var lwc LightWalletCoin
	session := ormEngin.NewSession()
	defer session.Close()

	session.Where("user_id = ?", userId)
	session.And("wallet_name = ?", walletName)
	session.And("coin_id = ?", coinId)
	has, err := session.Get(&lwc)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &lwc, nil
}

func ListLWalletCoinByUser(userId int64) ([]*LightWalletCoin, error) {
	var coins []*LightWalletCoin
	err := ormEngin.Where("user_id = ?", userId).Find(&coins)
	if err != nil {
		return nil, err
	}
	return coins, nil
}

func ListLWalletCoinByUserAndWallet(userId int64, walletName string) ([]*LightWalletCoin, error) {
	var coins []*LightWalletCoin
	err := ormEngin.Where("user_id = ?", userId).And("wallet_name = ?", walletName).Find(&coins)
	if err != nil {
		return nil, err
	}
	return coins, nil
}
