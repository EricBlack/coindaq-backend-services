package controller

import (
	"context"
	"fmt"

	"bx.com/platform-wallet/model"
	"bx.com/platform-wallet/proto"
	"bx.com/platform-wallet/wallet"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PWalletRpc struct{}

func (pwrpc *PWalletRpc) parse2Reply(pc *model.PWalletCoin, c *model.CoinInfo) *proto.WalletCoin {
	return &proto.WalletCoin{
		Id:          c.Id,
		Code:        c.Code,
		Symbol:      c.Symbol,
		Name:        c.Name,
		Decimals:    int32(c.Decimals),
		CoinType:    proto.CoinType(c.Kind),
		MinConfirms: int32(c.MinConfirms),
	}
}

func (pwrpc *PWalletRpc) CreatePlatformWallet(context context.Context, in *proto.PlatformWalletRequest) (*proto.PlatformWalletReply, error) {
	if in.UserId == 0 || in.WalletType == "" {
		return nil, status.Error(codes.InvalidArgument, "userId and walletType required")
	}
	coins, err := model.ListCoinInfo(0)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var reply []*proto.WalletCoin
	for _, coin := range coins {
		walletName := fmt.Sprintf(wallet.WalletNameTemplate, in.WalletType, coin.Symbol)
		account, addr, err := wallet.CoinWallet[walletName].GetNewAccount()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		pWalletCoin := &model.PWalletCoin{
			UserId:     in.UserId,
			CoinId:     coin.Id,
			WalletName: walletName,
			Account:    account,
			Address:    addr,
		}
		if err := model.CreatePWalletCoin(pWalletCoin); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		reply = append(reply, pwrpc.parse2Reply(pWalletCoin, coin))
	}
	return &proto.PlatformWalletReply{Coins: reply}, nil
}

func (pwrpc *PWalletRpc) GetPlatformWalletCoins(context context.Context, in *proto.PlatformWalletRequest) (*proto.PlatformWalletReply, error) {
	if in.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "userId and walletType required")
	}
	coins, err := model.ListPWCByUser(in.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var reply []*proto.WalletCoin
	for _, pwc := range coins {
		coin, err := model.GetCoinInfo(pwc.Id)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		reply = append(reply, pwrpc.parse2Reply(pwc, coin))
	}
	return &proto.PlatformWalletReply{Coins: reply}, nil
}

func (pwrpc *PWalletRpc) UpdatePlatformWalletCoins(context context.Context, in *proto.UpdatePWalletRequest) (*proto.WalletCoin, error) {
	if in.UserId == 0 || in.WalletType == "" || in.CoinSymbol == "" {
		return nil, status.Error(codes.InvalidArgument, "userId, walletType, coinSymbol required")
	}
	walletName := fmt.Sprintf(wallet.WalletNameTemplate, in.WalletType, in.CoinSymbol)
	account, addr, err := wallet.CoinWallet[walletName].GetNewAccount()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	coin, err := model.GetCoinInfoBySymbol(in.CoinSymbol)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	pwc, err := model.GetPWalletCoin(in.UserId, walletName, coin.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := pwc.UpdatePWCAccountAndAddr(account, addr); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	reply := proto.WalletCoin{
		Id:          coin.Id,
		Code:        coin.Code,
		Symbol:      coin.Symbol,
		Name:        coin.Name,
		Decimals:    int32(coin.Decimals),
		CoinType:    proto.CoinType(coin.Kind),
		MinConfirms: int32(coin.MinConfirms),
	}
	return &reply, nil
}

func (pwrpc *PWalletRpc) Transfer2L(context context.Context, in *proto.TransferRequest) (*proto.TransferReply, error) {
	if in.UserId == 0 || in.PWalletType == "" || in.CoinSymbol == "" || in.LWalletAddr == "" || in.Amount == "" {
		return nil, status.Error(codes.InvalidArgument, "missing reqired field")
	}
	_, err := decimal.NewFromString(in.Amount)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	pWalletName := fmt.Sprintf(wallet.WalletNameTemplate, in.PWalletType, in.CoinSymbol)
	coin, err := model.GetCoinInfoBySymbol(in.CoinSymbol)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	pwc, err := model.GetPWalletCoin(in.UserId, pWalletName, coin.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	txId, err := wallet.CoinWallet[pWalletName].SendTransaction(pwc.Account, in.LWalletAddr, in.Amount)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.TransferReply{TxId: txId}, nil
}
