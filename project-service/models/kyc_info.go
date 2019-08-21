package models

import (
	"time"
	"bx.com/project-service/bxgo"
)

type KycInfo struct {
	Id           int64     `xorm:"pk autoincr BIGINT(20)"`
	UserId       int64     `xorm:"comment('用户id') BIGINT(20)"`
	State        int       `xorm:"comment('状态') INT(11)"`
	Kind         int       `xorm:"comment('用户类型') INT(11)"`
	RealName     string    `xorm:"comment('真实姓名') TEXT"`
	CountryCode  int       `xorm:"comment('国家代码') INT(11)"`
	IdentityType int       `xorm:"comment('证件类型') INT(11)"`
	IdentityId   string    `xorm:"comment('证件Id') TEXT"`
	PhotoFront   string    `xorm:"comment('证件正面照') TEXT"`
	PhotoBack    string    `xorm:"comment('证件背面照') TEXT"`
	PhotoHand    string    `xorm:"comment('证件手持照') TEXT"`
	Reason       string    `xorm:"comment('审核原因') TEXT"`
	CreatedAt    time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdatedAt    time.Time `xorm:"comment('更新时间') DATETIME"`
	RejectedAt   time.Time `xorm:"comment('拒绝时间') DATETIME"`
	PassedAt     time.Time `xorm:"comment('通过时间') DATETIME"`
}

func CheckUserKycStatus(userId int64) (bool, error) {
	kycInfo := KycInfo{}

	_, err := bxgo.OrmEngin.Where("user_id=? ", userId).
		Where("state=? ", 1).Get(&kycInfo)
	if err != nil {
		return false, err
	}else {
		if kycInfo.Id == 0 {
			return false, nil
		} else{
			return true, nil
		}
	}
}
