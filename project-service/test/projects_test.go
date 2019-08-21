package test

import (
	"bx.com/project-service/config"
	"bx.com/project-service/bxgo"
	"bx.com/project-service/models"
	"testing"
	"bx.com/project-service/utils"
)

func init() {
	bxgo.CreateOrmEngin(config.DataSource{
		DriverName: "mysql",
		URI:        "root:123456@tcp(192.168.1.145:3306)/test?charset=utf8mb4",
		MaxIdle:    10,
		MaxOpen:    5,
		ShowSQL:    false,
	})

	//bxgo.OrmEngin.Sync2(new(models.Projects))
}

func TestQueryProjectById(t *testing.T) {
	projetId := int64(1)

	result, err := models.QueryProjectById(projetId)
	if err != nil {
		t.Error("%s", err.Error())
	}else{
		t.Logf("%s", result)
	}
}

func TestQueryProjectByUser(t *testing.T) {
	userId := int64(4)

	result, err := models.QueryProjectByUser(userId)
	if err != nil {
		t.Error("%s", err.Error())
	}else{
		for _, item := range result {
			t.Logf("%s", item)
		}
	}
}

func TestQueryAllProjects(t *testing.T) {
	result, err := models.QueryAllProjects()
	if err != nil {
		t.Error("%s", err.Error())
	}else{
		for _, item := range result {
			t.Logf("%s", item)
		}
	}
}

func TestGetDaysInterval(t *testing.T){
	timeInfo, _ := utils.String2TimeWithLocation("2018-06-3 14:01:09")
	days := models.GetDaysInterval(timeInfo)

	t.Logf("Days: %s", days)
}

func TestDiffDays(t *testing.T){
	timeInfo1, _ := utils.String2TimeWithLocation("2018-06-24 19:00:00")
	timeInfo2, _ := utils.String2TimeWithLocation("2018-05-30 18:00:00")
	days1 := models.GetDaysInterval(timeInfo1)
	days2 := models.GetDaysInterval(timeInfo2)

	t.Logf("Days1 %d", days1)
	t.Logf("Days2: %d", days2)

}


