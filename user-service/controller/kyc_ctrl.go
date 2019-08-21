package controller

import (
	"bx.com/user-service/model"
	"bx.com/user-service/proto"
	"bx.com/user-service/utils"
	"context"
	"time"
	"errors"
	"fmt"
)

type KycController struct{}

func (ctrl *KycController) CreateKyc(ctx context.Context, in *proto.CreateKycReq) (*proto.Empty, error) {
	kyc := model.KycInfo{
		UserId:    			in.UserId,
		State:     			model.KycRequestCreated,
		Kind:      			int32(in.Kind),
		RealName:  			in.RealName,
		CountryCode:		in.CountryCode,
		IdentityType:		int32(in.IdentityType),
		IdentityId:			in.IdentityId,
		PhotoFront:			in.PhotoFront,
		PhotoBack:			in.PhotoBack,
		PhotoHand:			in.PhotoHand,
		CreatedAt: 			time.Now(),
	}
	err := kyc.CreateKycInfo()

	return &proto.Empty{}, err
}
/*
func (ctrl *KycController) RejectKyc(ctx context.Context, in *proto.HandleKycReq) (*proto.Empty, error) {
	err := model.RejectKyc(in.Id, in.Reason)
	if err != nil {
		return &proto.Empty{}, err
	}
	return &proto.Empty{}, nil
}

func (ctrl *KycController) PassKyc(ctx context.Context, in *proto.HandleKycReq) (*proto.Empty, error) {
	err := model.PassKyc(in.Id, in.Reason)
	if err != nil {
		return &proto.Empty{}, err
	}
	return &proto.Empty{}, nil
}
*/
func (ctrl *KycController) UpdateKycInfo(ctx context.Context, in *proto.UpdateKycReq) (*proto.Empty, error) {
	kycInfo := model.KycInfo{
		Id:				in.Id,
		Kind:      			int32(in.Kind),
		RealName:  			in.RealName,
		CountryCode:		in.CountryCode,
		IdentityType:		int32(in.IdentityType),
		IdentityId:			in.IdentityId,
		PhotoFront:			in.PhotoFront,
		PhotoBack:			in.PhotoBack,
		PhotoHand:			in.PhotoHand,
	}

	err := kycInfo.UpdateKycInfo()

	return &proto.Empty{}, err
}

func (ctrl *KycController) QueryKycInfoById(ctx context.Context, in *proto.IdReq) (*proto.KycInfoReply, error) {
	kycF, err := model.QueryKycById(in.Id)
	if err != nil || kycF.Id == 0 {
		return &proto.KycInfoReply{}, err
	}
	//查询国家
	var country string
	countryInfo, err := model.QueryCountryById(int64(kycF.CountryCode))
	if err != nil || countryInfo.CName == ""{
		fmt.Errorf("Get country info got error: %s", err.Error())
		country = ""
	}else {
		country = countryInfo.EName
	}

	return &proto.KycInfoReply{
		Id:      kycF.Id,
		UserId:  kycF.UserId,
		State:     			proto.KycState(kycF.State),
		Kind:      			proto.UserKind(kycF.Kind),
		RealName:  			kycF.RealName,
		CountryCode:		kycF.CountryCode,
		CountryName:		country,
		IdentityType:		proto.IdentityType(kycF.IdentityType),
		IdentityId:			kycF.IdentityId,
		PhotoFront:			kycF.PhotoFront,
		PhotoBack:			kycF.PhotoBack,
		PhotoHand:			kycF.PhotoHand,
		CreateTime:			utils.Time2String(kycF.CreatedAt),
		UpdateTime:			utils.Time2String(kycF.UpdatedAt),
	}, nil
}

func (ctrl *KycController) QueryKycLastInfo(ctx context.Context, in *proto.LatestKycReq) (*proto.KycInfoReply, error) {
	if in.UserId == 0 {
		return &proto.KycInfoReply{}, errors.New("User id should not be zero")
	}

	kycF, err := model.QueryLastKycInfo(in.UserId)
	if err != nil || kycF.Id == 0 {
		return &proto.KycInfoReply{}, err
	}

	//查询国家
	var country string
	countryInfo, err := model.QueryCountryById(int64(kycF.CountryCode))
	if err != nil || countryInfo.CName == ""{
		fmt.Errorf("Get country info got error: %s", err.Error())
		country = ""
	}else {
		country = countryInfo.EName
	}

	return &proto.KycInfoReply{
		Id:      kycF.Id,
		UserId:  kycF.UserId,
		State:     			proto.KycState(kycF.State),
		Kind:      			proto.UserKind(kycF.Kind),
		RealName:  			kycF.RealName,
		CountryCode:		kycF.CountryCode,
		CountryName:		country,
		IdentityType:		proto.IdentityType(kycF.IdentityType),
		IdentityId:			kycF.IdentityId,
		PhotoFront:			kycF.PhotoFront,
		PhotoBack:			kycF.PhotoBack,
		PhotoHand:			kycF.PhotoHand,
		CreateTime:			utils.Time2String(kycF.CreatedAt),
		UpdateTime:			utils.Time2String(kycF.UpdatedAt),
	}, nil
}

func (ctrl *KycController) QueryKycInfos(ctx context.Context, in *proto.FilterReq) (*proto.KycListReply, error) {
	kycList, err := model.QueryKycByFilter(in.UserId, int32(in.State))
	if err != nil {
		return &proto.KycListReply{}, err
	}

	if len(kycList) != 0 {
		results := &proto.KycListReply{}
		for _, kycF := range kycList {
			//查询国家
			var country string
			countryInfo, err := model.QueryCountryById(int64(kycF.CountryCode))
			if err != nil || countryInfo.CName == ""{
				fmt.Errorf("Get country info got error: %s", err.Error())
				country = ""
			}else {
				country = countryInfo.EName
			}
			results.KycInfos = append(results.KycInfos, &proto.KycInfoReply{
				Id:      kycF.Id,
				UserId:  kycF.UserId,
				State:     			proto.KycState(kycF.State),
				Kind:      			proto.UserKind(kycF.Kind),
				RealName:  			kycF.RealName,
				CountryCode:		kycF.CountryCode,
				CountryName:		country,
				IdentityType:		proto.IdentityType(kycF.IdentityType),
				IdentityId:			kycF.IdentityId,
				PhotoFront:			kycF.PhotoFront,
				PhotoBack:			kycF.PhotoBack,
				PhotoHand:			kycF.PhotoHand,
				CreateTime:			utils.Time2String(kycF.CreatedAt),
				UpdateTime:			utils.Time2String(kycF.UpdatedAt),
			})
		}
		return results, nil
	}

	return &proto.KycListReply{}, nil
}

func (ctrl *KycController) QueryCountryById(ctx context.Context, in *proto.IdReq) (*proto.CountryInfoReply, error) {
	country, err := model.QueryCountryById(in.Id)
	if err != nil || country.CName == "" {
		return &proto.CountryInfoReply{}, err
	}else{
		return &proto.CountryInfoReply{
			Id:				country.Id,
			CountryNb:		country.CountryNB,
			Mark:			country.Mark,
			Ename:			country.EName,
			Cname:			country.CName,
		}, nil
	}
}

func (ctrl *KycController) QueryAllCountryInfo(ctx context.Context, in *proto.Empty) (*proto.CountryListReply, error){
	countryList, err := model.QueryAllCountry()
	if err != nil {
		return &proto.CountryListReply{}, err
	}else{
		if len(countryList) >0 {
			respList := &proto.CountryListReply{}
			for _, country := range countryList {
				respList.CountryList = append(respList.CountryList, &proto.CountryInfoReply{
					Id:				country.Id,
					CountryNb:		country.CountryNB,
					Mark:			country.Mark,
					Ename:			country.EName,
					Cname:			country.CName,
				})
			}

			return respList, nil
		}else {
			return &proto.CountryListReply{}, nil
		}
	}
}
