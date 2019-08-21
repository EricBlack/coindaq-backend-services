package models

import (
	"time"
	"bx.com/project-service/bxgo"
)

type ProjectsText struct {
	Id           int64     `xorm:"pk autoincr BIGINT(20)"`
	ProjectId    int64     `xorm:"BIGINT(20)"`
	Title        string    `xorm:"TEXT"`
	ProjectText  string    `xorm:"TEXT"`
	Enable       int       `xorm:"INT(11)"`
	PrioritySort int       `xorm:"INT(11)"`
	CreateTime   time.Time `xorm:"TIMESTAMP"`
	UpdateTime   time.Time `xorm:"TIMESTAMP"`
}

func QueryProjectsTextByProject(projectId int64) ([]*ProjectsText, error) {
	var projectTextList []*ProjectsText

	err := bxgo.OrmEngin.Where("project_id=? ", projectId).
		Asc("priority_sort").Find(&projectTextList)

	return projectTextList, err
}
