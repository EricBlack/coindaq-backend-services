package test

import (
	"context"
	"bx.com/user-service/bxgo"
	"bx.com/user-service/config"
	"bx.com/user-service/model"
	"bx.com/user-service/proto"
	ctr "bx.com/user-service/controller"
	log "github.com/sirupsen/logrus"
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

	bxgo.OrmEngin.Sync2(new(model.User))
}

func TestQueryKyc(t *testing.T){
	s := ctr.KycController{}
	req := &proto.IdReq{
		Id: 2,
	}

	kycF, err := s.QueryKycInfoById(context.Background(), req)
	if err != nil {
		log.Errorf(err.Error())
		t.Errorf("TestQueryKyc got unexpected error")
	}

	if kycF.Id == 0 {
		t.Errorf("TestQueryKyc return null")
	}
	t.Logf("KycResult: %s", kycF)
}