package models

import (
	"time"
	"bx.com/project-service/bxgo"
	"fmt"
)

type Projects struct {
	Id               int64     `xorm:"pk autoincr BIGINT(20)"`
	AdminId          int64     `xorm:"comment('管理员Id') index BIGINT(20)"`
	Summary          string    `xorm:"comment('项目名称') TEXT"`
	Description      string    `xorm:"comment('项目详细介绍') TEXT"`
	TargetValue      int64     `xorm:"comment('项目集资总额') BIGINT(20)"`
	SoftValue        int64     `xorm:"comment('软顶') BIGINT(20)"`
	HardValue        int64     `xorm:"comment('硬顶') BIGINT(20)"`
	IssueCoina       string    `xorm:"comment('项目发布币种') VARCHAR(20)"`
	IssueCoinb       string    `xorm:"comment('项目提币币种') VARCHAR(20)"`
	StageCount       int       `xorm:"comment('项目分期数量') INT(11)"`
	LockType         int       `xorm:"comment('结算类型 1: 统一结算，2：锁定结算') INT(11)"`
	Classify         string    `xorm:"comment('项目类别') TEXT"`
	WhitePaper       string    `xorm:"comment('白皮书') TEXT"`
	OfficialSite     string    `xorm:"comment('官网地址') TEXT"`
	CommunityAddress string    `xorm:"comment('社群地址') TEXT"`
	Status           int       `xorm:"comment('项目状态 0:未开始， 1:已开始, 2:已结束') INT(11)"`
	PrioritySort     int       `xorm:"comment('排序优先级') INT(11)"`
	CreateTime       time.Time `xorm:"DATETIME"`
	BeginTime        time.Time `xorm:"DATETIME"`
	EndTime          time.Time `xorm:"DATETIME"`
	UpdateTime       time.Time `xorm:"DATETIME"`
}

func (p Projects) TableName() string {
	return "projects"
}

func QueryProjectById(id int64) (*Projects, error) {
	var project Projects

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	_, err := session.Where("id=? ", id).Get(&project)

	return &project, err
}

func QueryAllProjects() ([]*Projects, error) {
	var projectList []*Projects

	err := bxgo.OrmEngin.Find(&projectList)

	return projectList, err
}

func QueryProjectByUser(userId int64) ([]*Projects, error) {
	var projectList []*Projects

	//获取关注项目id
	userProjectIds, err := QueryUserIcoProject(userId)
	if err != nil || len(userProjectIds) == 0 {
		return projectList, err
	}

	err = bxgo.OrmEngin.In("id", userProjectIds).
		Desc("priority_sort").
		Find(&projectList)

	return projectList, err
}

func CheckUserCanVoteProject(projectId, userId int64) (bool, error) {
	//查找项目
	projectInfo, err := QueryProjectById(projectId)
	if err != nil {
		return false, err
	}
	//查找用户钱包
	userBalance, err := QueryUserBalanceByFilter(userId, projectInfo.IssueCoina)
	if err != nil {
		return false, err
	}

	if userBalance.BalanceValue >0 {
		return true, nil
	}else {
		return false, nil
	}
}

func GetDaysInterval(timeInfo time.Time) int {
	year1 := time.Now().Year()
	month1 := time.Now().Month()
	day1 := time.Now().Day()

	year2 := timeInfo.Year()
	month2 := timeInfo.Month()
	day2 := timeInfo.Day()

	d2 := time.Date(year2, time.Month(month2), day2, timeInfo.Hour(), 0, 0, 0, time.UTC)
	d1 := time.Date(year1, time.Month(month1), day1, time.Now().Hour(), 0, 0, 0, time.UTC)
	diff := d2.YearDay() - d1.YearDay()

	for y := year1; y < year2; y++ {
		diff += time.Date(y, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	}

	if diff <0 {
		return 0
	}

	return diff
}

func UpdateProjectsStatusInfo() error {
	projects, err := QueryAllProjects()
	if err != nil {
		fmt.Errorf("Query all projects error：%s", err)
	}
	for _, item := range projects {
		if item.Status == NotStart && time.Now().After(item.BeginTime) && time.Now().Before(item.EndTime){
			item.Status = Started
			if _, err = bxgo.OrmEngin.Id(item.Id).Cols("status").Update(item); err != nil {
				fmt.Errorf("Update project status error: %s", err.Error())
			}

			continue
		}
		if item.Status == Started && time.Now().After(item.EndTime) {
			item.Status = Completed
			if _, err = bxgo.OrmEngin.Id(item.Id).Cols("status").Update(item); err != nil {
				fmt.Errorf("Update project status error: %s", err.Error())

				continue
			}
		}
	}
	return nil
}