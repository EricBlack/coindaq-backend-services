package models

import (
	"time"
	"github.com/pkg/errors"
	"bx.com/project-service/bxgo"
)

type MembersInfo struct {
	Id           int64     `xorm:"pk autoincr BIGINT(20)"`
	ProjectId    int64     `xorm:"comment('项目id') index BIGINT(20)"`
	Name         string    `xorm:"comment('姓名') TEXT"`
	Position     string    `xorm:"comment('职位') TEXT"`
	Description  string    `xorm:"comment('用户描述i信息') TEXT"`
	MemberType   int       `xorm:"comment('用户类型') INT(11)"`
	Image        string    `xorm:"comment('头像') TEXT"`
	PrioritySort int       `xorm:"comment('排序') INT(11)"`
	JoinTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('加入项目时间') TIMESTAMP"`
	UpdateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
}

func QueryProjectMemberList(projectId int64, memeberType int) ([]*MembersInfo, error) {
	var memberList []*MembersInfo
	if projectId == 0 {
		return memberList, errors.New("Project id parameter should be provided.")
	}

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	session.Where("project_id=? ", projectId)
	if memeberType != 0 {
		session.Where("member_type=? ", memeberType)
	}
	session.Desc("priority_sort")

	err := session.Find(&memberList)

	return memberList, err
}