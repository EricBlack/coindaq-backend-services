package controller

import (
	"bx.com/project-service/proto"
	"bx.com/project-service/models"
	"bx.com/project-service/utils"

	"context"
	"fmt"
	"bx.com/user-service/model"
	"errors"
)

type OtcController struct{}

//添加支付方式 - POST
func (ctrl *OtcController) AddPaymentAccount(ctx context.Context, in *proto.PaymentAccountReq) (*proto.Empty, error) {
	err := models.AddPayAccount(in.UserId, in.AccountName, in.AccountNumber, in.OpenBank, in.SubBank, in.PaymentPassword, int(in.AccountType))

	return &proto.Empty{}, err
}

//支付账号删除 - POST
func (ctrl *OtcController) DeletePaymentAccount(ctx context.Context, in *proto.AccountFilterReq) (*proto.Empty, error) {
	err := models.DeletePayAccount(in.UserId, int(in.AccountType))

	return &proto.Empty{}, err
}

//支付账号查询 - GET
func (ctrl *OtcController) QueryPaymentAccount(ctx context.Context, in *proto.AccountFilterReq) (*proto.AccountListReply, error) {
	accountListResult := proto.AccountListReply{}
	accoutList, err := models.QueryUserPayAccountByFilter(in.UserId, int(in.AccountType))
	if err !=nil {
		fmt.Errorf("%s", "Query user pay account got errr : %s", err.Error())
		return &accountListResult, err
	}
	if len(accoutList) == 0 {
		return &accountListResult, nil
	}

	for _, account := range accoutList {
		accountListResult.AccountList = append(accountListResult.AccountList, &proto.AccountReply{
			Id:				account.Id,
			UserId:			account.UserId,
			AccountName:	account.AccountName,
			AccountType:	proto.AccountType(account.AccountType),
			AccountNumber:  account.AccountNumber,
			OpenBank:		account.OpenBank,
			SubBank:		account.SubBank,
			IsDeleted:		proto.BoolValue(account.IsDeleted),
			CreateTime:		utils.Time2String(account.CreateTime),
			UpdateTime:		utils.Time2String(account.UpdateTime),
		})
	}

	return &accountListResult, nil
}

//消息添加 - POST
func (ctrl *OtcController) AddShortMessage(ctx context.Context, in *proto.MessageReq) (*proto.Empty, error) {
	err := models.SendMessage(in.UserFrom, in.UserTo, in.Message)

	return &proto.Empty{}, err
}

//消息列表查询 - GET
func (ctrl *OtcController) QueryShortMessages(ctx context.Context, in *proto.UserPairReq) (*proto.MessageListReply, error) {
	messageList, err := models.QueryUserMessage(in.FromId, in.ToId, in.LastId)
	if err != nil {
		return &proto.MessageListReply{}, err
	}
	messageResults := proto.MessageListReply{}
	for _, message := range messageList {
		messageResults.MessageList = append(messageResults.MessageList, &proto.MessageReply{
			Id:			message.Id,
			UserFrom:	message.UserFrom,
			UserTo:		message.UserTo,
			Message:	message.Message,
			CreateTime:	utils.Time2String(message.CreateTime),
		})
	}

	return &messageResults, nil
}

//交易对信息查询 - GET
func (ctrl *OtcController) QueryCurrencyPairs(ctx context.Context, in *proto.PairFilterReq) (*proto.CurrencyPairListReply, error) {
	currencyList, err := models.QueryCurrencyPairByFilter(in.ACurrencyName, in.BCurrencyName, int(in.TypeInfo), int(in.MatchAll))
	if err != nil || len(currencyList) ==0 {
		return &proto.CurrencyPairListReply{}, err
	}

	currencyResults := proto.CurrencyPairListReply{}
	for _, currency := range currencyList {
		currencyResults.PairList = append(currencyResults.PairList, &proto.CurrencyPairReply{
			Id:				currency.Id,
			CoinaId:		currency.CoinaId,
			CoinaName:		currency.CoinaName,
			CoinbId:		currency.CoinbId,
			CoinbName:		currency.CoinbName,
			PrioritySort:	int32(currency.PrioritySort),
			Type:			proto.CurrencyType(currency.Type),
		})
	}

	return &proto.CurrencyPairListReply{}, nil
}

//创建Otc订单 - POST
func (ctrl *OtcController) CreateOtcOrder(ctx context.Context, in *proto.OtcOrderReq) (*proto.Empty, error) {
	otcOrder := models.OtcOrder{
		UserId:				in.UserId,
		BuySell:			int(in.BuySell),
		CurrencyPair:		in.CurrencyPair,
		QuotePrice:			in.QuotePrice,
		QuoteAmount:		in.QuoteAmount,
		MinValue:			in.MinValue,
		MaxValue:			in.MaxValue,
		ExpireTime:			int(in.ExpireTime),
		OpenLimitHour:		int(in.OpenLimitHour),
		OpenLimitMinute:	int(in.OpenLimitMinute),
		CloseLimitHour:		int(in.CloseLimitHour),
		CloseLimitMinute:	int(in.CloseLimitMinute),
		AutoReply:			int(in.AutoReply),
		AutoMessage:		in.AutoMessage,
	}

	err := otcOrder.CreateOtcOrder()

	return &proto.Empty{}, err
}

//查询广告单列表信息 - GET
func (ctrl *OtcController) QueryOtcOrderList(ctx context.Context, in *proto.OtcFilterReq) (*proto.OtcOrderListReply, error) {
	otcFilter := models.OtcOrderFilter{
		UserId:			in.UserId,
		BuySell:		int(in.BuySell),
		CurrencyPair:	in.CurrencyPairId,
		AvailableTime:	int(in.OpenState),
		Status:			int(in.OtcStatus),
	}
	otcOrderList, err := models.QueryOtcOrdersByFilter(otcFilter)
	if err != nil {
		return  &proto.OtcOrderListReply{}, err
	}

	otcOrders := proto.OtcOrderListReply{}
	for _, otcOrder := range otcOrderList {
		//查询CurrencyPair
		currency, err := models.QueryCurrencyPairById(otcOrder.CurrencyPair)
		if err != nil || currency.CoinaId == "" {
			return &proto.OtcOrderListReply{}, err
		}
		currencyResult := &proto.CurrencyPairReply{
			Id:				currency.Id,
			CoinaId:		currency.CoinaId,
			CoinaName:  	currency.CoinaName,
			CoinbId:		currency.CoinbId,
			CoinbName:		currency.CoinbName,
			PrioritySort:	int32(currency.PrioritySort),
			Type:			proto.CurrencyType(currency.Type),
		}

		//查询收款地址信息
		accountListResult := proto.AccountListReply{}
		payeeList, err := models.QueryOtcPayAddress(otcOrder.Id)
		if err != nil {
			fmt.Errorf("Query Otc Pay address got error: %s", err.Error())
			return &proto.OtcOrderListReply{}, err
		}
		if len(payeeList) == 0 {
			fmt.Errorf("Query Otc Pay address return null.")
		} else {
			for _, payee := range payeeList {
				accountListResult.AccountList = append(accountListResult.AccountList, &proto.AccountReply{
					Id:				payee.Id,
					UserId:			payee.UserId,
					AccountName:	payee.AccountName,
					AccountType:	proto.AccountType(payee.AccountType),
					AccountNumber:  payee.AccountNumber,
					OpenBank:		payee.OpenBank,
					SubBank:		payee.SubBank,
					IsDeleted:		proto.BoolValue(payee.IsDeleted),
					CreateTime:		utils.Time2String(payee.CreateTime),
					UpdateTime:		utils.Time2String(payee.UpdateTime),
				})
			}
		}

		otcOrders.OrderList = append(otcOrders.OrderList, &proto.OtcOrderReply{
			Id:					otcOrder.Id,
			UserId:				otcOrder.UserId,
			BuySell:			proto.TradeType(otcOrder.BuySell),
			CurrencyPair:		currencyResult,
			PayeeAddressList:	&accountListResult,
			QuotePrice:			otcOrder.QuotePrice,
			QuoteAmount:		otcOrder.QuoteAmount,
			LeftAmount:			otcOrder.LeftAmount,
			MinValue:			otcOrder.MinValue,
			MaxValue:			otcOrder.MaxValue,
			ExpireTime:			int32(otcOrder.ExpireTime),
			OpenLimitHour:		int32(otcOrder.OpenLimitHour),
			OpenLimitMinute:	int32(otcOrder.OpenLimitMinute),
			CloseLimitHour:		int32(otcOrder.CloseLimitHour),
			CloseLimitMinute:	int32(otcOrder.CloseLimitMinute),
			Status:				proto.OtcStatusType(otcOrder.Status),
			AutoReply:			proto.BoolValue(otcOrder.AutoReply),
			AutoMessage:		otcOrder.AutoMessage,
			CreateTime:			utils.Time2String(otcOrder.CreateTime),
			UpdateTime:			utils.Time2String(otcOrder.UpdateTime),
		})
	}

	return &proto.OtcOrderListReply{}, nil
}

//查询广告单详细信息 - GET
func (ctrl *OtcController) QueryOtcOrderDetails(ctx context.Context, in *proto.IdReq) (*proto.OtcOrderReply, error) {
	otcOrder, err := models.QueryOtcOrderById(in.Id)
	if err != nil || otcOrder.UserId == 0 {
		return &proto.OtcOrderReply{}, err
	}

	//查询CurrencyPair
	currency, err := models.QueryCurrencyPairById(otcOrder.CurrencyPair)
	if err != nil || currency.CoinaId == "" {
		return &proto.OtcOrderReply{}, err
	}
	currencyResult := &proto.CurrencyPairReply{
		Id:				currency.Id,
		CoinaId:		currency.CoinaId,
		CoinaName:  	currency.CoinaName,
		CoinbId:		currency.CoinbId,
		CoinbName:		currency.CoinbName,
		PrioritySort:	int32(currency.PrioritySort),
		Type:			proto.CurrencyType(currency.Type),
	}

	//查询收款地址信息
	accountListResult := proto.AccountListReply{}
	payeeList, err := models.QueryOtcPayAddress(otcOrder.Id)
	if err != nil {
		fmt.Errorf("Query Otc Pay address got error: %s", err.Error())
		return &proto.OtcOrderReply{}, err
	}
	if len(payeeList) == 0 {
		fmt.Errorf("Query Otc Pay address return null.")
	} else {
		for _, payee := range payeeList {
			accountListResult.AccountList = append(accountListResult.AccountList, &proto.AccountReply{
				Id:				payee.Id,
				UserId:			payee.UserId,
				AccountName:	payee.AccountName,
				AccountType:	proto.AccountType(payee.AccountType),
				AccountNumber:  payee.AccountNumber,
				OpenBank:		payee.OpenBank,
				SubBank:		payee.SubBank,
				IsDeleted:		proto.BoolValue(payee.IsDeleted),
				CreateTime:		utils.Time2String(payee.CreateTime),
				UpdateTime:		utils.Time2String(payee.UpdateTime),
			})
		}
	}

	return &proto.OtcOrderReply{
		Id:					otcOrder.Id,
		UserId:				otcOrder.UserId,
		BuySell:			proto.TradeType(otcOrder.BuySell),
		CurrencyPair:		currencyResult,
		PayeeAddressList:	&accountListResult,
		QuotePrice:			otcOrder.QuotePrice,
		QuoteAmount:		otcOrder.QuoteAmount,
		LeftAmount:			otcOrder.LeftAmount,
		MinValue:			otcOrder.MinValue,
		MaxValue:			otcOrder.MaxValue,
		ExpireTime:			int32(otcOrder.ExpireTime),
		OpenLimitHour:		int32(otcOrder.OpenLimitHour),
		OpenLimitMinute:	int32(otcOrder.OpenLimitMinute),
		CloseLimitHour:		int32(otcOrder.CloseLimitHour),
		CloseLimitMinute:	int32(otcOrder.CloseLimitMinute),
		Status:				proto.OtcStatusType(otcOrder.Status),
		AutoReply:			proto.BoolValue(otcOrder.AutoReply),
		AutoMessage:		otcOrder.AutoMessage,
		CreateTime:			utils.Time2String(otcOrder.CreateTime),
		UpdateTime:			utils.Time2String(otcOrder.UpdateTime),
	}, nil
}

//确认下单交易 - POST
func (ctrl *OtcController) ConfirmDealOrder(ctx context.Context, in *proto.OtcDetailReq) (*proto.OtcDetailsReply, error){
	//查询广告单
	otcOrder, err := models.QueryOtcOrderById(in.OtcOrderId)
	if err != nil || otcOrder.UserId == 0 {
		return &proto.OtcDetailsReply{}, err
	}
	//查询交易对信息
	currencyPair, err := models.QueryCurrencyPairById(otcOrder.CurrencyPair)
	if err != nil || currencyPair.Type == 0 {
		return &proto.OtcDetailsReply{}, err
	}

	//添加记录单
	orderDetails := models.OtcOrderDetails{
		OtcOrderId:			otcOrder.Id,
		JoinUserId:			in.UserId,
		BuyAmount:			in.BuyAmount,
		PayAmount:			in.PayAmount,
	}

	var detailsResult *models.OtcOrderDetails
	//法币交易
	if currencyPair.Type == models.FiatDeal {
		detailsResult, err = orderDetails.ConfirmCoinDealOrder()
		if err != nil {
			return &proto.OtcDetailsReply{}, err
		}
	}

	//币币交易
	if currencyPair.Type == models.CoinDeail {
		detailsResult, err = orderDetails.ConfirmCoinDealOrder()
		if err != nil {
			return &proto.OtcDetailsReply{}, err
		}
	}

	//转换输出
	otcDetailsResult := proto.OtcDetailsReply{}

	otcDetailsResult.Id = detailsResult.Id
	otcDetailsResult.JoinUserId = detailsResult.JoinUserId
	otcDetailsResult.BuyAmount = detailsResult.BuyAmount
	otcDetailsResult.PayAmount = detailsResult.PayAmount
	otcDetailsResult.OrderTime = utils.Time2String(detailsResult.OrderTime)
	otcDetailsResult.SettlementTime = utils.Time2String(detailsResult.SettlementTime)
	otcDetailsResult.SettlementNote = detailsResult.SettlementNote
	otcDetailsResult.Status = proto.OtcDetailsStatusType(detailsResult.Status)
	otcDetailsResult.OwnerPay = proto.BoolValue(detailsResult.OwnerPay)
	otcDetailsResult.JoinerPay = proto.BoolValue(detailsResult.JoinerPay)
	otcDetailsResult.AppealOwner = proto.BoolValue(detailsResult.AppealOwner)
	otcDetailsResult.AppealOwnerMessage = detailsResult.AppealOwnerMessage
	otcDetailsResult.AppealJoiner = proto.BoolValue(detailsResult.AppealJoiner)
	otcDetailsResult.AppealJoinerMessage = detailsResult.AppealJoinerMessage
	otcDetailsResult.OrderNumber = detailsResult.OrderNumber
	otcDetailsResult.UpdateTime = utils.Time2String(detailsResult.UpdateTime)

	//查询CurrencyPair
	currency, err := models.QueryCurrencyPairById(otcOrder.CurrencyPair)
	if err != nil || currency.CoinaId == "" {
		return &otcDetailsResult, err
	}
	currencyResult := &proto.CurrencyPairReply{
		Id:				currency.Id,
		CoinaId:		currency.CoinaId,
		CoinaName:  	currency.CoinaName,
		CoinbId:		currency.CoinbId,
		CoinbName:		currency.CoinbName,
		PrioritySort:	int32(currency.PrioritySort),
		Type:			proto.CurrencyType(currency.Type),
	}

	otcOrderResult := proto.OtcOrderReply {
		Id:					otcOrder.Id,
		UserId:				otcOrder.UserId,
		BuySell:			proto.TradeType(otcOrder.BuySell),
		CurrencyPair:		currencyResult,
		PayeeAddressList:	nil,
		QuotePrice:			otcOrder.QuotePrice,
		QuoteAmount:		otcOrder.QuoteAmount,
		LeftAmount:			otcOrder.LeftAmount,
		MinValue:			otcOrder.MinValue,
		MaxValue:			otcOrder.MaxValue,
		ExpireTime:			int32(otcOrder.ExpireTime),
		OpenLimitHour:		int32(otcOrder.OpenLimitHour),
		OpenLimitMinute:	int32(otcOrder.OpenLimitMinute),
		CloseLimitHour:		int32(otcOrder.CloseLimitHour),
		CloseLimitMinute:	int32(otcOrder.CloseLimitMinute),
		Status:				proto.OtcStatusType(otcOrder.Status),
		AutoReply:			proto.BoolValue(otcOrder.AutoReply),
		AutoMessage:		otcOrder.AutoMessage,
		CreateTime:			utils.Time2String(otcOrder.CreateTime),
		UpdateTime:			utils.Time2String(otcOrder.UpdateTime),
	}
	otcDetailsResult.OtcOrder = &otcOrderResult

	//查询商家用户信息
	ownerInfo, err := model.QueryUserById(otcOrderResult.Id)
	if err != nil || ownerInfo.Email == "" {
		otcDetailsResult.OrderOwnerName = ""
		return &otcDetailsResult, err
	} else {
		otcDetailsResult.OrderOwnerName = ownerInfo.Email
	}

	//查询参与者用户信息
	joinerInfo, err := model.QueryUserById(detailsResult.JoinUserId)
	if err != nil || joinerInfo.Email == "" {
		otcDetailsResult.OrderJoinerName = ""
		return &otcDetailsResult, err
	} else {
		otcDetailsResult.OrderJoinerName = joinerInfo.Email
	}

	return &otcDetailsResult, err
}

//取消订单 - POST
func (ctrl *OtcController) CancelDealOrder(ctx context.Context, in *proto.IdReq) (*proto.Empty, error) {
	_, err := models.CancelBuySellOrder(in.Id)

	return &proto.Empty{}, err
}

//确认已付款 - POST
func (ctrl *OtcController) MarkOrderPayment(ctx context.Context, in *proto.MarkUserPaiedReq) (*proto.Empty, error) {
	if in.PersonType == proto.OrderPersonType_InitiatorType {
		_, err := models.MarkOrderPayment(in.OrderDetailId)

		return &proto.Empty{}, err
	}
	if in.PersonType == proto.OrderPersonType_AcceptorType {
		_, err := models.MarkOrderCoinPayment(in.OrderDetailId)

		return &proto.Empty{}, err
	}else {
		return &proto.Empty{}, errors.New("Wrong type of order person type.")
	}
}

//订单申诉 - POST
func (ctrl *OtcController) ComplainOrderPayment(ctx context.Context, in *proto.ComplainReq) (*proto.Empty, error) {
	if in.OrderDetailId == 0 {
		return &proto.Empty{}, errors.New("Order details id should not be zero.")
	}
	if in.PersonType != proto.OrderPersonType_InitiatorType && in.PersonType != proto.OrderPersonType_AcceptorType {
		return &proto.Empty{}, errors.New("Order person type parameter is not correct.")
	}
	_, err := models.ComplainOrderPayment(in.OrderDetailId, int(in.PersonType), in.Message)
	return &proto.Empty{}, err
}

//查询交易订单详情 - GET
func (ctrl *OtcController) QueryOtcDetailsById(ctx context.Context, in *proto.IdReq) (*proto.OtcDetailsReply, error) {
	otcDetailsResult := proto.OtcDetailsReply{}
	//查询订单信息
	otcDetails, err := models.GetOrderDetailsById(in.Id)
	if err != nil || otcDetails.JoinUserId == 0 {
		return &otcDetailsResult, err
	}
	otcDetailsResult.Id = otcDetails.Id
	otcDetailsResult.JoinUserId = otcDetails.JoinUserId
	otcDetailsResult.BuyAmount = otcDetails.BuyAmount
	otcDetailsResult.PayAmount = otcDetails.PayAmount
	otcDetailsResult.OrderTime = utils.Time2String(otcDetails.OrderTime)
	otcDetailsResult.SettlementTime = utils.Time2String(otcDetails.SettlementTime)
	otcDetailsResult.SettlementNote = otcDetails.SettlementNote
	otcDetailsResult.Status = proto.OtcDetailsStatusType(otcDetails.Status)
	otcDetailsResult.OwnerPay = proto.BoolValue(otcDetails.OwnerPay)
	otcDetailsResult.JoinerPay = proto.BoolValue(otcDetails.JoinerPay)
	otcDetailsResult.AppealOwner = proto.BoolValue(otcDetails.AppealOwner)
	otcDetailsResult.AppealOwnerMessage = otcDetails.AppealOwnerMessage
	otcDetailsResult.AppealJoiner = proto.BoolValue(otcDetails.AppealJoiner)
	otcDetailsResult.AppealJoinerMessage = otcDetails.AppealJoinerMessage
	otcDetailsResult.OrderNumber = otcDetails.OrderNumber
	otcDetailsResult.UpdateTime = utils.Time2String(otcDetails.UpdateTime)

	//查询广告单
	otcOrder, err := models.QueryOtcOrderById(otcDetails.OtcOrderId)
	if err != nil || otcOrder.UserId == 0 {
		return &otcDetailsResult, err
	}

	//查询CurrencyPair
	currency, err := models.QueryCurrencyPairById(otcOrder.CurrencyPair)
	if err != nil || currency.CoinaId == "" {
		return &otcDetailsResult, err
	}
	currencyResult := &proto.CurrencyPairReply{
		Id:				currency.Id,
		CoinaId:		currency.CoinaId,
		CoinaName:  	currency.CoinaName,
		CoinbId:		currency.CoinbId,
		CoinbName:		currency.CoinbName,
		PrioritySort:	int32(currency.PrioritySort),
		Type:			proto.CurrencyType(currency.Type),
	}

	otcOrderResult := proto.OtcOrderReply {
		Id:					otcOrder.Id,
		UserId:				otcOrder.UserId,
		BuySell:			proto.TradeType(otcOrder.BuySell),
		CurrencyPair:		currencyResult,
		PayeeAddressList:	nil,
		QuotePrice:			otcOrder.QuotePrice,
		QuoteAmount:		otcOrder.QuoteAmount,
		LeftAmount:			otcOrder.LeftAmount,
		MinValue:			otcOrder.MinValue,
		MaxValue:			otcOrder.MaxValue,
		ExpireTime:			int32(otcOrder.ExpireTime),
		OpenLimitHour:		int32(otcOrder.OpenLimitHour),
		OpenLimitMinute:	int32(otcOrder.OpenLimitMinute),
		CloseLimitHour:		int32(otcOrder.CloseLimitHour),
		CloseLimitMinute:	int32(otcOrder.CloseLimitMinute),
		Status:				proto.OtcStatusType(otcOrder.Status),
		AutoReply:			proto.BoolValue(otcOrder.AutoReply),
		AutoMessage:		otcOrder.AutoMessage,
		CreateTime:			utils.Time2String(otcOrder.CreateTime),
		UpdateTime:			utils.Time2String(otcOrder.UpdateTime),
	}
	otcDetailsResult.OtcOrder = &otcOrderResult

	//查询商家用户信息
	ownerInfo, err := model.QueryUserById(otcOrderResult.Id)
	if err != nil || ownerInfo.Email == "" {
		otcDetailsResult.OrderOwnerName = ""
		return &otcDetailsResult, err
	} else {
		otcDetailsResult.OrderOwnerName = ownerInfo.Email
	}

	//查询参与者用户信息
	joinerInfo, err := model.QueryUserById(otcDetails.JoinUserId)
	if err != nil || joinerInfo.Email == "" {
		otcDetailsResult.OrderJoinerName = ""
		return &otcDetailsResult, err
	} else {
		otcDetailsResult.OrderJoinerName = joinerInfo.Email
	}

	return &otcDetailsResult, err
}

//查询广告单关联订单信息 - GET
func (ctrl *OtcController) QueryOtcOrderDetailsListByOtcOrderId(ctx context.Context, in *proto.IdReq) (*proto.OtcDetailsListReply, error) {
	otcDetailsList := proto.OtcDetailsListReply{}
	//查询订单列表
	detailsList, err := models.GetOrderDetailsByOtcOrderId(in.Id)
	if err != nil || len(detailsList) == 0 {
		return &otcDetailsList, err
	}

	for _, otcDetails := range detailsList {
		otcDetailsResult := proto.OtcDetailsReply{}
		otcDetailsResult.Id = otcDetails.Id
		otcDetailsResult.JoinUserId = otcDetails.JoinUserId
		otcDetailsResult.BuyAmount = otcDetails.BuyAmount
		otcDetailsResult.PayAmount = otcDetails.PayAmount
		otcDetailsResult.OrderTime = utils.Time2String(otcDetails.OrderTime)
		otcDetailsResult.SettlementTime = utils.Time2String(otcDetails.SettlementTime)
		otcDetailsResult.SettlementNote = otcDetails.SettlementNote
		otcDetailsResult.Status = proto.OtcDetailsStatusType(otcDetails.Status)
		otcDetailsResult.OwnerPay = proto.BoolValue(otcDetails.OwnerPay)
		otcDetailsResult.JoinerPay = proto.BoolValue(otcDetails.JoinerPay)
		otcDetailsResult.AppealOwner = proto.BoolValue(otcDetails.AppealOwner)
		otcDetailsResult.AppealOwnerMessage = otcDetails.AppealOwnerMessage
		otcDetailsResult.AppealJoiner = proto.BoolValue(otcDetails.AppealJoiner)
		otcDetailsResult.AppealJoinerMessage = otcDetails.AppealJoinerMessage
		otcDetailsResult.OrderNumber = otcDetails.OrderNumber
		otcDetailsResult.UpdateTime = utils.Time2String(otcDetails.UpdateTime)

		//查询广告单
		otcOrder, err := models.QueryOtcOrderById(otcDetails.OtcOrderId)
		if err != nil || otcOrder.UserId == 0 {
			return &otcDetailsList, err
		}

		//查询CurrencyPair
		currency, err := models.QueryCurrencyPairById(otcOrder.CurrencyPair)
		if err != nil || currency.CoinaId == "" {
			return &otcDetailsList, err
		}
		currencyResult := &proto.CurrencyPairReply{
			Id:				currency.Id,
			CoinaId:		currency.CoinaId,
			CoinaName:  	currency.CoinaName,
			CoinbId:		currency.CoinbId,
			CoinbName:		currency.CoinbName,
			PrioritySort:	int32(currency.PrioritySort),
			Type:			proto.CurrencyType(currency.Type),
		}

		otcOrderResult := proto.OtcOrderReply {
			Id:					otcOrder.Id,
			UserId:				otcOrder.UserId,
			BuySell:			proto.TradeType(otcOrder.BuySell),
			CurrencyPair:		currencyResult,
			PayeeAddressList:	nil,
			QuotePrice:			otcOrder.QuotePrice,
			QuoteAmount:		otcOrder.QuoteAmount,
			LeftAmount:			otcOrder.LeftAmount,
			MinValue:			otcOrder.MinValue,
			MaxValue:			otcOrder.MaxValue,
			ExpireTime:			int32(otcOrder.ExpireTime),
			OpenLimitHour:		int32(otcOrder.OpenLimitHour),
			OpenLimitMinute:	int32(otcOrder.OpenLimitMinute),
			CloseLimitHour:		int32(otcOrder.CloseLimitHour),
			CloseLimitMinute:	int32(otcOrder.CloseLimitMinute),
			Status:				proto.OtcStatusType(otcOrder.Status),
			AutoReply:			proto.BoolValue(otcOrder.AutoReply),
			AutoMessage:		otcOrder.AutoMessage,
			CreateTime:			utils.Time2String(otcOrder.CreateTime),
			UpdateTime:			utils.Time2String(otcOrder.UpdateTime),
		}
		otcDetailsResult.OtcOrder = &otcOrderResult

		//查询商家用户信息
		ownerInfo, err := model.QueryUserById(otcOrderResult.Id)
		if err != nil || ownerInfo.Email == "" {
			otcDetailsResult.OrderOwnerName = ""
			return &otcDetailsList, err
		} else {
			otcDetailsResult.OrderOwnerName = ownerInfo.Email
		}

		//查询参与者用户信息
		joinerInfo, err := model.QueryUserById(otcDetails.JoinUserId)
		if err != nil || joinerInfo.Email == "" {
			otcDetailsResult.OrderJoinerName = ""
			return &otcDetailsList, err
		} else {
			otcDetailsResult.OrderJoinerName = joinerInfo.Email
		}
		otcDetailsList.DetailList = append(otcDetailsList.DetailList, &otcDetailsResult)
	}

	return &otcDetailsList, nil
}

//查询用户参与订单信息 - GET
func(ctrl *OtcController) QueryUserOrderDetailsListByFilter(ctx context.Context, in *proto.OtcDetailsFilterReq) (*proto.OtcDetailsListReply, error) {
	otcDetailsList := proto.OtcDetailsListReply{}
	var detailsList []*models.OtcOrderDetails
	var err error
	//查询订单列表
	if in.State == proto.OrderStateType_OngoingType {
		detailsList, err = models.GetUserOngoingOrderDetails(in.UserId)
	} else if in.State == proto.OrderStateType_CompleteType {
		detailsList, err = models.GetUserCompleteOrderDetails(in.UserId)
	} else if in.State == proto.OrderStateType_CancelType {
		detailsList, err = models.GetUserCancelOrderDetails(in.UserId)
	}
	if err != nil || len(detailsList) == 0 {
		return &otcDetailsList, err
	}

	for _, otcDetails := range detailsList {
		otcDetailsResult := proto.OtcDetailsReply{}
		otcDetailsResult.Id = otcDetails.Id
		otcDetailsResult.JoinUserId = otcDetails.JoinUserId
		otcDetailsResult.BuyAmount = otcDetails.BuyAmount
		otcDetailsResult.PayAmount = otcDetails.PayAmount
		otcDetailsResult.OrderTime = utils.Time2String(otcDetails.OrderTime)
		otcDetailsResult.SettlementTime = utils.Time2String(otcDetails.SettlementTime)
		otcDetailsResult.SettlementNote = otcDetails.SettlementNote
		otcDetailsResult.Status = proto.OtcDetailsStatusType(otcDetails.Status)
		otcDetailsResult.OwnerPay = proto.BoolValue(otcDetails.OwnerPay)
		otcDetailsResult.JoinerPay = proto.BoolValue(otcDetails.JoinerPay)
		otcDetailsResult.AppealOwner = proto.BoolValue(otcDetails.AppealOwner)
		otcDetailsResult.AppealOwnerMessage = otcDetails.AppealOwnerMessage
		otcDetailsResult.AppealJoiner = proto.BoolValue(otcDetails.AppealJoiner)
		otcDetailsResult.AppealJoinerMessage = otcDetails.AppealJoinerMessage
		otcDetailsResult.OrderNumber = otcDetails.OrderNumber
		otcDetailsResult.UpdateTime = utils.Time2String(otcDetails.UpdateTime)

		//查询广告单
		otcOrder, err := models.QueryOtcOrderById(otcDetails.OtcOrderId)
		if err != nil || otcOrder.UserId == 0 {
			return &otcDetailsList, err
		}

		//查询CurrencyPair
		currency, err := models.QueryCurrencyPairById(otcOrder.CurrencyPair)
		if err != nil || currency.CoinaId == "" {
			return &otcDetailsList, err
		}
		currencyResult := &proto.CurrencyPairReply{
			Id:				currency.Id,
			CoinaId:		currency.CoinaId,
			CoinaName:  	currency.CoinaName,
			CoinbId:		currency.CoinbId,
			CoinbName:		currency.CoinbName,
			PrioritySort:	int32(currency.PrioritySort),
			Type:			proto.CurrencyType(currency.Type),
		}

		otcOrderResult := proto.OtcOrderReply {
			Id:					otcOrder.Id,
			UserId:				otcOrder.UserId,
			BuySell:			proto.TradeType(otcOrder.BuySell),
			CurrencyPair:		currencyResult,
			PayeeAddressList:	nil,
			QuotePrice:			otcOrder.QuotePrice,
			QuoteAmount:		otcOrder.QuoteAmount,
			LeftAmount:			otcOrder.LeftAmount,
			MinValue:			otcOrder.MinValue,
			MaxValue:			otcOrder.MaxValue,
			ExpireTime:			int32(otcOrder.ExpireTime),
			OpenLimitHour:		int32(otcOrder.OpenLimitHour),
			OpenLimitMinute:	int32(otcOrder.OpenLimitMinute),
			CloseLimitHour:		int32(otcOrder.CloseLimitHour),
			CloseLimitMinute:	int32(otcOrder.CloseLimitMinute),
			Status:				proto.OtcStatusType(otcOrder.Status),
			AutoReply:			proto.BoolValue(otcOrder.AutoReply),
			AutoMessage:		otcOrder.AutoMessage,
			CreateTime:			utils.Time2String(otcOrder.CreateTime),
			UpdateTime:			utils.Time2String(otcOrder.UpdateTime),
		}
		otcDetailsResult.OtcOrder = &otcOrderResult

		//查询商家用户信息
		ownerInfo, err := model.QueryUserById(otcOrderResult.Id)
		if err != nil || ownerInfo.Email == "" {
			otcDetailsResult.OrderOwnerName = ""
			return &otcDetailsList, err
		} else {
			otcDetailsResult.OrderOwnerName = ownerInfo.Email
		}

		//查询参与者用户信息
		joinerInfo, err := model.QueryUserById(otcDetails.JoinUserId)
		if err != nil || joinerInfo.Email == "" {
			otcDetailsResult.OrderJoinerName = ""
			return &otcDetailsList, err
		} else {
			otcDetailsResult.OrderJoinerName = joinerInfo.Email
		}
		otcDetailsList.DetailList = append(otcDetailsList.DetailList, &otcDetailsResult)
	}

	return &otcDetailsList, nil
}