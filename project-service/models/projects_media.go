package models

import (
	"time"
	"errors"
	"bx.com/project-service/bxgo"
)

type ProjectsMedia struct {
	Id           int64     `xorm:"pk autoincr BIGINT(20)"`
	ProjectId    int64     `xorm:"comment('项目id') index BIGINT(20)"`
	Title        string    `xorm:"comment('标题') TEXT"`
	Address      string    `xorm:"comment('地址') TEXT"`
	Type         int       `xorm:"default 0 comment('类型，图片或视频') INT(11)"`
	Enable       int       `xorm:"default 0 comment('是否显示') INT(11)"`
	PrioritySort int       `xorm:"comment('排序优先级') INT(11)"`
	CreateTime   time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

func QueryProjectMediaList(projectId int64, mediaType int) ([]*ProjectsMedia, error) {
	var mediaList []*ProjectsMedia

	if projectId == 0 {
		return mediaList, errors.New("Project id should be provided.")
	}

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	session.Where("project_id=? ", projectId)
	if mediaType != 0 {
		session.Where("type=? ", mediaType)
	}
	session.Desc("priority_sort")

	err := session.Find(&mediaList)

	return mediaList, err
}
