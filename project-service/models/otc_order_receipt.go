package models

import (
	"time"
	"github.com/pkg/errors"
	"bx.com/project-service/bxgo"
)

type OtcOrderReceipt struct {
	Id           int64     `xorm:"pk autoincr BIGINT(20)"`
	OtcId        int64     `xorm:"BIGINT(20)"`
	PayAccountId int64     `xorm:"comment('用户收款地址id') BIGINT(20)"`
	CreateTime   time.Time `xorm:"comment('创建时间') DATETIME"`
}

func QueryReceiptAddress(otcId, userAccountId int64) (*OtcOrderReceipt, error) {
	var otcReceipt OtcOrderReceipt
	_, err := bxgo.OrmEngin.Where("otc_id=? ", otcId).
		Where("pay_account_id=? ", userAccountId).Get(&otcReceipt)

	return &otcReceipt, err
}

func AddReceiptAddress(otcId, userId int64, userAccountList []int64) error {
	if otcId == 0 || userId == 0 {
		return errors.New("OtcId and UserId should not be zero.")
	}
	if len(userAccountList) == 0 {
		return errors.New("User pay account should not be empty.")
	}

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	//事物开始
	session.Begin()
	for _, accountId := range userAccountList {
		existAddress, err := QueryReceiptAddress(otcId, accountId)
		if err != nil {
			return err
		}
		if existAddress.Id != 0 {
			continue
		}

		otcReceipt := OtcOrderReceipt{
			OtcId:        otcId,
			PayAccountId: accountId,
			CreateTime:   time.Now(),
		}
		_, err = session.Insert(&otcReceipt)
		if err != nil {
			session.Rollback()
			return err
		}
	}
	//事物提交
	err := session.Commit()

	return err
}

func QueryOtcPayAddress(otcId int64) ([]*UserPayAccount, error) {
	var payAccountList []*UserPayAccount
	var otcReceipts []*OtcOrderReceipt
	err := bxgo.OrmEngin.Where("otc_id=? ", otcId).Find(&otcReceipts)
	if err != nil || len(otcReceipts) == 0 {
		return payAccountList, err
	}else {
		var accountIds []int64
		for _, otcRec := range otcReceipts {
			accountIds = append(accountIds, otcRec.PayAccountId)
		}

		err = bxgo.OrmEngin.Where("is_deleted!=? ", TrueValue).
			In("id", accountIds).Find(&payAccountList)

		return payAccountList, err
	}
}
