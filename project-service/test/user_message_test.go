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

	bxgo.OrmEngin.Sync2(new(models.UserMessage))
}

func TestAddMessage(t *testing.T){
	models.SendMessage(1,2,"Hello, this is mike.")
	models.SendMessage(1,2,"Hello, Is there?")
	models.SendMessage(2,1,"Hello, this is Jake.")
	models.SendMessage(1,2,"Question 1")
	models.SendMessage(2,1,"Answer 1")
	models.SendMessage(2,1,"Answer 2")
	models.SendMessage(1,2,"Good boy")
}

func TestQueryMessage(t *testing.T){
	result, err :=models.QueryUserMessage(1,2, 0)
	if err != nil {
		t.Errorf("%s", err.Error())
	}else{
		for _, item := range result {
			t.Logf("%s", item)
		}
	}
}
