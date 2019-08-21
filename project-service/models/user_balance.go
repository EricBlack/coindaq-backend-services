package models

import (
	"time"
	"strconv"
	"bx.com/project-service/bxgo"
	"github.com/pkg/errors"
	"github.com/go-xorm/xorm"
)

type UserBalance struct {
	BalanceId          string    `xorm:"not null pk comment('余额主键') VARCHAR(32)"`
	UserId             string    `xorm:"not null VARCHAR(32)"`
	CurrencyId         string    `xorm:"not null comment('币表主键') VARCHAR(32)"`
	BalanceValue       int64     `xorm:"comment('余额数字 1表示 0.00000001') BIGINT(50)"`
	RechargeAddress    string    `xorm:"comment('更改地址？') VARCHAR(100)"`
	QrcodeAddress      string    `xorm:"comment('二维码地址') VARCHAR(100)"`
	TotalBalance       int64     `xorm:"comment('账户总额') BIGINT(20)"`
	ChargeUnAccount    int64     `xorm:"comment('充值未到账') BIGINT(20)"`
	WithdrawUnTransfer int64     `xorm:"comment('提现未转出') BIGINT(20)"`
	LockPosition       int64     `xorm:"comment('锁仓不可卖') BIGINT(20)"`
	IcoUndue           int64     `xorm:"comment('ico未到期') BIGINT(20)"`
	UpdateTime         time.Time `xorm:"DATETIME"`
}

func QueryUserBalanceList(userId int64) ([]*UserBalance, error) {
	var balanceList []*UserBalance
	userIdInfo := strconv.FormatInt(userId, 10)
	err := bxgo.OrmEngin.Where("user_id=? ", userIdInfo).Find(&balanceList)

	return balanceList, err
}

func QueryUserBalanceByFilter(userId int64, currencyId string) (*UserBalance, error) {
	var userBalance UserBalance

	_, err := bxgo.OrmEngin.Where("user_id=? ", userId).
		Where("currency_id=? ", currencyId).Get(&userBalance)

	return &userBalance, err
}

func QueryUserProjectBalance(userId, projectId int64) (*UserBalance, error) {
	var userBalance *UserBalance
	projectInfo, err := QueryProjectById(projectId)
	if err != nil {
		return userBalance, err
	}

	userBalance, err = QueryUserBalanceByFilter(userId, projectInfo.IssueCoina)
	return userBalance, err
}

func LockUserBalance(session *xorm.Session, userId, balance int64, currencyId string) error {
	userBalance, err := QueryUserBalanceByFilter(userId, currencyId)
	if err != nil {
		return err
	}
	if userBalance.BalanceValue< balance {
		return errors.New("Lock balance amount over user balance value.")
	}

	userBalance.BalanceValue -= balance
	userBalance.LockPosition += balance
	userBalance.UpdateTime = time.Now()

	_, err = session.Where("balance_id=? ", userBalance.BalanceId).
		Cols("balance_value", "lock_position", "update_time").Update(userBalance)

	return err
}

func UnlockUserBalance(session *xorm.Session, userId, balance int64, currencyId string) error {
	userBalance, err := QueryUserBalanceByFilter(userId, currencyId)
	if err != nil {
		return err
	}
	if userBalance.LockPosition< balance {
		return errors.New("Unlock balance amount over user lock position value.")
	}

	userBalance.BalanceValue += balance
	userBalance.LockPosition -= balance
	userBalance.UpdateTime = time.Now()

	_, err = session.Where("balance_id=? ", userBalance.BalanceId).
		Cols("balance_value", "lock_position", "update_time").Update(userBalance)

	return err
}
