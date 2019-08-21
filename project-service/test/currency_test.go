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

func TestQueryCurrency(t *testing.T) {
	result, err := models.QueryCurrencyById("100001")
	if err != nil {
		t.Error("%s", err.Error())
	}else{
		t.Logf("%s", result)
	}
}
