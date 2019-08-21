package model

import (
	"fmt"
	"testing"

	"bx.com/user-service/config"
)

func init() {
	CreateOrmEngin(config.DataSource{
		DriverName: "postgres",
		URI:        "postgres://postgres:123456@localhost:5432/wallet?sslmode=disable",
		MaxIdle:    3,
		MaxOpen:    3,
		ShowSQL:    true,
	})
	ormEngin.Sync2(new(LightWalletCoin))
}

func TestCreateLightWalletCoin(t *testing.T) {
	lwc := LightWalletCoin{
		WalletName: "test",
		UserId:     int64(1),
		CoinId:     int64(3),
		Address:    "asdfasdfasdfasdf",
	}
	err := CreateLightWalletCoin(&lwc)
	if err != nil {
		t.Fatal(err)
	}
	// existedId, err := CreateLightWalletCoin(lwc)
	// if id != existedId {
	// 	t.Fatal("not return exist lwc id when create exist")
	// }
}

func TestGetLightWalletCoin(t *testing.T) {
	userId := int64(1)
	coinId := int64(3)
	walletName := "test"
	lwc, err := GetLightWalletCoin(userId, walletName, coinId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(lwc)
	userId = int64(2)
	lwc, err = GetLightWalletCoin(userId, walletName, coinId)
	if lwc != nil {
		t.Error("lwc should be nil, but not")
	}
}
