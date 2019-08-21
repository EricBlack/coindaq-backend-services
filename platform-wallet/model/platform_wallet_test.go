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
	ormEngin.Sync2(new(PWalletCoin))
}

func TestCreatePWalletCoin(t *testing.T) {
	pwc := PWalletCoin{
		UserId:     int64(1),
		CoinId:     int64(3),
		WalletName: "test",
		Account:    "aafasdfasdfasdf",
		Address:    "adfasdfasdfasdfad",
	}
	err := CreatePWalletCoin(&pwc)
	if err != nil {
		t.Fatal(err)
	}
	err = CreatePWalletCoin(&pwc)
	if err == nil {
		t.Fatal("expect duplicated pwc error, but not")
	}
}

func TestListPWCByUser(t *testing.T) {
	pwcs, err := ListPWCByUser(int64(1))
	if err != nil {
		t.Fatal(err)
	}
	if len(pwcs) <= 0 {
		t.Error("len lte 0")
	}
	fmt.Printf("list len is %d \n", len(pwcs))
}

func TestGetPWalletCoin(t *testing.T) {
	pwc, err := GetPWalletCoin(int64(1), "test", int64(3))
	if err != nil {
		t.Fatal(err)
	}
	if pwc == nil {
		t.Fatal("should found, but not")
	}
	fmt.Printf("pwc is %v \n", pwc)
	pwc, err = GetPWalletCoin(int64(2), "test", int64(3))
	if err != nil {
		t.Fatal(err)
	}
	if pwc != nil {
		t.Fatal("should not found, but found")
	}
}

func TestUpdatePWCAccountAndAddr(t *testing.T) {
	expectedAccount := "2222222"
	expectedAddress := "1111111"
	pWalletCoin, _ := GetPWalletCoin(1, "test", 3)
	err := pWalletCoin.UpdatePWCAccountAndAddr(expectedAccount, expectedAddress)
	if err != nil {
		t.Fatal(err)
	}
	pwc, err := GetPWalletCoin(1, "test", 3)
	if err != nil {
		t.Fatal(err)
	}
	if pwc.Account != expectedAccount {
		t.Error("not expected account")
	}
	if pwc.Address != expectedAddress {
		t.Error("not expected address")
	}
}
