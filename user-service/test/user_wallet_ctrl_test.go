package test

import (
	"bx.com/user-service/bxgo"
	"bx.com/user-service/model"
	"bx.com/user-service/config"
	ctr "bx.com/user-service/controller"
	"testing"
	"context"
	log "github.com/sirupsen/logrus"
	"bx.com/user-service/proto"
)
func init() {
	bxgo.CreateOrmEngin(config.DataSource{
		DriverName: "mysql",
		URI:        "root:123456@tcp(192.168.1.145:3306)/test?charset=utf8mb4",
		MaxIdle:    10,
		MaxOpen:    5,
		ShowSQL:    true,
	})
	bxgo.OrmEngin.Sync2(new(model.UserBalance))
	bxgo.OrmEngin.Sync2(new(model.UserWallet))
}

func TestUserWalletControl_AddWalletAddress(t *testing.T) {
	walletAddress := &proto.WalletAddressReq{
		UserId:	10,
		CurrencyName: "BTC",
		WalletAddress: "dfvghnjkljydfgdshgfhgfjhgjfgdguhgtrfd",
	}
	s := ctr.UserWalletControl{}
	_, err := s.AddWalletAddress(context.Background(), walletAddress)
	if err != nil {
		log.Errorf(err.Error())
		t.Errorf("TestSignup got unexpected error")
	}
}

func TestUserWalletControl_DeleteUserWalletAddress(t *testing.T) {
	s := ctr.UserWalletControl{}
	_, err := s.DeleteUserWalletAddress(context.Background(), &proto.IdReq{Id:2})
	if err != nil {
		log.Errorf(err.Error())
		t.Errorf("TestSignup got unexpected error")
	}
}

func TestUserWalletControl_BindUserCoinAddress(t *testing.T) {
	userCoinReq := &proto.UserCoinAddressReq{
		UserId:	10,
		CoinId: "100001",
	}
	s := ctr.UserWalletControl{}
	_, err := s.BindUserCoinAddress(context.Background(), userCoinReq)
	if err != nil {
		log.Errorf(err.Error())
		t.Errorf("TestSignup got unexpected error")
	}
}

func TestUserWalletControl_QueryUserCoinAddress(t *testing.T) {
	s := ctr.UserWalletControl{}
	userCoinReq := &proto.UserCoinAddressReq{
		UserId:		10,
	}
	resp, err := s.QueryUserCoinAddress(context.Background(), userCoinReq)
	if err != nil {
		log.Errorf(err.Error())
		t.Errorf("TestSignup got unexpected error")
	}else{
		log.Printf("%s", resp.UserCoinList)
	}
}

func TestUserWalletControl_QueryUserWalletAddress(t *testing.T) {
	/*
	walletAddress := &proto.WalletAddressReq{
		UserId:	10,
		CurrencyName: "BTC",
	}
	*/

	walletAddress := &proto.WalletAddressReq{
		UserId:	10,
	}

	s := ctr.UserWalletControl{}
	resp, err := s.QueryUserWalletAddress(context.Background(), walletAddress)
	if err != nil {
		log.Errorf(err.Error())
		t.Errorf("TestSignup got unexpected error")
	}else{
		log.Printf("%s", resp.AddressList)
	}
}

func TestQueryUserBalanceByFilter(t *testing.T) {
	result, err := model.QueryBalanceByFilter(10, "100001")
	if err != nil {
		t.Errorf("%s", err.Error())
	}else{
		t.Logf("%s", result)
	}
}
