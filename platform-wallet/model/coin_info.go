package model

import (
	"errors"
	"time"
)

const (
	_ = iota
	BTC
	ETH
	ERC20
)

const (
	CoinWithdrawDisable = iota
	CoinWithdrawEnable
)

const (
	CoinReceiveDisable = iota
	CoinReceiveEnable
)

type CoinInfo struct {
	Id          int64
	Code        string    `xorm:"code text"`
	Symbol      string    `xorm:"symbol text"`
	Name        string    `xorm:"name text"`
	Decimals    int       `xorm:"decimals integer"`
	Kind        int       `xorm:"coin_type integer"`
	MinConfirms int       `xorm:"min_confirms integer"`
	Withdrawal  int       `xorm:"withdrawal integer"`
	Receivable  int       `xorm:"receivable integer"`
	Path        string    `xorm:"path text"`
	DeletedAt   time.Time `xorm:"deleted_at datetime"`
	CreatedAt   time.Time `xorm:"created_at datetime"`
	UpdatedAt   time.Time `xorm:"updated_at datetime"`
}

func (ci CoinInfo) TableName() string {
	return "coin_info"
}

func (ci *CoinInfo) ValidateFields() error {
	if ci.Code == "" {
		return errors.New("code field can not be nil or empty")
	}
	if ci.Symbol == "" {
		return errors.New("symbol field can not be nil or empty")
	}
	if ci.Name == "" {
		return errors.New("name field can not be nil or empty")
	}
	if ci.Decimals <= 0 || ci.Decimals > 18 {
		return errors.New("decimals field out of range, [0, 18)")
	}
	if !(ci.Kind == BTC || ci.Kind == ETH || ci.Kind == ERC20) {
		return errors.New("unknow coin info kind, enum(BTC, ETH, ERC20)")
	}
	return nil
}

func CreateCoinInfo(coinInfo *CoinInfo) error {
	exists := CoinInfo{}
	session := ormEngin.NewSession()
	defer session.Close()
	session.Where("code = ?", coinInfo.Code)
	session.Or("symbol = ?", coinInfo.Symbol)
	session.And("deleted_at IS NULL")
	has, err := session.Exist(&exists)
	if err != nil {
		return err
	}
	if has {
		return errors.New("duplicated value in filed code, symbol")
	}
	if err != nil {
		return err
	}
	coinInfo.Withdrawal = CoinWithdrawDisable
	coinInfo.Receivable = CoinReceiveDisable
	coinInfo.CreatedAt = time.Now()
	coinInfo.UpdatedAt = time.Now()
	_, err = ormEngin.Insert(coinInfo)
	if err != nil {
		return err
	}
	return nil
}

func ListCoinInfo(kind int) ([]*CoinInfo, error) {
	var coinInfoList []*CoinInfo
	session := ormEngin.NewSession()
	defer session.Close()
	session.Where("deleted_at IS NULL")
	if kind != 0 {
		session.And("coin_type = ?", kind)
	}
	err := session.Find(&coinInfoList)
	return coinInfoList, err
}

func GetCoinInfo(id int64) (*CoinInfo, error) {
	var coinInfo CoinInfo
	has, err := ormEngin.ID(id).Get(&coinInfo)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("coin info not found")
	}
	return &coinInfo, nil
}

func GetCoinInfoBySymbol(symbol string) (*CoinInfo, error) {
	var coinInfo CoinInfo
	has, err := ormEngin.Where("symbol = ?", symbol).And("deleted_at IS NULL").Get(&coinInfo)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("coin info not found")
	}
	return &coinInfo, nil
}

func (ci *CoinInfo) Save() error {
	updatingCoinInfo := CoinInfo{
		Id:          ci.Id,
		Code:        ci.Code,
		Symbol:      ci.Symbol,
		Name:        ci.Symbol,
		Decimals:    ci.Decimals,
		Kind:        ci.Kind,
		MinConfirms: ci.MinConfirms,
		UpdatedAt:   time.Now(),
	}
	if updatingCoinInfo.Id == 0 {
		return errors.New("no id, call create first")
	}
	exists := CoinInfo{}
	session := ormEngin.NewSession()
	defer session.Close()
	session.Where("id <> ?", updatingCoinInfo.Id)
	session.And("deleted_at IS NULL")
	session.And("(code = ? OR symbol = ?)", updatingCoinInfo.Code, updatingCoinInfo.Symbol)
	has, err := session.Exist(&exists)
	if err != nil {
		return err
	}
	if has {
		return errors.New("duplicated value in filed code, symbol")
	}

	_, err = ormEngin.Id(updatingCoinInfo.Id).Omit("withdrawal", "receivable", "created_at", "deleted_at").Update(&updatingCoinInfo)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCoinInfo(id int64) error {
	_, err := ormEngin.ID(id).Cols("deleted_at").Update(&CoinInfo{DeletedAt: time.Now()})
	return err
}

func EnableWithdraw(id int64) error {
	_, err := ormEngin.ID(id).Cols("withdrawal").Update(&CoinInfo{Withdrawal: CoinWithdrawEnable})
	return err
}

func DisableWithdraw(id int64) error {
	_, err := ormEngin.ID(id).Cols("withdrawal").Update(&CoinInfo{Withdrawal: CoinWithdrawDisable})
	return err
}

func EnableReceive(id int64) error {
	_, err := ormEngin.ID(id).Cols("receivable").Update(&CoinInfo{Receivable: CoinReceiveEnable})
	return err
}

func DisableReceive(id int64) error {
	_, err := ormEngin.ID(id).Cols("receivable").Update(&CoinInfo{Receivable: CoinReceiveDisable})
	return err
}
