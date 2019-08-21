package models

import (
	"time"
	"bx.com/project-service/bxgo"
)

type CertificationInfo struct {
	Id           int64     `xorm:"pk autoincr BIGINT(20)"`
	Name         string    `xorm:"comment('认证名称') TEXT"`
	Description  string    `xorm:"comment('认证描述信息') TEXT"`
	Image        string    `xorm:"comment('认证图像地址') TEXT"`
	PrioritySort int       `xorm:"comment('排序优先级') INT(11)"`
	CreateTime   time.Time `xorm:"DATETIME"`
}

func QueryCertificationById(id int64) (*CertificationInfo, error) {
	var certificationInfo CertificationInfo

	_, err := bxgo.OrmEngin.Id(id).Get(&certificationInfo)

	return &certificationInfo, err
}

func QueryCertificationByIdList(idList []int64) ([]*CertificationInfo, error) {
	var certificationList []*CertificationInfo

	err := bxgo.OrmEngin.In("id", idList).
		Desc("priority_sort").
		Find(&certificationList)

	return certificationList, err
}
