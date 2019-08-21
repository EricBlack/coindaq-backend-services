package model

import (
	"bx.com/user-service/bxgo"
	"errors"
	"time"
)

type UserWallet struct {
	Id            int64
	UserId        int64     `xorm:"user_id bigint notnull"`
	Currency      string    `xorm:"currency text"`
	WalletAddress string    `xorm:"wallet_address text"`
	CreateTime    time.Time `xorm:"create_time datetime"`
	UpdateTime    time.Time `xorm:"update_time datetime"`
	IsDeleted     int32     `xorm:"is_deleted int"`
}

func (uw UserWallet) TableName() string {
	return "user_wallet"
}

func (uw *UserWallet) AddWalletAddress() error {
	uw.CreateTime = time.Now()
	uw.UpdateTime = time.Now()
	uw.IsDeleted = False

	wallet, err := uw.QueryWalletByAddress(uw.WalletAddress)
	if err != nil {
		return err
	}
	if wallet.Id != 0 {
		return errors.New("Wallet path existed.")
	}

	_, err = bxgo.OrmEngin.Insert(uw)

	return err
}

func (uw *UserWallet) QueryWalletByAddress(address string) (*UserWallet, error) {
	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	_, err := session.Where("wallet_address=? ", address).Get(uw)

	return uw, err
}

func QueryUserWalletAddress(userId int64, currency, address string) ([]*UserWallet, error) {
	var userWallets []*UserWallet
	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	session.Where("user_id=? ", userId)
	session.Where("is_deleted=? ", False)
	if currency != "" {
		session.Where("currency=? ", currency)
	}
	if address != "" {
		session.Where("wallet_address=? ", address)
	}

	err := session.Find(&userWallets)
	return userWallets, err
}

func DeleteUserWallet(id int64) error {
	_, err := bxgo.OrmEngin.Id(id).Cols("is_deleted", "update_time").
		Update(&UserWallet{
			IsDeleted:  True,
			UpdateTime: time.Now(),
		})

	return err
}
