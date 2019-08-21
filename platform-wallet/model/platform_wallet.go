package model

import (
	"errors"
	"time"
)

type PWalletCoin struct {
	Id         int64
	UserId     int64     `xorm:"user_id bigint"`
	CoinId     int64     `xorm:"coin_id int"`
	WalletName string    `xorm:"wallet_name text"`
	Account    string    `xorm:"account text index(user_wallet_account_idx)"`
	Address    string    `xorm:"address text index(user_wallet_address_idx)"`
	CreatedAt  time.Time `xorm:"created_at datetime"`
	UpdatedAt  time.Time `xorm:"updated_at datetime"`
}

func (pwc *PWalletCoin) TableName() string {
	return "platform_wallet"
}

func (pwc *PWalletCoin) validateFields() error {
	if pwc.UserId == 0 {
		return errors.New("user id can not be empty")
	}
	if pwc.CoinId == 0 {
		return errors.New("coin id can not be empty")
	}
	if pwc.WalletName == "" {
		return errors.New("wallet name can not be empty")
	}
	if pwc.Account == "" {
		return errors.New("account can not be empty")
	}
	if pwc.Address == "" {
		return errors.New("address can not be empty")
	}
	return nil
}

func CreatePWalletCoin(pwc *PWalletCoin) error {
	if err := pwc.validateFields(); err != nil {
		return err
	}
	exists, err := GetPWalletCoin(pwc.UserId, pwc.WalletName, pwc.CoinId)
	if err != nil {
		return err
	}
	if exists != nil {
		return errors.New("coin exists")
	}
	pwc.CreatedAt = time.Now()
	pwc.UpdatedAt = time.Now()
	_, err = ormEngin.Insert(&pwc)
	if err != nil {
		return err
	}
	return nil
}

func ListPWCByUser(uid int64) ([]*PWalletCoin, error) {
	var pwcList []*PWalletCoin
	err := ormEngin.Where("user_id = ?", uid).Find(&pwcList)
	return pwcList, err
}

func GetPWalletCoin(uid int64, walletName string, coinId int64) (*PWalletCoin, error) {
	var pwc PWalletCoin
	session := ormEngin.NewSession()
	defer session.Close()

	session.Where("user_id = ?", uid)
	session.And("wallet_name = ?", walletName)
	session.And("coin_id = ?", coinId)
	has, err := session.Get(&pwc)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &pwc, nil
}

func (pwc *PWalletCoin) UpdatePWCAccountAndAddr(newAccount string, newAddr string) error {
	pwc.Account = newAccount
	pwc.Address = newAddr
	pwc.UpdatedAt = time.Now()
	if _, err := ormEngin.ID(pwc.Id).Cols("account", "address", "updated_at").Update(&pwc); err != nil {
		return err
	}
	return nil
}
