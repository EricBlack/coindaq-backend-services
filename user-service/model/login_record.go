package model

import (
	"time"

	"bx.com/user-service/bxgo"
	"bx.com/user-service/utils"
)

const (
	_ = iota
	LoginPassed
	LoginFailed
)

type LoginRecord struct {
	Id           int64
	UserId       int64     `xorm:"user_id bigint notnull"`
	LoginIp    	 string    `xorm:"login_ip text"`
	DeviceId     string    `xorm:"device_id text"`
	LoginStatus  int32     `xorm:"login_status int"`
	LoginComment string	   `xorm:"login_comment text"`
	LoginTime    time.Time `xorm:"login_time datetime"`
}

type LoginFilter struct {
	UserId 	int64
	LoginStatus	int32
	LoginTime	string
}

func (lr LoginRecord) TableName() string {
	return "login_records"
}

func CreateLoginRecord(lr *LoginRecord) (err error){
	lr.LoginTime = time.Now()

	_, err = bxgo.OrmEngin.Insert(lr)
	return err
}

func QueryLoginRecordsByFilter(filter *LoginFilter) ([]*LoginRecord, error){
	var loginList []*LoginRecord

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	if filter.UserId != 0 {
		session.Where("user_id=? ", filter.UserId)
	}
	if filter.LoginStatus != 0 {
		session.Where("login_status=? ", filter.LoginStatus)
	}
	if filter.LoginTime != "" {
		loginTime, err := utils.String2Time(filter.LoginTime)
		if err != nil{
			return loginList, err
		}
		session.Where("login_time>? ", loginTime)
	}

	err := session.Find(&loginList)

	return loginList, err
}