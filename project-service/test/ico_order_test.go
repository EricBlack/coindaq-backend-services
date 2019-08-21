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

	bxgo.OrmEngin.Sync2(new(models.IcoOrder))
}

func TestJoinProjectIco(t *testing.T){
	projectId := int64(1)
	userId := int64(4)
	buyCount := int64(2000)
	payCount := int64(20)
	stageNo := 1
	currencyId := "10001"

	err := models.JoinProjectIco(projectId, userId, buyCount, payCount, stageNo, currencyId)
	if err != nil {
		t.Error("%s", err.Error())
	}
}
