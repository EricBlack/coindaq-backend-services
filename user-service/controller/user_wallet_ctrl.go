package controller

import (
	"bx.com/user-service/model"
	"bx.com/user-service/proto"
	"bx.com/user-service/utils"
	"context"
	"errors"
)

type UserWalletControl struct{}

func (ctrl *UserWalletControl) BindUserCoinAddress(ctx context.Context, in *proto.UserCoinAddressReq) (*proto.Empty, error) {
	if in.UserId == 0 || in.CoinId == "" {
		return &proto.Empty{}, errors.New("Invalid request parameter.")
	}
	err := model.CreateUserCoinAddress(in.UserId, in.CoinId)
	return &proto.Empty{}, err
}

func (ctrl *UserWalletControl) QueryUserCoinAddress(ctx context.Context, in *proto.UserCoinAddressReq) (*proto.UserCoinListReply, error) {
	balanceList, err := model.QueryBalanceByFilter(in.UserId, in.CoinId) //需添加更新Coin地址转换
	if err != nil {
		return &proto.UserCoinListReply{}, err
	}
	if len(balanceList) != 0 {
		result := &proto.UserCoinListReply{}
		for _, ub := range balanceList {
			result.UserCoinList = append(result.UserCoinList, &proto.UserCoinAddressReply{
				BalanceId:       ub.BalanceId,
				UserId:          ub.UserId,
				CurrencyId:      ub.CurrencyId,
				BalanceValue:    ub.BalanceValue,
				RechargeAddress: ub.RechargeAddress,
				QRCodeAddress:   ub.QrcodeAddress,
			})
		}

		return result, nil
	}
	return &proto.UserCoinListReply{}, nil
}

func (ctrl *UserWalletControl) AddWalletAddress(ctx context.Context, in *proto.WalletAddressReq) (*proto.Empty, error) {
	wallet := model.UserWallet{
		UserId:        in.UserId,
		Currency:      in.CurrencyName,
		WalletAddress: in.WalletAddress,
	}
	err := wallet.AddWalletAddress()

	return &proto.Empty{}, err
}

func (ctrl *UserWalletControl) QueryUserWalletAddress(ctx context.Context, in *proto.WalletAddressReq) (*proto.WalletAddressListReply, error) {
	walletList, err := model.QueryUserWalletAddress(in.UserId, in.CurrencyName, in.WalletAddress)
	if err != nil {
		return &proto.WalletAddressListReply{}, err
	}
	if len(walletList) == 0 {
		return &proto.WalletAddressListReply{}, nil
	} else {
		resp := proto.WalletAddressListReply{}
		for _, wallet := range walletList {
			resp.AddressList = append(resp.AddressList, &proto.WalletAddressInfoReply{
				Id:            wallet.Id,
				UserId:        wallet.UserId,
				CurrencyName:  wallet.Currency,
				WalletAddress: wallet.WalletAddress,
				CreateTime:    utils.Time2String(wallet.CreateTime),
				UpdateTime:    utils.Time2String(wallet.UpdateTime),
			})
		}

		return &resp, nil
	}
}

func (ctrl *UserWalletControl) DeleteUserWalletAddress(ctx context.Context, in *proto.IdReq) (*proto.Empty, error) {
	err := model.DeleteUserWallet(in.Id)

	return &proto.Empty{}, err
}

//User Platform Balance
func (ctrl *UserWalletControl) QueryUserBalanceByFilter(ctx context.Context, in *proto.UserCoinAddressReq) (*proto.UserBalanceListReply, error) {
	userBalanceList, err := model.QueryBalanceByFilter(in.UserId, in.CoinId)
	if err != nil || len(userBalanceList) == 0 {
		return &proto.UserBalanceListReply{}, nil
	} else {
		balanceResult := proto.UserBalanceListReply{}
		for _, balance := range userBalanceList {
			balanceResult.BalanceList = append(balanceResult.BalanceList, &proto.UserBalanceReply{
				BalanceId:			balance.BalanceId,
				UserId:				balance.UserId,
				CurrencyId:			balance.CurrencyId,
				BalanceValue:		balance.BalanceValue,
				RechargeAddress:	balance.RechargeAddress,
				QrcodeAddress:		balance.QrcodeAddress,
				TotalBalance:		balance.TotalBalance,
				ChargeUnAccount:    balance.ChargeUnAccount,
				WithdrawUnTransfer: balance.WithdrawUnTransfer,
				LockPosition:		balance.LockPosition,
				IcoUndue:			balance.IcoUndue,
			})
		}

		return &balanceResult, nil
	}
}