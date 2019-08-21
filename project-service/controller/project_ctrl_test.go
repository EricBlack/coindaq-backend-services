package controller

import (
	"bx.com/project-service/bxgo"
	"bx.com/project-service/config"
	"bx.com/project-service/models"
	"testing"
	"context"
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

func TestQueryProjectPhotosInfo(t *testing.T){
	s := ProjectController{}
	resp, err := s.QueryProjectPhotosInfo(context.Background(), &proto.IdReq{Id:1})
	if err != nil {
		t.Errorf("%s", err)
	} else {
		for _, item := range resp.MediaList {
			t.Logf("%s", item)
		}
	}
}

func TestQueryProjectVideosInfo(t *testing.T){
	s := ProjectController{}
	resp, err := s.QueryProjectVideosInfo(context.Background(), &proto.IdReq{Id:1})
	if err != nil {
		t.Errorf("%s", err)
	} else {
		for _, item := range resp.MediaList {
			t.Logf("%s", item)
		}
	}
}