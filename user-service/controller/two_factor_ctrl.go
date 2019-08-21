package controller

import (
	"context"

	"bx.com/user-service/model"
	"bx.com/user-service/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/satori/go.uuid"
)

type TwoFactorController struct{}

func (ctrl *TwoFactorController) GenerateFactor(ctx context.Context, in *proto.FactorReq) (*proto.FactorReply, error) {
	resultF := model.TwoFactor{
		UserId:     in.UserId,
		VerifyType: int32(in.Type),
	}
	var err error
	if resultF, err = model.GetTwoFactor(resultF.UserId, resultF.VerifyType); err != nil {
		return &proto.FactorReply{}, status.Errorf(codes.InvalidArgument, "Parameter error, user id and factor type needed.")
	}

	if in.Type == model.GoogleAuthType {
		if err = resultF.GenOTPSecret(false); err != nil {
			return &proto.FactorReply{}, err
		}
	} else {
		resultF, err = model.CreateTwoFactor(resultF)
		if err != nil {
			return &proto.FactorReply{}, err
		}
	}

	return &proto.FactorReply{
		Id:           resultF.Id,
		UserId:       resultF.UserId,
		Code:         resultF.OtpSecret,
		Activated:    proto.BoolValue(resultF.Activated),
		LastVerifyAt: resultF.LastVerifyAt.Format("2006-01-02 15:04:05"),
		RefreshedAt:  resultF.RefreshedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (ctrl *TwoFactorController) QueryFactor(ctx context.Context, in *proto.FactorReq) (*proto.FactorReply, error) {
	resultF, err := model.GetTwoFactor(in.UserId, int32(in.Type))
	if err != nil || resultF.Id == 0 {
		return &proto.FactorReply{}, err
	}

	return &proto.FactorReply{
		Id:           resultF.Id,
		UserId:       resultF.UserId,
		Code:         resultF.OtpSecret,
		Activated:    proto.BoolValue(resultF.Activated),
		LastVerifyAt: resultF.LastVerifyAt.Format("2006-01-02 15:04:05"),
		RefreshedAt:  resultF.RefreshedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (ctrl *TwoFactorController) RefreshFactor(ctx context.Context, in *proto.FactorReq) (*proto.FactorReply, error) {
	var resultF model.TwoFactor
	var err error

	if in.Type == model.GoogleAuthType {
		resultF = model.TwoFactor{
			UserId:     in.UserId,
			VerifyType: int32(in.Type),
		}
		if err = resultF.Refresh(); err != nil {
			return &proto.FactorReply{}, err
		}
	} else {
		resultF, err = model.RefreshFactor(in.UserId, int32(in.Type))
		if err != nil {
			return &proto.FactorReply{}, err
		}
	}

	return &proto.FactorReply{
		Id:           resultF.Id,
		UserId:       resultF.UserId,
		Code:         resultF.OtpSecret,
		Activated:    proto.BoolValue(resultF.Activated),
		LastVerifyAt: resultF.LastVerifyAt.Format("2006-01-02 15:04:05"),
		RefreshedAt:  resultF.RefreshedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (ctrl *TwoFactorController) VerifyFactorCode(ctx context.Context, in *proto.InfoReq) (*proto.VerifyReply, error) {
	if in.Type == proto.FactorType_EmailAuthType || in.Type == proto.FactorType_PhoneAuthType {
		resultF, err := model.GetVaildTwoFactor(in.UserId, in.Code)
		if err != nil {
			return &proto.VerifyReply{}, err
		}

		if resultF.Id == 0 {
			return &proto.VerifyReply{}, status.Error(codes.InvalidArgument, "Factor not correct or timeout.")
		}

		//更新
		resultF, err = model.UpdateFactorStatus(resultF.Id, model.True)
		if err != nil {
			return &proto.VerifyReply{}, err
		}

		return &proto.VerifyReply{
			Result: true,
		}, nil
	} else if in.Type == proto.FactorType_GoogleAuthType {
		resultF, err := model.GetTwoFactor(in.UserId, model.GoogleAuthType)
		if err != nil {
			return nil, err
		}
		if resultF.Id == 0 {
			return nil, status.Error(codes.InvalidArgument, "Factor not found.")
		}

		result := resultF.VerifyOTP(in.Code)

		if result {
			return &proto.VerifyReply{
				Result: result,
			}, nil
		} else {
			return &proto.VerifyReply{}, errors.New("Google factor verification failed.")
		}
	} else {
		return &proto.VerifyReply{}, errors.New("Input factor type error.")
	}
}

func (ctrl *TwoFactorController) VerifyUserFactors(ctx context.Context, in *proto.FactorListReq) (*proto.VerifyFactorsReply, error) {
	if in.Email == "" {
		return &proto.VerifyFactorsReply{}, errors.New("Email should not be blank.")
	}
	userInfo, err := model.GetUserByEmail(in.Email)
	if err != nil {
		return &proto.VerifyFactorsReply{}, err
	}
	if userInfo.Id == 0 {
		return &proto.VerifyFactorsReply{}, errors.New("No such user.")
	}

	//验证Email Code
	if in.EmailCode == "" {
		return &proto.VerifyFactorsReply{}, errors.New("Email code and phone code should not blank.")
	}
	emailF, err := model.GetVaildTwoFactor(userInfo.Id, in.EmailCode)
	if err != nil {
		return &proto.VerifyFactorsReply{}, err
	}
	if emailF.Id == 0 {
		return &proto.VerifyFactorsReply{}, errors.New("Email code not correct or timeout.")
	}

	//验证Phone Code
	if userInfo.PhoneNumber != "" {
		if in.PhoneCode == "" {
			return &proto.VerifyFactorsReply{}, errors.New("Phone code should not blank.")
		}
		phoneF, err := model.GetVaildTwoFactor(userInfo.Id, in.PhoneCode)
		if err != nil {
			return &proto.VerifyFactorsReply{}, err
		}

		if phoneF.Id == 0 {
			return &proto.VerifyFactorsReply{}, errors.New("Phone code not correct or timeout.")
		}
	}

	//验证 Google Code
	if err != nil {
		return &proto.VerifyFactorsReply{}, err
	}

	if userInfo.GoogleIsBind == model.True {
		if in.GoogleCode == "" {
			return &proto.VerifyFactorsReply{}, errors.New("Google code should not blank.")
		}
		googleF, err := model.GetTwoFactor(userInfo.Id, model.GoogleAuthType)
		if err != nil {
			return nil, err
		}
		result := googleF.VerifyOTP(in.GoogleCode)
		if !result {
			return &proto.VerifyFactorsReply{}, errors.New("Google code not correct.")
		}
	}

	//生成验证凭证
	token, err := uuid.NewV4()
	if err != nil {
		return &proto.VerifyFactorsReply{}, err
	}
	err = model.UpdateUserToken(userInfo.Id, token.String())
	if err != nil {
		return &proto.VerifyFactorsReply{}, err
	}

	return &proto.VerifyFactorsReply{ Code: token.String()}, err
}