package test

import (
	"bx.com/project-service/config"
	"bx.com/project-service/bxgo"
	"bx.com/project-service/models"
	"testing"
	"context"
	"bx.com/project-service/controller"
	"bx.com/project-service/proto"
)


func init() {
	bxgo.CreateOrmEngin(config.DataSource{
		DriverName: "mysql",
		URI:        "root:123456@tcp(192.168.1.145:3306)/test?charset=utf8mb4",
		MaxIdle:    10,
		MaxOpen:    5,
		ShowSQL:    true,
	})

	bxgo.OrmEngin.Sync2(new(models.Projects))
}

func TestRecommendProject (t *testing.T) {
	s := controller.ProjectController{}
	projectList, err := s.QueryRecommendProjectsInfo(context.Background(), &proto.Empty{})
	if err != nil {
		t.Errorf("%s", err.Error())
	}else {
		for _, project := range projectList.ProjectList {
			t.Logf("%s", project)
		}
	}
}

func TestProjectInfo(t *testing.T) {
	s := controller.ProjectController{}
	idReq := &proto.IdReq{ Id:	1}
	projectResult, err := s.QueryProjectDetailsInfo(context.Background(), idReq)

	if err != nil {
		t.Errorf("%s", err.Error())
	}else {
		t.Logf("%s", projectResult)
	}
}


