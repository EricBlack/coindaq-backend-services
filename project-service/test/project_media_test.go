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

	bxgo.OrmEngin.Sync2(new(models.ProjectsMedia))
}


func TestQueryAllProjectMedias(t *testing.T) {
	projectId := int64(1)

	result, err := models.QueryProjectMediaList(projectId, 0)
	if err != nil {
		t.Error("%s", err.Error())
	}else{
		for _, item := range result {
			t.Logf("%s", item)

		}
	}
}

func TestQueryAllProjectPhoto(t *testing.T) {
	projectId := int64(1)

	result, err := models.QueryProjectMediaList(projectId, 1)
	if err != nil {
		t.Error("%s", err.Error())
	}else{
		for _, item := range result {
			t.Logf("%s", item)

		}
	}
}

func TestQueryAllProjectVedio(t *testing.T) {
	projectId := int64(1)

	result, err := models.QueryProjectMediaList(projectId, 2)
	if err != nil {
		t.Error("%s", err.Error())
	}else{
		for _, item := range result {
			t.Logf("%s", item)
		}
	}
}