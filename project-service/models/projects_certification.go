package models

import (
	"time"
	"bx.com/project-service/bxgo"
)

type ProjectsCertification struct {
	Id              int64     `xorm:"pk autoincr BIGINT(20)"`
	ProjectId       int64     `xorm:"index BIGINT(20)"`
	CertificationId int64     `xorm:"index BIGINT(20)"`
	PrioritySort    int       `xorm:"comment('排序优先级') INT(11)"`
	CreateTime      time.Time `xorm:"comment('添加时间') DATETIME"`
}

func QueryProjectCertificationList(projectId int64) ([]*CertificationInfo, error){
	var projectCertificatnList []*ProjectsCertification
	var certificationList []*CertificationInfo

	err := bxgo.OrmEngin.Where("project_id=? ", projectId).
		Desc("priority_sort").
		Find(&projectCertificatnList)

	if err != nil || len(projectCertificatnList) == 0 {
		return certificationList, err
	}

	var certificationIdList []int64
	for _, projectCertification := range projectCertificatnList {
		certificationIdList = append(certificationIdList, projectCertification.CertificationId)
	}
	certificationList, err = QueryCertificationByIdList(certificationIdList)

	return certificationList, err
}
