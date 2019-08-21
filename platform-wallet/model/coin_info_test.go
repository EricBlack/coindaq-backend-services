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
	ormEngin.Sync2(new(CoinInfo))
}

func TestCreateCoinInfo(t *testing.T) {
	coin := CoinInfo{
		Code:        "bitcoin",
		Symbol:      "BTC",
		Name:        "Bitcoin",
		Decimals:    8,
		Kind:        BTC,
		MinConfirms: 6,
	}
	err := CreateCoinInfo(&coin)
	if err != nil {
		t.Error(err)
	}
}

func TestListCoinInfo(t *testing.T) {
	kind := BTC
	coinInfoList, err := ListCoinInfo(kind)
	if err != nil {
		t.Error(err)
	}
	if len(coinInfoList) <= 0 {
		t.Error("len is 0")
	}
	fmt.Printf("coin list is : %v \n", coinInfoList)
}

func TestGetCoinInfo(t *testing.T) {
	coin, err := GetCoinInfo(3)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("coin is  %v \n", coin)
	_, err = GetCoinInfo(0)

	if err == nil {
		t.Error("expect not found, but passed")
	}
	fmt.Println(err)
}

func TestGetCoinInfoBySymbol(t *testing.T) {
	coin, err := GetCoinInfoBySymbol("BTC")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("coin is %v \n", coin)
}

func TestSaveCoinInfo(t *testing.T) {
	coin := CoinInfo{
		Id:          1,
		Code:        "bitcoin",
		Symbol:      "BTC",
		Name:        "Bitcoin1",
		Decimals:    8,
		Kind:        BTC,
		MinConfirms: 6,
	}
	err := coin.Save()
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteCoinInfo(t *testing.T) {
	err := DeleteCoinInfo(2)
	if err != nil {
		t.Error(err)
	}
}

func TestEnableWithdraw(t *testing.T) {
	err := EnableWithdraw(3)
	if err != nil {
		t.Fatal(err)
	}
	coin, err := GetCoinInfo(3)
	if err != nil {
		t.Fatal(err)
	}
	if coin.Withdrawal == CoinWithdrawDisable {
		t.Fatal("enabled, but get disable")
	}
}

func TestDisableWithdraw(t *testing.T) {
	err := DisableWithdraw(3)
	if err != nil {
		t.Fatal(err)
	}
	coin, err := GetCoinInfo(3)
	if err != nil {
		t.Fatal(err)
	}
	if coin.Withdrawal == CoinWithdrawEnable {
		t.Fatal("disabled, but get enable")
	}
}

func TestEnableReceive(t *testing.T) {
	err := EnableReceive(3)
	if err != nil {
		t.Fatal(err)
	}
	coin, err := GetCoinInfo(3)
	if err != nil {
		t.Fatal(err)
	}
	if coin.Receivable == CoinReceiveDisable {
		t.Fatal("enabled, but get disable")
	}
}

func TestDisableReceive(t *testing.T) {
	err := DisableReceive(3)
	if err != nil {
		t.Fatal(err)
	}
	coin, err := GetCoinInfo(3)
	if err != nil {
		t.Fatal(err)
	}
	if coin.Withdrawal == CoinReceiveEnable {
		t.Fatal("disabled, but get enable")
	}
}
