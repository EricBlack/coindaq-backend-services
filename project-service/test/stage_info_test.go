package test

import (
	"bx.com/project-service/config"
	"bx.com/project-service/bxgo"
	"bx.com/project-service/models"
	"testing"
)


func init() {
	bxgo.CreateOrmEngin(config.DataSource{
		DriverName: "mysql",
		URI:        "root:123456@tcp(192.168.1.145:3306)/test?charset=utf8mb4",
		MaxIdle:    10,
		MaxOpen:    5,
		ShowSQL:    true,
	})

	bxgo.OrmEngin.Sync2(new(models.StageInfo))
}

func TestQueryProjectStageList(t *testing.T){
	projectId := int64(1)
	result, err := models.QueryProjectStageList(projectId)
	if err != nil {
		t.Error("%s", err.Error())
	}else{
		for _, item := range result {
			t.Logf("%s", item)
		}
	}
}

func TestQueryProjectStageByFilter1(t *testing.T){
	projectId := int64(1)
	result, err := models.QueryProjectStageByFilter(projectId, 1)
	if err != nil {
		t.Error("%s", err.Error())
	}else{
		t.Logf("%s", result)
	}
}

func TestQueryProjectStageByFilter2(t *testing.T){
	projectId := int64(1)
	result, err := models.QueryProjectStageByFilter(projectId, 2)
	if err != nil {
		t.Error("%s", err.Error())
	}else{
		t.Logf("%s", result)
	}
}

func TestCheckProjectStageIsCompleted(t *testing.T){
	projectId := int64(1)
	stageNo := 3
	result, err := models.CheckProjectStageIsCompleted(projectId, stageNo)

	if err != nil {
		t.Errorf("%s", err.Error())
	}else{
		t.Logf("%s", result)
	}
}