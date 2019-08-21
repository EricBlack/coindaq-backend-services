package test

import (
	"bx.com/user-service/bxgo"
	"bx.com/user-service/config"
	"bx.com/user-service/model"
	"testing"
)

func init() {
	/*
	bxgo.CreateOrmEngin(config.DataSource{
		DriverName: "postgres",
		URI:        "postgres://postgres:postgres@localhost:5432/userinfo?sslmode=disable",
		MaxIdle:    10,
		MaxOpen:    5,
		ShowSQL:    true,
	})
	*/
	bxgo.CreateOrmEngin(config.DataSource{
		DriverName: "mysql",
		URI:        "root:123456@tcp(192.168.1.145:3306)/test?charset=utf8mb4",
		MaxIdle:    10,
		MaxOpen:    5,
		ShowSQL:    true,
	})

	bxgo.OrmEngin.Sync2(new(model.MsgRecords))
}

func TestMessageInHour(t *testing.T){
	result, err := model.MessageInHour("+8618500797779")
	if err != nil {
		t.Errorf("%s", err)
	}else {
		t.Logf("%s", result)
	}
}

