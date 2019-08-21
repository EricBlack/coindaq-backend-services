package models

import (
	"time"
	"bx.com/project-service/bxgo"
	"github.com/pkg/errors"
	"math/rand"
	"fmt"
)

type OtcOrderDetails struct {
	Id                  int64
	OtcOrderId          int64     `xorm:"BIGINT(20)"`
	JoinUserId          int64     `xorm:"BIGINT(20)"`
	BuyAmount           int64     `xorm:"comment('得到金额') BIGINT(20)"`
	PayAmount           int64     `xorm:"comment('付款数量') BIGINT(20)"`
	OrderTime           time.Time `xorm:"comment('下单时间') DATETIME"`
	SettlementTime      time.Time `xorm:"comment('结算时间') DATETIME"`
	SettlementNote      string    `xorm:"comment('结算备注') TEXT"`
	Status              int       `xorm:"comment('订单详情状态 0：创建， 1：过期， 2：已付款， 3：已打币， 4： 取消， 5：申诉阶段， 6：完成') INT(11)"`
	OwnerPay            int       `xorm:"comment('卖家打款') INT(11)"`
	JoinerPay           int       `xorm:"comment('买家打款') INT(11)"`
	AppealOwner         int       `xorm:"comment('卖家申诉') INT(11)"`
	AppealOwnerMessage  string    `xorm:"comment('申诉信息') TEXT"`
	AppealJoiner        int       `xorm:"comment('买家申诉') INT(11)"`
	AppealJoinerMessage string    `xorm:"comment('申诉信息') TEXT"`
	OrderNumber         string    `xorm:"comment('订单随机号') TEXT"`
	UpdateTime          time.Time `xorm:"comment('更新时间') DATETIME"`
}

type UserOtcFilter struct {
	UserId			int64
	OtcStatus		int
}

//根据Id查询
func GetOrderDetailsById(id int64) (*OtcOrderDetails, error) {
	var otcOrderDetails OtcOrderDetails
	if id == 0 {
		return &otcOrderDetails, errors.New("Otc order details Id should not be zero.")
	}

	_, err := bxgo.OrmEngin.Id(id).Get(&otcOrderDetails)

	return &otcOrderDetails, err
}

//根据广告单Id查询相关订单
func GetOrderDetailsByOtcOrderId(otcOrderId int64) ([]*OtcOrderDetails, error) {
	var orderDetailsList []*OtcOrderDetails

	err := bxgo.OrmEngin.Where("otc_order_id=? ", otcOrderId).Find(&orderDetailsList)

	return orderDetailsList, err
}

//查询用户完成订单
func GetUserCompleteOrderDetails(userId int64) ([]*OtcOrderDetails, error) {
	var orderDetailsList []*OtcOrderDetails

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	session.Where("join_user_id=? ", userId)
	session.Where("status=?", OrderCompleted)

	err := session.Find(&orderDetailsList)

	return orderDetailsList, err
}

//查询用户取消订单
func GetUserCancelOrderDetails(userId int64) ([]*OtcOrderDetails, error) {
	var orderDetailsList []*OtcOrderDetails

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	session.Where("join_user_id=? ", userId)
	session.Where("status=?", OrderCanceled)

	err := session.Find(&orderDetailsList)

	return orderDetailsList, err
}

//查询用户正在进行订单
func GetUserOngoingOrderDetails(userId int64) ([]*OtcOrderDetails, error) {
	var orderDetailsList []*OtcOrderDetails

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	session.Where("join_user_id=? ", userId)
	session.Where("status != and status !=?", OrderCompleted, OrderCanceled)

	err := session.Find(&orderDetailsList)

	return orderDetailsList, err
}

//检查法币交易是否合格
func (ood *OtcOrderDetails) CheckFiatDealOrder (order *OtcOrder) error {
	//检查货币数量限制
	if order.MinValue > ood.BuyAmount || order.MaxValue < ood.BuyAmount {
		return errors.New("Order amount is out of bounds.")
	}
	//检查是否超库存
	if order.LeftAmount < ood.BuyAmount {
		return errors.New("Order amount is over stock quantity.")
	}

	return nil
}

//法币订单操作
func (ood *OtcOrderDetails) ConfirmFiatDealOrder() (*OtcOrderDetails, error) {
	//查询广告信息
	orderInfo, err := QueryOtcOrderById(ood.OtcOrderId)
	if err != nil || orderInfo.UserId == 0 {
		return ood, err
	}

	//判断是否可下单
	err = ood.CheckFiatDealOrder(orderInfo)
	if err != nil {
		return ood, err
	}

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	//开启事务
	session.Begin()

	//修改余量
	orderInfo.LeftAmount -= ood.BuyAmount
	orderInfo.UpdateTime = time.Now()
	_, err = session.Where("id=? ", orderInfo.Id).Update(orderInfo)
	if err != nil {
		session.Rollback()
		return ood, err
	}

	//添加Otc Details订单
	ood.OrderNumber = RandOtcNumberSalt()
	ood.OrderTime = time.Now()

	_, err = session.Insert(ood)
	if err != nil {
		session.Rollback()
		return ood, err
	}

	//提交事务
	if err = session.Commit(); err != nil {
		return ood, err
	}

	//发送消息通知
	err = orderInfo.SendMessage(ood.JoinUserId)
	return ood, err
}

//检查币币交易是否合格
func (ood *OtcOrderDetails) CheckCoinDealOrder (order *OtcOrder) error {
	return nil
}

//币币订单操作
func (ood *OtcOrderDetails) ConfirmCoinDealOrder() (*OtcOrderDetails, error) {
	//查询广告信息
	orderInfo, err := QueryOtcOrderById(ood.OtcOrderId)
	if err != nil || orderInfo.UserId == 0 {
		return ood, err
	}

	//检查是否可下单
	err = ood.CheckCoinDealOrder(orderInfo)
	if err != nil {
		return ood, err
	}

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	//开启事务
	session.Begin()

	//修改余量
	orderInfo.LeftAmount -= ood.BuyAmount
	orderInfo.UpdateTime = time.Now()
	_, err = session.Where("id=? ", orderInfo.Id).Update(orderInfo)
	if err != nil {
		session.Rollback()
		return ood, err
	}

	//添加Otc Details订单
	ood.Status = OrderCompleted
	ood.OrderNumber = RandOtcNumberSalt()
	ood.OrderTime = time.Now()

	_, err = session.Insert(ood)
	if err != nil {
		session.Rollback()
		return ood, err
	}

	//提交事务
	if err = session.Commit(); err != nil {
		return ood, err
	}

	//发送消息通知
	err = orderInfo.SendMessage(ood.JoinUserId)
	return ood, err
}

//取消订单 - 法币
func CancelBuySellOrder(orderDetailId int64) (*OtcOrderDetails, error) {
	orderDetails, err := GetOrderDetailsById(orderDetailId)
	if err != nil || orderDetails.JoinUserId == 0 {
		return orderDetails, err
	}

	if orderDetails.Status != OrderCreated {
		return orderDetails, errors.New("Current order cannot canceled.")
	}

	orderDetails.Status = OrderCanceled
	orderDetails.UpdateTime = time.Now()

	orderInfo, err := QueryOtcOrderById(orderDetails.OtcOrderId)
	if err != nil || orderInfo.UserId == 0 {
		return orderDetails, err
	}

	//事务更新
	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	//开始事务
	session.Begin()

	//恢复库存
	orderInfo.LeftAmount += orderDetails.BuyAmount
	orderInfo.UpdateTime = time.Now()
	_, err = session.Where("id=? ", orderInfo.Id).Update(orderInfo)
	if err != nil {
		session.Rollback()
		return orderDetails, err
	}

	//更新状态
	_, err = session.Where("id=? ", orderDetails.Id).Cols("status", "update_time").Update(orderDetails)
	if err != nil {
		session.Rollback()
		return orderDetails, err
	}

	//提交事务
	err = session.Commit()

	return orderDetails, err
}

//标记已付款
func MarkOrderPayment(orderDetailId int64) (*OtcOrderDetails, error) {
	orderDetails, err := GetOrderDetailsById(orderDetailId)
	if err != nil || orderDetails.JoinUserId == 0 {
		return orderDetails, err
	}

	orderDetails.JoinerPay = TrueValue
	orderDetails.UpdateTime = time.Now()

	_, err = bxgo.OrmEngin.Id(orderDetailId).Cols("buyer_pay", "update_time").Update(orderDetails)
	return orderDetails, err
}

//标记已转币
func MarkOrderCoinPayment(orderDetailId int64) (*OtcOrderDetails, error) {
	orderDetails, err := GetOrderDetailsById(orderDetailId)
	if err != nil || orderDetails.JoinUserId == 0 {
		return orderDetails, err
	}

	orderDetails.OwnerPay = TrueValue
	orderDetails.UpdateTime = time.Now()

	_, err = bxgo.OrmEngin.Id(orderDetailId).Cols("owner_pay", "update_time").Update(orderDetails)
	return orderDetails, err
}

//订单申诉
func ComplainOrderPayment(orderDetailId int64, who int, message string) (*OtcOrderDetails, error) {
	orderDetails, err := GetOrderDetailsById(orderDetailId)
	if err != nil || orderDetails.JoinUserId == 0 {
		return orderDetails, err
	}

	if who == Initiator {
		orderDetails.AppealOwner = TrueValue
		orderDetails.AppealOwnerMessage = message
	} else {
		orderDetails.AppealJoiner = TrueValue
		orderDetails.AppealJoinerMessage = message
	}

	orderDetails.Status = OrderAppealed
	orderDetails.UpdateTime = time.Now()

	_, err = bxgo.OrmEngin.Id(orderDetailId).Update(orderDetails)

	return orderDetails, err
}

//生成随机订单号
func RandOtcNumberSalt() string {
	baseStr := "0123456789"
	bytes := []byte(baseStr)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 8; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

//更新订单超时未付款状态
func UpdateOrderDetailsStatus() error {
	orderList, err := QueryOtcOrdersByFilter(OtcOrderFilter{Status:Started})
	if err != nil {
		fmt.Errorf("Query Otc Order info failed: %s", err.Error())
	}
	for _, order := range orderList {
		orderDetailsList, err := GetOrderDetailsByOtcOrderId(order.Id)
		if err != nil {
			fmt.Errorf("Query Order details collection info failed: %s", err.Error())
			continue
		}
		for _, detail := range orderDetailsList {
			minutes, _ := time.ParseDuration(fmt.Sprintf("%sm", order.ExpireTime))
			if detail.Status == OrderCreated && detail.OrderTime.Add(minutes).Before(time.Now()) {
				detail.Status = OrderTimedOut
				detail.UpdateTime = time.Now()

				_, err := bxgo.OrmEngin.Id(detail.Id).Cols("status").Update(detail)
				if err != nil {
					fmt.Errorf("Update Order details collection status info failed: %s", err.Error())
				}
			}
		}

	}

	return nil
}
