package model

import (
	"bx.com/user-service/bxgo"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type UserBalance struct {
	BalanceId          string    `xorm:"not null pk comment('余额主键') VARCHAR(32)"`
	UserId             string    `xorm:"not null VARCHAR(32)"`
	CurrencyId         string    `xorm:"not null comment('币表主键') VARCHAR(32)"`
	BalanceValue       int64     `xorm:"comment('余额数字 1表示 0.00000001') BIGINT(50)"`
	RechargeAddress    string    `xorm:"VARCHAR(100)"`
	QrcodeAddress      string    `xorm:"comment('二维码地址') VARCHAR(100)"`
	TotalBalance       int64     `xorm:"comment('账户总额') BIGINT(20)"`
	ChargeUnAccount    int64     `xorm:"comment('充值未到账') BIGINT(20)"`
	WithdrawUnTransfer int64     `xorm:"comment('提现未转出') BIGINT(20)"`
	LockPosition       int64     `xorm:"comment('锁仓不可卖') BIGINT(20)"`
	IcoUndue           int64     `xorm:"comment('ico未到期') BIGINT(20)"`
	UpdateTime         time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

func (ub UserBalance) TableName() string {
	return "user_balance"
}

func QueryBalanceByUserId(userId int64) ([]*UserBalance, error) {
	var userList []*UserBalance
	userIdInfo := fmt.Sprintf("%d", userId)
	_, err := bxgo.OrmEngin.Where("user_id=? ", userIdInfo).Desc("currency_id").Get(userList)

	return userList, err
}

func QueryBalanceByFilter(userId int64, currencyId string) ([]*UserBalance, error) {
	var userBalance []*UserBalance
	userIdInfo := fmt.Sprintf("%d", userId)

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	session.Where("user_id=? ", userIdInfo)
	if currencyId != "" {
		session.Where("currency_id=? ", currencyId)
	}

	if err := session.Find(&userBalance); err != nil {
		return userBalance, err
	}

	return userBalance, nil
}

func CheckCoinAddressExist(userId int64, currencyId string) (bool, error) {
	userIdInfo := fmt.Sprintf("%d", userId)
	userBalance := UserBalance{}
	_, err := bxgo.OrmEngin.Where("user_id=?", userIdInfo).
		Where("currency_id=? ", currencyId).Get(&userBalance)
	if err != nil {
		return false, err
	}else {
		if userBalance.BalanceId == ""{
			return false, nil
		}else {
			return true, nil
		}
	}
}

func CreateAllUserCoinAddress(userId int64)  {
	currencyList, err := GetAllCurrency()
	if err != nil || len(currencyList) == 0 {
		return
	} else {
		for _, item := range currencyList {
			flag, err := CheckCoinAddressExist(userId, item.CurrencyId)
			if err != nil || flag {
				continue
			} else {
				_ = CreateUserCoinAddress(userId, item.CurrencyId)
			}
		}
	}
}

func CreateUserCoinAddress(userId int64, currencyId string) error {
	var tableName string
	userIdInfo := fmt.Sprintf("%d", userId)

	//区分钱包生成地址
	switch currencyId {
	case "100001":
		tableName = "receive_btc"
		break
	case "100004":
		tableName = "receive_bch"
	default:
		tableName = "receive_eth"
	}

	//查询是否已存在
	isExist, err := CheckCoinAddressExist(userId, currencyId)
	if err != nil || isExist {
		return err
	}

	//查询地址
	var sqlString string = ""
	sqlString = fmt.Sprintf("select receive_address from %s where receive_user_id is null limit 1", tableName)
	results, err := bxgo.OrmEngin.QueryString(sqlString)
	if err != nil {
		return err
	}

	if len(results) == 1 {
		address := results[0]["receive_address"]
		//事物保存
		session := bxgo.OrmEngin.NewSession()
		defer session.Close()

		err = session.Begin()
		//保存Receive Table
		sqlString = fmt.Sprintf("update %s set receive_user_id='%s' where receive_address='%s'", tableName, userIdInfo, address)
		_, err = session.Exec(sqlString)
		if err != nil {
			session.Rollback()
			return err
		}

		//保存UserBalance Table
		//uid, _ := uuid.NewV4()
		uid := GetGuid()
		userBalance := UserBalance{
			BalanceId:       fmt.Sprintf("%s", uid),
			UserId:          userIdInfo,
			CurrencyId:      currencyId,
			RechargeAddress: address,
		}
		_, err = session.Insert(&userBalance)
		if err != nil {
			session.Rollback()
			return err
		}

		// add Commit() after all actions
		if err = session.Commit(); err != nil {
			return err
		}
		return nil

	} else {
		return errors.New("No available address to allocation.")
	}
}

func GetGuid() string {
	timeString := time.Now().String()
	h := md5.New()
	h.Write([]byte(timeString))
	return hex.EncodeToString(h.Sum(nil))
}
