package models

import (
	"time"
	"bx.com/project-service/bxgo"
	"errors"
)

type NoticeNews struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	ProjectId   int64     `xorm:"comment('项目id') index BIGINT(20)"`
	Title       string    `xorm:"comment('公告标题') TEXT"`
	Description string    `xorm:"comment('公告描述信息') TEXT"`
	Image       string    `xorm:"comment('上传图片') TEXT"`
	ImageW      int       `xorm:"comment('图片宽') INT(11)"`
	ImageH      int       `xorm:"comment('图片高') INT(11)"`
	UrlSite     string    `xorm:"comment('用户超链接') TEXT"`
	NoticeType  int       `xorm:"INT(11)"`
	SendType    int       `xorm:"comment('公告接受用户') INT(11)"`
	Status      int       `xorm:"comment('信息状态') INT(11)"`
	Reason      string    `xorm:"comment('审核原因') TEXT"`
	CreateTime  time.Time `xorm:"not null DATETIME"`
	ExpireTime  time.Time `xorm:"not null comment('过期时间') DATETIME"`
	UpdateTime  time.Time `xorm:"not null DATETIME"`
}

func QueryNoticeNewsById(id int64) (*NoticeNews, error){
	var noticeNews NoticeNews

	_, err := bxgo.OrmEngin.Id(id).Get(&noticeNews)

	return &noticeNews, err
}

func QueryNoticesByProjectId(projectId int64) ([]*NoticeNews, error) {
	var noticeList []*NoticeNews

	err := bxgo.OrmEngin.Where("project_id=? ", projectId).Find(&noticeList)

	return noticeList, err
}

func QueryNoticesByFilter(projectId int64, noticeType, status int) ([]*NoticeNews, error) {
	var noticeList []*NoticeNews

	if projectId == 0 {
		return noticeList, errors.New("Project id should be provided.")
	}
	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	session.Where("project_id=? ", projectId)
	if noticeType != 0 {
		session.Where("notice_type=? ", noticeType)
	}
	if status != 0 {
		session.Where("status=? ", status)
	}

	err := session.Find(&noticeList)

	return noticeList, err
}