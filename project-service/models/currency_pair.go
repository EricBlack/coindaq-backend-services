package models

import (
	"time"
	"bx.com/project-service/bxgo"
	"github.com/pkg/errors"
)

type CurrencyPair struct {
	Id           int64     `xorm:"pk autoincr BIGINT(20)"`
	CoinaId      string    `xorm:"comment('交易币A') VARCHAR(20)"`
	CoinaName    string    `xorm:"VARCHAR(20)"`
	CoinbId      string    `xorm:"comment('交易币B') VARCHAR(20)"`
	CoinbName    string    `xorm:"VARCHAR(20)"`
	PrioritySort int       `xorm:"INT(11)"`
	Type         int       `xorm:"comment('是否法币还是币币 1:法币, 2:币币') INT(11)"`
	CreateTime   time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}

func QueryCurrencyPairById(id int64) (*CurrencyPair, error) {
	var currencyPair CurrencyPair
	if id == 0 {
		return &currencyPair, errors.New("Currency pair id should not be zero.")
	}

	_, err := bxgo.OrmEngin.Id(id).Get(&currencyPair)

	return &currencyPair, err
}

func QueryCurrencyPairByFilter(coinaName, coinbName string, typeInfo, matchAll int) ([]*CurrencyPair, error) {
	var currencyPaiars []*CurrencyPair

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	if typeInfo != 0 {
		session.Where("type=? ", typeInfo)
	}
	if matchAll == 0 {
		if coinaName != "" {
			session.Where("coina_name=? ", coinaName)
		}
		if coinbName != "" {
			session.Where("coinb_name=? ", coinbName)
		}
	} else {
		if coinaName != ""{
			session.Where("coina_name=? or coinb_name=? ", coinaName, coinaName)
		}

		if coinbName != "" {
			session.Where("coina_name=? or coinb_name=? ", coinaName, coinaName)
		}
	}

	err := session.Asc("priority_sort").Find(&currencyPaiars)

	return currencyPaiars, err
}