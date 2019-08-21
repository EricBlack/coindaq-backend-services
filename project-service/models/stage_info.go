package models

import (
	"time"
	"bx.com/project-service/bxgo"
	"github.com/pkg/errors"
	"github.com/go-xorm/xorm"
	"fmt"
)

type StageInfo struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	ProjectId     int64     `xorm:"comment('项目Id') index BIGINT(20)"`
	StageNumber   int       `xorm:"comment('项目期数信息') INT(11)"`
	StageName     string    `xorm:"comment('阶段名称') TEXT"`
	CoinCount     int       `xorm:"comment('募资币种') INT(11)"`
	TargetValue   int64     `xorm:"comment('募集资金') BIGINT(20)"`
	SoftValue     int64     `xorm:"comment('软顶') BIGINT(20)"`
	Discount      string    `xorm:"comment('折扣') TEXT"`
	CompleteValue int64     `xorm:"comment('当前完成金额') BIGINT(20)"`
	StageStatus   int       `xorm:"comment('分期状态') INT(11)"`
	BeginTime     time.Time `xorm:"not null DATETIME"`
	EndTime       time.Time `xorm:"not null DATETIME"`
	CreateTime    time.Time `xorm:"not null DATETIME"`
	UpdateTime    time.Time `xorm:"not null DATETIME"`
}

func QueryProjectStageList(projectId int64) ([]*StageInfo, error){
	var stageList []*StageInfo

	err := bxgo.OrmEngin.Where("project_id=? ", projectId).
		Asc("stage_number").Find(&stageList)

	return stageList, err
}

func QueryAllStageList() ([]*StageInfo, error) {
	var stageList []*StageInfo

	err := bxgo.OrmEngin.Find(&stageList)

	return stageList, err
}

func QueryProjectStageByFilter(projectId int64, stageNo int) (*StageInfo, error) {
	var stageInfo StageInfo

	if projectId ==0 || stageNo == 0 {
		return &stageInfo, errors.New("Both project id and stage number should be not zero.")
	}
	_, err := bxgo.OrmEngin.Where("project_id=? ", projectId).
		Where("stage_number=? ", stageNo).Get(&stageInfo)

	return &stageInfo, err
}

//查询当前运行中的阶段
func QueryProjectAvailableStage(projectId int64) (*StageInfo, error) {
	stageList, err := QueryProjectStageList(projectId)
	if err != nil {
		return nil, err
	}
	for _, stage := range stageList {
		if stage.StageStatus == Started {
			return stage, nil
		}
	}
	return nil, nil
}

//查询唯一展示阶段
func QueryDisplayStageInfo(projectId int64) ([]*StageInfo, error) {
	var stageList []*StageInfo
	var resultStages []*StageInfo
	err := bxgo.OrmEngin.Where("project_id=? ", projectId).Find(&stageList)
	if err != nil || len(stageList) == 0 {
		return stageList, err
	}

	for _, stage := range stageList {
		if stage.StageStatus == Started {
			resultStages = append(resultStages, stage)
			return resultStages, nil
		}
		if time.Now().After(stage.BeginTime) && time.Now().Before(stage.EndTime){
			resultStages = append(resultStages, stage)
			return resultStages, nil
		}
	}

	if stageList[0].StageStatus == NotStart {
		resultStages = append(resultStages, stageList[0])
		return resultStages, nil
	} else {
		length := len(stageList)
		resultStages = append(resultStages, stageList[length-1])
		return resultStages, nil
	}
}

func CheckProjectStageIsCompleted(projectId int64, stageNo int) (bool, error){
	stageInfo, err := QueryProjectStageByFilter(projectId, stageNo)
	if err != nil {
		return false, err
	}
	if stageInfo.Id == 0 {
		return false, errors.New("No such stage information.")
	}

	if stageInfo.StageStatus == Completed {
		return true, nil
	}
	if stageInfo.CompleteValue >= stageInfo.TargetValue {
		return true, nil
	}
	return false, nil
}

func (stageInfo *StageInfo) UpdateStageInfo(session *xorm.Session, investmentVolumn int64) error {
	stageInfo.CompleteValue += investmentVolumn
	stageInfo.UpdateTime = time.Now()
	if stageInfo.CompleteValue >= stageInfo.TargetValue {
		stageInfo.StageStatus = Completed
	}

	_, err := session.Where("id=?", stageInfo.Id).Update(stageInfo)

	return err
}

//更新阶段状态信息
func UpdateStageStatusInfo() error {
	stages, err := QueryAllStageList()
	if err != nil {
		fmt.Errorf("Query all stages error：%s", err)
	}
	for _, item := range stages {
		if item.StageStatus == NotStart && time.Now().After(item.BeginTime) && time.Now().Before(item.EndTime){
			item.StageStatus = Started
			if _, err = bxgo.OrmEngin.Id(item.Id).Cols("status").Update(item); err != nil {
				fmt.Errorf("Update project status error: %s", err.Error())
			}

			continue
		}
		if item.StageStatus == Started && time.Now().After(item.EndTime) {
			item.StageStatus = Completed
			if _, err = bxgo.OrmEngin.Id(item.Id).Cols("status").Update(item); err != nil {
				fmt.Errorf("Update project status error: %s", err.Error())

				continue
			}
		}
	}
	return nil
}