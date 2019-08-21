package test

import (
	"testing"
	"bx.com/user-service/model"
)

func TestQuryCurrency(t *testing.T){
	result ,err := model.GetCurrency("100001")
	if err != nil {
		t.Errorf("%s", err.Error())
	}else{
		t.Logf("%s", result)
	}
}

func TestQueryAllCurrency(t *testing.T) {
	result ,err := model.GetAllCurrency()
	if err != nil {
		t.Errorf("%s", err.Error())
	}else{
		for _, item := range result {
			t.Logf("%s", item)
		}
	}
}

func TestCreateAllCurrency(t *testing.T){
	model.CreateAllUserCoinAddress(1)
}
