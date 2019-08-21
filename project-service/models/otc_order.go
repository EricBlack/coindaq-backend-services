package models

import (
	"time"
	"bx.com/project-service/bxgo"
	"github.com/pkg/errors"
	"github.com/go-xorm/xorm"
)

type OtcOrder struct {
	Id               int64
	UserId           int64     `xorm:"comment('用户id') BIGINT(20)"`
	BuySell          int       `xorm:"comment('买或卖 1=买， 2=卖') INT(11)"`
	CurrencyPair     int64     `xorm:"comment('交易对') BIGINT(20)"`
	QuotePrice       int64     `xorm:"comment('报价') BIGINT(20)"`
	QuoteAmount      int64     `xorm:"comment('出售数量') BIGINT(20)"`
	LeftAmount       int64     `xorm:"comment('剩余数量') BIGINT(20)"`
	MinValue         int64     `xorm:"comment('最小限额') BIGINT(20)"`
	MaxValue         int64     `xorm:"comment('最大限额') BIGINT(20)"`
	ExpireTime       int       `xorm:"comment('过期时间') INT(11)"`
	OpenLimitHour    int       `xorm:"comment('开发时间-小时') INT(11)"`
	OpenLimitMinute  int       `xorm:"comment('开放时间-分钟') INT(11)"`
	CloseLimitHour   int       `xorm:"comment('关闭时间-小时') INT(11)"`
	CloseLimitMinute int       `xorm:"comment('关闭时间-分钟') INT(11)"`
	Status           int       `xorm:"comment('订单状态，1：进行中， 2：已完成， 3：已关闭') INT(11)"`
	AutoReply        int       `xorm:"comment('是否自动回复') INT(11)"`
	AutoMessage      string    `xorm:"comment('回复信息') TEXT"`
	CreateTime       time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateTime       time.Time `xorm:"comment('更新时间') DATETIME"`
}

type OtcOrderFilter struct {
	UserId  		int64
	BuySell			int
	CurrencyPair	int64
	AvailableTime	int
	Status			int
}

type UpdateOtcOrderInfo struct {
	Id				int64
	LeftAmount		int64
	Status			int
}

func (oo *OtcOrder) CheckOtcVaild() (error) {
	if oo.UserId == 0 {
		return errors.New("User Id should be not zero")
	}
	if oo.BuySell == 0 {
		return errors.New("Buy or sell parameter not set.")
	}
	if oo.CurrencyPair == 0 {
		return errors.New("Currency pair parameter not set.")
	}
	if oo.MaxValue < oo.MinValue {
		return errors.New("Order limit value set conflicted.")
	}
	if oo.CloseLimitHour < oo.OpenLimitHour {
		return errors.New("Order limit display time set conflicted.")
	}

	return nil
}

//广告订单
func (oo *OtcOrder) CreateOtcOrder() (error) {
	if err := oo.CheckOtcVaild(); err != nil {
		return err
	}
	oo.CreateTime = time.Now()

	_, err := bxgo.OrmEngin.Insert(oo)

	return err
}

//发送短消息
func (oo *OtcOrder) SendMessage(userId int64) error {
	if oo.AutoReply == TrueValue {
		err := SendMessage(oo.UserId, userId, oo.AutoMessage)
		return err
	}
	return nil
}

//更新订单
func UpdateOtcOrder(session *xorm.Session, updateInfo UpdateOtcOrderInfo) (error) {
	if updateInfo.Id == 0 {
		return errors.New("Otc Id should not be zero.")
	}

	otcInfo, err := QueryOtcOrderById(updateInfo.Id)
	if err != nil || otcInfo.UserId == 0 {
		return err
	}

	if updateInfo.LeftAmount != 0 {
		otcInfo.LeftAmount = updateInfo.LeftAmount
	}
	if updateInfo.Status != 0 {
		otcInfo.Status = updateInfo.Status
	}
	otcInfo.UpdateTime = time.Now()

	_, err = session.Where("id=? ", otcInfo.Id).Update(otcInfo)

	return err
}

//查询广告订单
func QueryOtcOrderById(id int64) (*OtcOrder, error) {
	var order OtcOrder

	_, err := bxgo.OrmEngin.Id(id).Get(&order)

	return &order, err
}

//查询所有广告订单
func QueryAllOtcOrderList() ([]*OtcOrder, error){
	var allOtcOrders []*OtcOrder

	err := bxgo.OrmEngin.Find(&allOtcOrders)

	return allOtcOrders, err
}

//筛选所有广告订单
func QueryOtcOrdersByFilter(filter OtcOrderFilter) ([]*OtcOrder, error) {
	var otcOrderList []*OtcOrder

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	if filter.UserId != 0 {
		session.Where("user_id=? ", filter.UserId)
	}
	if filter.CurrencyPair != 0 {
		session.Where("currency_pair=? ", filter.CurrencyPair)
	}
	if filter.Status != 0 {
		session.Where("status=? ", filter.Status)
	}

	err := session.Find(&otcOrderList)
	if err != nil {
		return otcOrderList, err
	}

	//筛选日开放时间
	if filter.AvailableTime == 0 {
		return otcOrderList, err
	} else{
		var availableOtcOrders []*OtcOrder
		for _, item := range otcOrderList {
			if result := CheckOtcTimeAvailable(item); result {
				availableOtcOrders = append(availableOtcOrders, item)
			}
		}

		return availableOtcOrders, nil
	}

}

func CheckOtcTimeAvailable(order *OtcOrder) (bool) {
	currentHour := time.Now().Hour()
	currentMinute := time.Now().Minute()

	if order.OpenLimitHour < currentHour && order.CloseLimitHour > currentHour {
		return true
	}
	if order.OpenLimitHour == currentHour && order.OpenLimitMinute < currentMinute {
		return true
	}
	if order.CloseLimitHour == currentHour && order.CloseLimitMinute >currentMinute {
		return true
	}

	return false
}
