package controller

import (
	"context"

	"bx.com/platform-wallet/model"
	"bx.com/platform-wallet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LightWalletRpc struct{}

func (lwrpc *LightWalletRpc) SyncLightWallet(context context.Context, in *proto.SyncLightWalletRequest) (*proto.Empty, error) {
	for _, coin := range in.Coins {
		if coin.Symbol == "" {
			return nil, status.Error(codes.InvalidArgument, "symbol can not be empty")
		}
		coinInfo, err := model.GetCoinInfoBySymbol(coin.Symbol)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		lwCoin := model.LightWalletCoin{
			UserId:     in.UserId,
			WalletName: in.WalletName,
			CoinId:     coinInfo.Id,
			Address:    coin.Address,
		}
		err = model.CreateLightWalletCoin(&lwCoin)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &proto.Empty{}, nil
}

func (lwrpc *LightWalletRpc) parse2Reply(wc *model.LightWalletCoin, c *model.CoinInfo) *proto.WalletCoin {
	return &proto.WalletCoin{
		Id:          c.Id,
		Code:        c.Code,
		Symbol:      c.Symbol,
		Name:        c.Name,
		Decimals:    int32(c.Decimals),
		CoinType:    proto.CoinType(c.Kind),
		MinConfirms: int32(c.MinConfirms),
		WalletName:  wc.WalletName,
		Path:        c.Path,
	}
}

func (lwrpc *LightWalletRpc) GetLightWallet(context context.Context, in *proto.LightWalletRequest) (*proto.LightWalletCoinsReply, error) {
	if in.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "userId required")
	}
	var coins []*model.LightWalletCoin
	var err error
	if in.WalletName == "" {
		coins, err = model.ListLWalletCoinByUser(in.UserId)
	} else {
		coins, err = model.ListLWalletCoinByUserAndWallet(in.UserId, in.WalletName)
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var reply []*proto.WalletCoin
	for _, wc := range coins {
		c, err := model.GetCoinInfo(wc.CoinId)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		reply = append(reply, lwrpc.parse2Reply(wc, c))
	}
	return &proto.LightWalletCoinsReply{Coins: reply}, nil
}
