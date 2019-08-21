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

	bxgo.OrmEngin.Sync2(new(models.ProjectsText))
}

func TestQueryProjectsText(t *testing.T) {
	result, err := models.QueryProjectsTextByProject(1)
	if err != nil {
		t.Errorf("%s", err.Error())
	}else{
		for _, par := range result {
			t.Logf("%s", par)
		}
	}
}
