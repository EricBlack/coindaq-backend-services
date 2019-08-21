package test

import (
	"context"
	"testing"

	"bx.com/user-service/bxgo"
	"bx.com/user-service/config"
	"bx.com/user-service/model"
	"bx.com/user-service/proto"
	ctr "bx.com/user-service/controller"
	"github.com/sirupsen/logrus"
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

func TestGenerateFactor(t *testing.T) {
	tf := ctr.TwoFactorController{}
	factorReq := &proto.FactorReq{
		UserId: 1,
		Type:   model.GoogleAuthType,
	}
	res, err := tf.GenerateFactor(context.Background(), factorReq)
	if err != nil {
		logrus.Errorf(err.Error())
		t.Errorf("TestGenerateFactorTest got unexpected error")
	}

	logrus.Infof("%s", res)
}

func TestQueryFactor(t *testing.T) {
	tf := ctr.TwoFactorController{}
	factorReq := &proto.FactorReq{
		UserId: 1,
		Type:   1,
	}
	res, err := tf.QueryFactor(context.Background(), factorReq)
	if err != nil {
		logrus.Errorf(err.Error())
		t.Errorf("QueryFactorTest got unexpected error")
	}
	if res.Id == 0 {
		logrus.Errorf("QueryFactorTest failed, no factor refreshed.")
	}

	logrus.Infof("%s", res)
}

func TestRefreshFactor(t *testing.T) {
	tf := ctr.TwoFactorController{}
	factorReq := &proto.FactorReq{
		UserId: 1,
		Type:   1,
	}
	res, err := tf.RefreshFactor(context.Background(), factorReq)
	if err != nil {
		logrus.Errorf(err.Error())
		t.Errorf("RefreshFactorTest got unexpected error")
	}
	if res.Id == 0 {
		logrus.Errorf("RefreshFactorTest failed, no factor refreshed.")
	}

	logrus.Infof("%s", res)
}

func TestVerifyEmailCode(t *testing.T) {
	tf := ctr.TwoFactorController{}
	infoReq := &proto.InfoReq{
		UserId: 1,
		Type:   model.EmailAuthType,
		Code:	"8083",
	}
	res, err := tf.VerifyFactorCode(context.Background(), infoReq)
	if err != nil {
		logrus.Errorf(err.Error())
		t.Errorf("TestVerifyEmailCode got unexpected error")
	}
	if res.Result == false{
		t.Errorf("TestVerifyEmailCode verified failed")
	}

	logrus.Infof("%s", res)
}

func TestVerifyPhoneCode(t *testing.T) {
	tf := ctr.TwoFactorController{}
	infoReq := &proto.InfoReq{
		UserId: 1,
		Type:   model.PhoneAuthType,
		Code:	"7917",
	}
	res, err := tf.VerifyFactorCode(context.Background(), infoReq)
	if err != nil {
		logrus.Errorf(err.Error())
		t.Errorf("TestVerifyPhoneCode got unexpected error")
	}

	if res.Result == false{
		t.Errorf("VerifyPhoneCode verified failed.")
	}
	logrus.Infof("%s", res)
}

func TestVerifyGoogleCode(t *testing.T) {
	tf := ctr.TwoFactorController{}
	infoReq := &proto.InfoReq{
		UserId: 1,
		Type:   model.GoogleAuthType,
		Code:	"FN57TVHAUMJWI2UU",
	}
	res, err := tf.VerifyFactorCode(context.Background(), infoReq)
	if err != nil {
		logrus.Errorf(err.Error())
		t.Errorf("TestVerifyGoogleCode got unexpected error")
	}

	if res.Result == false{
		t.Errorf("TestVerifyGoogleCode verified failed.")
	}

	logrus.Infof("%s", res)
}

func TestGoogleImageGenerate(t *testing.T){
	factor, err := model.GetTwoFactor(15, 3)
	if err != nil {
		logrus.Errorf(err.Error())
	}else{
		factor.BarcodeImage(nil)
	}
}
