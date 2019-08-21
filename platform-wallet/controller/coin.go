package controller

import (
	"context"

	"bx.com/platform-wallet/model"
	"bx.com/platform-wallet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CoinRpc struct{}

func (crpc *CoinRpc) parse2Reply(coin *model.CoinInfo) *proto.WalletCoinReply {
	return &proto.WalletCoinReply{
		Id:          coin.Id,
		Code:        coin.Code,
		Symbol:      coin.Symbol,
		Name:        coin.Name,
		CoinType:    proto.CoinType(coin.Kind),
		Decimals:    int32(coin.Decimals),
		MinConfirms: int32(coin.MinConfirms),
	}
}

func (crpc *CoinRpc) CreatWalletCoin(context context.Context, in *proto.WalletCoinRequest) (*proto.WalletCoinReply, error) {
	coin := &model.CoinInfo{
		Code:        in.Code,
		Symbol:      in.Symbol,
		Name:        in.Name,
		Kind:        int(in.CoinType),
		Decimals:    int(in.Decimals),
		MinConfirms: int(in.MinConfirms),
	}
	if err := coin.ValidateFields(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err := model.CreateCoinInfo(coin)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return crpc.parse2Reply(coin), nil
}

func (crpc *CoinRpc) UpdateWalletCoin(context context.Context, in *proto.WalletCoinRequest) (*proto.Empty, error) {
	coin := model.CoinInfo{
		Code:        in.Code,
		Symbol:      in.Symbol,
		Name:        in.Name,
		Kind:        int(in.CoinType),
		Decimals:    int(in.Decimals),
		MinConfirms: int(in.MinConfirms),
	}
	if err := coin.ValidateFields(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := coin.Save(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.Empty{}, nil
}

func (crpc *CoinRpc) getIdFromRequest(in *proto.WalletCoinRequest) (int64, error) {
	id := in.Id
	if id == 0 {
		coin, err := model.GetCoinInfoBySymbol(in.Symbol)
		if err != nil {
			return 0, err
		}
		id = coin.Id
	}
	return id, nil
}

func (crpc *CoinRpc) EnableWalletCoinWithdraw(context context.Context, in *proto.WalletCoinRequest) (*proto.Empty, error) {
	if in.Id == 0 && in.Symbol == "" {
		return nil, status.Error(codes.InvalidArgument, "must contain id or symbol")
	}
	id, err := crpc.getIdFromRequest(in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := model.EnableWithdraw(id); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.Empty{}, nil
}

func (crpc *CoinRpc) DisableWalletCoinWithdraw(context context.Context, in *proto.WalletCoinRequest) (*proto.Empty, error) {
	if in.Id == 0 && in.Symbol == "" {
		return nil, status.Error(codes.InvalidArgument, "must contain id or symbol")
	}
	id, err := crpc.getIdFromRequest(in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := model.DisableWithdraw(id); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.Empty{}, nil
}

func (crpc *CoinRpc) EnableWalletCoinReceive(context context.Context, in *proto.WalletCoinRequest) (*proto.Empty, error) {
	if in.Id == 0 && in.Symbol == "" {
		return nil, status.Error(codes.InvalidArgument, "must contain id or symbol")
	}
	id, err := crpc.getIdFromRequest(in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := model.EnableReceive(id); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.Empty{}, nil
}

func (crpc *CoinRpc) DisableWalletCoinReceive(context context.Context, in *proto.WalletCoinRequest) (*proto.Empty, error) {
	if in.Id == 0 && in.Symbol == "" {
		return nil, status.Error(codes.InvalidArgument, "must contain id or symbol")
	}
	id, err := crpc.getIdFromRequest(in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := model.DisableReceive(id); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.Empty{}, nil
}

func (crpc *CoinRpc) GetWalletCoin(context context.Context, in *proto.WalletCoinRequest) (*proto.WalletCoinReply, error) {
	if in.Id == 0 && in.Symbol == "" {
		return nil, status.Error(codes.InvalidArgument, "must contain id or symbol")
	}
	var coin *model.CoinInfo
	var err error
	if in.Id != 0 {
		coin, err = model.GetCoinInfo(in.Id)
	} else {
		coin, err = model.GetCoinInfoBySymbol(in.Symbol)
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return crpc.parse2Reply(coin), nil
}

func (crpc *CoinRpc) ListWalletCoin(context context.Context, in *proto.Empty) (*proto.WalletCoinListReply, error) {
	var reply []*proto.WalletCoinReply
	coins, err := model.ListCoinInfo(0)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	for _, coin := range coins {
		reply = append(reply, crpc.parse2Reply(coin))
	}
	return &proto.WalletCoinListReply{Coins: reply}, nil
}
