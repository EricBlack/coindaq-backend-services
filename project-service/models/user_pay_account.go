package models

import (
	"time"
	"bx.com/project-service/bxgo"
	"github.com/pkg/errors"
	"bx.com/user-service/model"
)

type UserPayAccount struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	UserId        int64     `xorm:"comment('用户id') BIGINT(20)"`
	AccountName   string    `xorm:"comment('收款人名称') VARCHAR(255)"`
	AccountType   int       `xorm:"comment('账号类型') INT(11)"`
	AccountNumber string    `xorm:"VARCHAR(20)"`
	OpenBank      string    `xorm:"comment('开户银行') VARCHAR(255)"`
	SubBank       string    `xorm:"comment('开户支行') VARCHAR(255)"`
	IsDeleted     int       `xorm:"INT(11)"`
	CreateTime    time.Time `xorm:"DATETIME"`
	UpdateTime    time.Time `xorm:"DATETIME"`
}

func CheclPayAccountParameter(account UserPayAccount) error {
	if account.UserId == 0 {
		return errors.New("User id parameter not correct.")
	}
	if account.AccountType ==0 {
		return errors.New("Pay account type is not correct.")
	}
	if account.AccountType == 1 {
		if account.SubBank == "" || account.OpenBank == "" || account.AccountName == ""{
			return errors.New("AccountName, OpenBank and SubBank should not be empty.")
		}
 	}

 	return nil
}

func AddPayAccount(userId int64, accountName, accountNumber, openBank, subBank, paymentPwd string, accountType int) (error) {
	//判断类型是否重复添加
	payAccount, err := QueryUserPayAccountByFilter(userId, accountType)
	if err != nil {
		return err
	}
	if len(payAccount) >0 {
		return errors.New("Cannot add duplicate type of pay account address.")
	}
	newPayAccount := UserPayAccount{
		UserId:			userId,
		AccountName:	accountName,
		AccountType:	accountType,
		AccountNumber:	accountNumber,
		OpenBank:		openBank,
		SubBank:		subBank,
	}
	//检查参数
	if err = CheclPayAccountParameter(newPayAccount); err != nil {
		return err
	}

	//验证支付密码
	result, err := model.VerifyPaymentPassword(userId, paymentPwd)
	if err != nil {
		return err
	}
	if !result{
		return errors.New("User payment password not correct.")
	}

	//添加
	_, err = bxgo.OrmEngin.Insert(&newPayAccount)
	return err
}

func DeletePayAccount(userId int64, accountType int) (error) {
	//判断类型是否存在
	payAccounts, err := QueryUserPayAccountByFilter(userId, accountType)
	if err != nil {
		return err
	}
	if len(payAccounts) == 0 {
		return errors.New("No such type of pay account address.")
	}

	payAccounts[0].IsDeleted = TrueValue
	payAccounts[0].AccountType = 0
	payAccounts[0].UpdateTime = time.Now()

	_, err = bxgo.OrmEngin.Id(payAccounts[0].Id).
		Cols("is_deleted", "account_type", "update_time").
		Update(payAccounts[0])

	return err
}

func QueryUserPayAccountByFilter(userId int64, accountType int) ([]*UserPayAccount, error) {
	var userAccounts []*UserPayAccount

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	if userId == 0 {
		return userAccounts, errors.New("User Id should not be zero.")
	}

	session.Where("user_id=? ", userId)

	if accountType != 0 {
		session.Where("account_type=? ", accountType)
	}

	err := session.Find(&userAccounts)

	return userAccounts, err
}