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

	bxgo.OrmEngin.Sync2(new(models.NoticeVote))
}

func TestQueryNoticeVote(t *testing.T) {
	result, err := models.QueryNoticeVote(2)
	if err != nil {
		t.Error("%s", err.Error())
	}else{
		t.Logf("%s", result)
	}
}

func TestUpdateNoticeVoteTimeStatus(t *testing.T){
	begin := "2018-05-21 10:00:00"
	after := "2018-06-30 18:00:00"
	err := models.UpdateNoticeVoteTimeStatus(2, begin, after)

	if err != nil {
		t.Error("%s", err.Error())
	}
}

func TestQueryVoteVolumnInfo(t *testing.T){
	approve, disapprove, absence, err := models.QueryVoteVolumnInfo(2)
	if err != nil {
		t.Error("%s", err.Error())
	}else{
		t.Logf("Approve=%s, Disapprove=%s, Absence=%s", approve, disapprove, absence)
	}
}

func TestUserVoteNotice(t *testing.T) {
	err := models.UserVoteNotice(8, 2, 1)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}
