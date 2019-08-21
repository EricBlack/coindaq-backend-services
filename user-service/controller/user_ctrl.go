package controller

import (
	"context"

	"strings"
	"time"

	"bx.com/user-service/model"
	"bx.com/user-service/proto"
	"bx.com/user-service/utils"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"errors"
)

type UserController struct{}

func (ctrl *UserController) Signup(ctx context.Context, in *proto.RegisterReq) (*proto.UserReply, error) {
	//If Enable Invite Code
	var userId int64
	var err error
	if in.EnableCode {
		if userId, err = model.QueryInvitCode(in.InviteCode); err != nil {
			return &proto.UserReply{}, status.Error(codes.Internal, "Interal query error.")
		}
		if userId == 0 {
			return &proto.UserReply{}, status.Error(codes.Internal, "Invite code is invaild.")
		}
	}

	//Check
	user, err := model.GetUserByEmail(in.Email)
	if err != nil {
		return &proto.UserReply{}, err
	}

	if user.Id != 0 {
		return &proto.UserReply{}, status.Error(codes.InvalidArgument, "Provided email have been registed.")
	}

	//Create
	salt := model.RandSalt()
	password := model.GenPwd(in.Password, salt)
	if err != nil {
		return &proto.UserReply{}, status.Error(codes.Internal, "Password generated error.")
	}

	user = model.User{
		Email:       in.Email,
		Password:    password,
		Salt:        salt,
		CountryCode: in.CountryCode,
		DisplayName: in.DisplayName,
		InviteCode:  in.InviteCode,
		RegisterIp:  in.RegisterIp,
		DeviceId:    in.DeviceId,
	}
	id, err := model.CreateUser(user)
	if err != nil {
		return &proto.UserReply{}, status.Error(codes.Internal, "Create user error.")
	}

	if id == 0 {
		return &proto.UserReply{}, status.Error(codes.Internal, "Create user error.")
	}

	//Add invite info
	if userId != 0 {
		if _, err := model.CreateInviter(userId, id); err != nil {
			return &proto.UserReply{}, status.Error(codes.Internal, "Create invite info error.")
		}
	}

	//Add verification Items
	factor := model.TwoFactor{
		UserId:     id,
		VerifyType: model.EmailAuthType,
	}
	if factor, err = model.CreateTwoFactor(factor); err != nil {
		return &proto.UserReply{}, err
	}

	factor.VerifyType = model.PhoneAuthType
	if factor, err = model.CreateTwoFactor(factor); err != nil {
		return &proto.UserReply{}, err
	}

	factor.VerifyType = model.GoogleAuthType
	if factor, err = model.CreateTwoFactor(factor); err != nil {
		return &proto.UserReply{}, err
	}

	userReply := proto.UserReply{
		Id:          id,
		Email:       user.Email,
		RegistIp:    user.RegisterIp,
		DeviceId:    user.DeviceId,
		DisplayName: user.DisplayName,
	}
	return &userReply, nil
}

func (ctrl *UserController) Signin(ctx context.Context, in *proto.AuthReq) (*proto.UserReply, error) {
	user, err := model.GetUserByEmail(in.Email)
	if err != nil || user.Id == 0 {
		return &proto.UserReply{}, status.Error(codes.Internal, "Find user error.")
	}

	pwd := model.GenPwd(in.Password, user.Salt)
	loginInfo := model.LoginRecord{
		UserId:    user.Id,
		LoginIp:   in.LoginIp,
		DeviceId:  in.DeviceId,
		LoginTime: time.Now(),
	}
	if pwd != user.Password {
		//记录登陆失败信息
		loginInfo.LoginStatus = model.LoginFailed
		loginInfo.LoginComment = "password not correct"
		_ = model.CreateLoginRecord(&loginInfo)

		return &proto.UserReply{}, status.Error(codes.PermissionDenied, "Password verify error.")
	}
	if user.Activated == model.False {
		//记录登陆失败信息
		loginInfo.LoginStatus = model.LoginFailed
		loginInfo.LoginComment = "user not activated"
		_ = model.CreateLoginRecord(&loginInfo)

		return &proto.UserReply{}, status.Error(codes.PermissionDenied, "User not activated.")
	}

	if user.Disabled == model.True {
		//记录登陆失败信息
		loginInfo.LoginStatus = model.LoginFailed
		loginInfo.LoginComment = "account disabled"
		_ = model.CreateLoginRecord(&loginInfo)

		return &proto.UserReply{}, status.Error(codes.PermissionDenied, "Account was been disabled.")
	}

	//生成token信息
	token, err := uuid.NewV4()
	if err != nil {
		return &proto.UserReply{}, err
	}
	err = model.UpdateUserToken(user.Id, token.String())
	if err != nil {
		return &proto.UserReply{}, err
	}

	//记录登陆成功信息
	loginInfo.LoginStatus = model.LoginPassed
	loginInfo.LoginComment = "success"
	_ = model.CreateLoginRecord(&loginInfo)

	//生成钱包地址
	model.CreateAllUserCoinAddress(user.Id)

	userReply := proto.UserReply{
		Id:          	user.Id,
		Email:       	user.Email,
		DisplayName: 	user.DisplayName,
		PhoneNumber:	user.PhoneNumber,
		Kind:			proto.UserKind(user.Kind),
		Activated:		proto.BoolValue(user.Activated),
		CountryCode:	user.CountryCode,
		RealName:		user.RealName,
		Disabled:		proto.BoolValue(user.Disabled),
		IdentityType:	proto.IdentityType(user.IdentityType),
		IdentityId:		user.IdentityId,
		PhotoFront:		user.PhotoFront,
		PhotoBack:		user.PhotoBack,
		PhotoHand:		user.PhotoBack,
		InviteCode:		user.InviteCode,
		RegistIp:		user.RegisterIp,
		DeviceId:		user.DeviceId,
		GoogleIsBind: 	user.GoogleIsBind,
	}

	return &userReply, nil
}

func (ctrl *UserController) CheckUserStatus(ctx context.Context, in *proto.ForgetEmailReq) (*proto.UserStatusReply, error) {
	userInfo, err := model.GetUserByEmail(in.Email)
	if err != nil {
		return &proto.UserStatusReply{}, err
	}
	if userInfo.Id == 0 {
		return &proto.UserStatusReply{}, errors.New("Email account not registed.")
	} else {
		//手机绑定检查
		userStatus := proto.UserStatusReply{}
		if userInfo.PhoneNumber == "" {
			userStatus.IsPhoneBind = proto.BoolValue_False
		}else{
			userStatus.IsPhoneBind = proto.BoolValue_True
		}
		//Google绑定检查
		if userInfo.GoogleIsBind == model.True {
			userStatus.IsGoogleBind = proto.BoolValue_True
		}else {
			userStatus.IsGoogleBind = proto.BoolValue_False
		}
		//用户激活检查
		if userInfo.Activated == model.True {
			userStatus.IsActivated = proto.BoolValue_True
		}else {
			userStatus.IsActivated = proto.BoolValue_False
		}
		//用户被禁检查
		if userInfo.Disabled == model.True {
			userStatus.IsDisabled = proto.BoolValue_True
		}else {
			userStatus.IsDisabled = proto.BoolValue_False
		}

		//是否支付密码设置
		if userInfo.PaymentPassword != "" {
			userStatus.IsPayPassword = proto.BoolValue_True
		} else {
			userStatus.IsDisabled = proto.BoolValue_False
		}

		//检查Kyc状态
		if userInfo.RealName != "" && userInfo.IdentityId != "" {
			userStatus.IsKycPassed = proto.BoolValue_True
		} else {
			userStatus.IsKycPassed = proto.BoolValue_False
		}

		return &userStatus, nil
	}
}

func (ctrl *UserController) VerifyUserToken(ctx context.Context, in *proto.UserReq) (*proto.UserReply, error) {
	user, err := model.QueryUserByToken(in.Token)
	if err != nil {
		return &proto.UserReply{}, err
	}

	return &proto.UserReply{
		Id:          user.Id,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		PhoneNumber: user.PhoneNumber,
		Kind:        proto.UserKind(user.Kind),
		Activated:   proto.BoolValue(user.Activated),
		CountryCode: user.CountryCode,
		RealName:    user.RealName,
		Disabled:    proto.BoolValue(user.Disabled),
		IdentityType:proto.IdentityType(user.IdentityType),
		IdentityId:  user.IdentityId,
		PhotoFront:  user.PhotoFront,
		PhotoBack:   user.PhotoBack,
		PhotoHand:   user.PhotoHand,
		InviteCode:  user.InviteCode,
		RegistIp:    user.RegisterIp,
		DeviceId:    user.DeviceId,
	}, nil
}

func (ctrl *UserController) QueryLoginRecords(ctx context.Context, in *proto.RecordFilterReq) (*proto.LoginRecordListReply, error) {
	filter := model.LoginFilter{
		UserId:      in.UserId,
		LoginStatus: int32(in.Status),
		LoginTime:   in.LoginTime,
	}

	recordList, err := model.QueryLoginRecordsByFilter(&filter)
	if err != nil {
		return &proto.LoginRecordListReply{}, err
	}
	var resp *proto.LoginRecordListReply
	for _, record := range recordList {
		resp.RecordList = append(resp.RecordList, &proto.LoginRecordReply{
			Id:           record.Id,
			UserId:       record.UserId,
			LoginIp:      record.LoginIp,
			DeviceId:     record.DeviceId,
			Status:       proto.LoginStatus(record.LoginStatus),
			LoginComment: record.LoginComment,
			LoginTime:    utils.Time2String(record.LoginTime),
		})
	}

	return resp, nil
}

func (ctrl *UserController) ActivateNewUser(ctx context.Context, in *proto.ForgetEmailReq) (*proto.FactorReply, error) {
	userInfo, err := model.GetUserByEmail(in.Email)
	if err != nil {
		return &proto.FactorReply{}, err
	}
	if userInfo.Id == 0 {
		return &proto.FactorReply{}, errors.New("Email is not registed.")
	}
	if userInfo.Activated == model.True {
		return &proto.FactorReply{}, errors.New("User have been activated aleardy.")
	}

	resultF := model.TwoFactor{
		UserId:     userInfo.Id,
		VerifyType: model.EmailAuthType,
	}
	result ,err := model.CreateTwoFactor(resultF)
	if err != nil {
		return &proto.FactorReply{}, err
	} else {
		return &proto.FactorReply{
			Id:				result.Id,
			UserId:			result.UserId,
			Code:			result.OtpSecret,
			Activated:		proto.BoolValue(result.Activated),
			LastVerifyAt:	utils.Time2String(result.LastVerifyAt),
		}, nil
	}
}

func (ctrl *UserController) ActivateEmailUser(ctx context.Context, in *proto.ActivateReq) (*proto.Empty, error) {
	affted, err := model.ActivateUser(in.UserId, in.Secret)
	if err != nil {
		return nil, err
	}
	if affted == 1 {
		return &proto.Empty{}, nil
	}
	return nil, status.Error(codes.Internal, "user activated failed.")
}

func (ctrl *UserController) ForgetPasswordViaEmail(ctx context.Context, in *proto.ForgetEmailReq) (*proto.ForgetReply, error) {
	user, err := model.GetUserByEmail(in.Email)
	if err != nil {
		return &proto.ForgetReply{}, status.Error(codes.Internal, "Find user error.")
	}
	if user.Id == 0 {
		return &proto.ForgetReply{}, status.Error(codes.Internal, "Find user error.")
	}
	factor := model.TwoFactor{
		UserId:     user.Id,
		VerifyType: int32(proto.FactorType_EmailAuthType),
	}

	result, err := model.CreateTwoFactor(factor)
	if err != nil {
		return &proto.ForgetReply{}, err
	}

	if result.Id == 0 {
		return &proto.ForgetReply{}, status.Error(codes.Internal, "No factor created.")
	}

	return &proto.ForgetReply{
		Id:          user.Id,
		Information: user.Email,
		Secrect:     result.OtpSecret,
	}, nil
}

func (ctrl *UserController) ForgetPasswordViaPhone(ctx context.Context, in *proto.ForgetPhoneReq) (*proto.ForgetReply, error) {
	userInfo, err := model.GetUserByEmail(in.Email)
	if err != nil {
		return &proto.ForgetReply{}, errors.New("Find user error, wrong email.")
	}
	if userInfo.Id == 0 {
		return &proto.ForgetReply{}, errors.New("Cannot find related user.")
	}
	if userInfo.PhoneNumber != in.Phone {
		return &proto.ForgetReply{}, errors.New("Your phone number is not correct.")
	}

	factor := model.TwoFactor{
		UserId:     userInfo.Id,
		VerifyType: int32(proto.FactorType_PhoneAuthType),
	}

	result, err := model.CreateTwoFactor(factor)
	if err != nil {
		return &proto.ForgetReply{}, err
	}

	if result.Id == 0 {
		return &proto.ForgetReply{}, status.Error(codes.Internal, "No factor created.")
	}

	return &proto.ForgetReply{
		Id:          userInfo.Id,
		Information: userInfo.PhoneNumber,
		Secrect:     result.OtpSecret,
	}, nil
}

func (ctrl *UserController) RecordMessageInfo(ctx context.Context, in *proto.RecordMessageReq) (*proto.Empty, error) {
	msg := model.MsgRecords{
		Destination:   in.Destination,
		Message:       in.Message,
		SendStatus:    int(in.SendStatus),
		ReturnMessage: in.ReturnMessage,
		CreateTime:    time.Now(),
	}
	err := model.CreateMsgRecord(&msg)

	return &proto.Empty{}, err
}

func (ctrl *UserController) MessageInHour(ctx context.Context, in *proto.ForgetPhoneReq) (*proto.MessageCountReply, error) {
	count, err := model.MessageInHour(in.Phone)
	if err != nil {
		return &proto.MessageCountReply{}, err
	}
	return &proto.MessageCountReply{
		Count: count,
	}, nil
}

func (ctrl *UserController) ResetPassword(ctx context.Context, in *proto.ModifyPasswordReq) (*proto.Empty, error) {
	affted, err := model.ResetPassword(in.Password, in.Code)
	if err != nil {
		return &proto.Empty{}, err
	}
	if affted == 1 {
		return &proto.Empty{}, nil
	}
	return &proto.Empty{}, status.Error(codes.Internal, "No password item reseted.")
}

func (ctrl *UserController) UpdatePassword(ctx context.Context, in *proto.UpdatePasswordReq) (*proto.Empty, error) {
	affted, err := model.UpdatePassword(in.Id, in.OldPassword, in.NewPassword)
	if err != nil {
		return &proto.Empty{}, err
	}

	if affted == 1 {
		return &proto.Empty{}, nil
	}

	return &proto.Empty{}, status.Error(codes.Internal, "No password item updated.")
}

func (ctrl *UserController) Logout(ctx context.Context, in *proto.IdReq) (*proto.Empty, error) {
	err := model.UpdateUserToken(in.Id, "")

	return &proto.Empty{}, err
}

func (ctrl *UserController) UpdateUserInfo(ctx context.Context, in *proto.ModifyUserReq) (*proto.Empty, error) {
	queryUser, err := model.GetUserByID(in.Id)
	if err != nil {
		return &proto.Empty{}, err
	}
	if queryUser.Id == 0 {
		return &proto.Empty{}, status.Error(codes.Internal, "No such user.")
	}

	//参数判断
	if in.DisplayName != "" {
		queryUser.DisplayName = in.DisplayName
	}
	if in.CountryCode != "" {
		queryUser.CountryCode = in.CountryCode
	}
	if in.RealName != "" {
		queryUser.RealName = in.RealName
	}
	if in.IdentityId != "" {
		queryUser.IdentityId = in.IdentityId
	}
	queryUser.UpdatedAt = time.Now()

	userId, err := model.UpdateUserInfo(queryUser)
	if err != nil {
		return &proto.Empty{}, err
	}
	if userId == 0 {
		return &proto.Empty{}, status.Error(codes.Internal, "User information update failed.")
	}

	return &proto.Empty{}, nil
}

func (ctrl *UserController) SendBindMessage(ctx context.Context, in *proto.SendMessageReq) (*proto.FactorReply, error) {
	if in.UserId == 0 {
		return &proto.FactorReply{}, errors.New("User id parameter should be provided.")
	}
	userInfo, err := model.GetUserByPhone(in.Phone)
	if err != nil {
		return &proto.FactorReply{}, err
	}
	if userInfo.Id != 0 {
		return &proto.FactorReply{}, errors.New("Phone number have been binded.")
	}

	factor := model.TwoFactor{UserId:in.UserId, VerifyType:model.PhoneAuthType}
	factorResult, err := model.CreateTwoFactor(factor)
	if err != nil {
		return &proto.FactorReply{}, err
	} else {
		return &proto.FactorReply{
			Id:				factorResult.Id,
			UserId:			factorResult.UserId,
			Code:			factorResult.OtpSecret,
			Activated:		proto.BoolValue(factorResult.Activated),
			LastVerifyAt:	utils.Time2String(factorResult.LastVerifyAt),
			RefreshedAt:	utils.Time2String(factorResult.RefreshedAt),
		}, nil
	}
}

func (ctrl *UserController) BindUserPhone(ctx context.Context, in *proto.BindPhoneReq) (*proto.Empty, error) {
	affted, err := model.BindPhone(in.Id, in.Phone, strings.ToUpper(in.Code))
	if err != nil {
		return &proto.Empty{}, err
	}
	if affted == 1 {
		return &proto.Empty{}, nil
	}
	return &proto.Empty{}, status.Error(codes.Internal, "No phone item binded.")
}

func (ctrl *UserController) QueryUserById(ctx context.Context, in *proto.IdReq) (*proto.UserReply, error) {
	user, err := model.QueryUserById(in.Id)
	if err != nil || user.Email == "" {
		return &proto.UserReply{}, err
	}

	return &proto.UserReply{
		Id:          user.Id,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		PhoneNumber: user.PhoneNumber,
		Kind:        proto.UserKind(user.Kind),
		Activated:   proto.BoolValue(user.Activated),
		CountryCode: user.CountryCode,
		RealName:    user.RealName,
		Disabled:    proto.BoolValue(user.Disabled),
		IdentityType:proto.IdentityType(user.IdentityType),
		IdentityId:  user.IdentityId,
		PhotoFront:  user.PhotoFront,
		PhotoBack:   user.PhotoBack,
		PhotoHand:   user.PhotoHand,
		InviteCode:  user.InviteCode,
		RegistIp:    user.RegisterIp,
		DeviceId:    user.DeviceId,
	}, nil
}

func (ctrl *UserController) QueryUsers(ctx context.Context, in *proto.QueryUserReq) (*proto.UserListReply, error) {
	list, err := model.QueryUsersByFilter(model.UserFilter{
		Email:       in.Email,
		DisplayName: in.DisplayName,
		PhoneNumber: in.PhoneNumber,
		RealName:    in.RealName,
		Kind:        int32(in.Kind),
		Activated:   int32(in.Activated),
		Disabled:    int32(in.Disabled),
	})
	if err != nil {
		return &proto.UserListReply{}, err
	}
	if len(list) != 0 {
		results := &proto.UserListReply{}
		for i := 0; i < len(list); i++ {
			results.Users = append(results.Users, &proto.UserReply{
				Id:          list[i].Id,
				Email:       list[i].Email,
				DisplayName: list[i].DisplayName,
				PhoneNumber: list[i].PhoneNumber,
				Kind:        proto.UserKind(list[i].Kind),
				Activated:   proto.BoolValue(list[i].Activated),
				CountryCode: list[i].CountryCode,
				RealName:    list[i].RealName,
				Disabled:    proto.BoolValue(list[i].Disabled),
				IdentityType:proto.IdentityType(list[i].IdentityType),
				IdentityId:  list[i].IdentityId,
				PhotoFront:  list[i].PhotoFront,
				PhotoBack:   list[i].PhotoBack,
				PhotoHand:   list[i].PhotoHand,
				InviteCode:  list[i].InviteCode,
				RegistIp:    list[i].RegisterIp,
				DeviceId:    list[i].DeviceId,
			})
		}
		return results, nil
	}

	return &proto.UserListReply{}, nil
}

func (ctrl *UserController) QueryInvitedUsersByInvitorId(ctx context.Context, in *proto.IdReq) (*proto.UserListReply, error) {
	list, err := model.QueryInvitedUser(in.Id)
	if err != nil {
		return &proto.UserListReply{}, err
	}
	if len(list) != 0 {
		results := &proto.UserListReply{}
		for i := 0; i < len(list); i++ {
			results.Users = append(results.Users, &proto.UserReply{
				Id:          list[i].Id,
				Email:       list[i].Email,
				DisplayName: list[i].DisplayName,
				PhoneNumber: list[i].PhoneNumber,
				Kind:        proto.UserKind(list[i].Kind),
				Activated:   proto.BoolValue(list[i].Activated),
				CountryCode: list[i].CountryCode,
				RealName:    list[i].RealName,
				Disabled:    proto.BoolValue(list[i].Disabled),
				IdentityType:proto.IdentityType(list[i].IdentityType),
				IdentityId:  list[i].IdentityId,
				PhotoFront:  list[i].PhotoFront,
				PhotoBack:   list[i].PhotoBack,
				PhotoHand:   list[i].PhotoHand,
				InviteCode:  list[i].InviteCode,
				RegistIp:    list[i].RegisterIp,
				DeviceId:    list[i].DeviceId,
			})
		}
		return results, nil
	}

	return &proto.UserListReply{}, nil
}

func (ctrl *UserController) QueryInvitorInfoById(ctx context.Context, in *proto.IdReq) (*proto.InviterInfoReply, error) {
	inviter, err := model.QueryInvitUserById(in.Id)
	if err != nil || inviter.UserId == 0 {
		return &proto.InviterInfoReply{}, err
	}
	return &proto.InviterInfoReply{
		Id:        inviter.Id,
		UserId:    inviter.UserId,
		InviterId: inviter.InviterId,
		CreateAt:  utils.Time2String(inviter.CreatedAt),
	}, nil
}

func (ctrl *UserController) UpdatePaymentPassword(ctx context.Context, in *proto.PaymentPasswordReq) (*proto.Empty, error) {
	err := model.UpdatePaymentPassword(in.UserId, in.Password)
	return &proto.Empty{}, err
}

func (ctrl *UserController) VerifyPaymentPassword(ctx context.Context, in *proto.PaymentPasswordReq) (*proto.BoolReply, error) {
	result, err := model.VerifyPaymentPassword(in.UserId, in.Password)
	if err != nil {
		return &proto.BoolReply{
			Result: proto.BoolValue_False,
		}, err
	} else {
		if result {
			return &proto.BoolReply{
				Result: proto.BoolValue_True,
			}, nil
		} else {
			return &proto.BoolReply{
				Result: proto.BoolValue_False,
			}, nil
		}
	}
}

func (ctrl *UserController) VerifyFinancialOperation(ctx context.Context, in *proto.IdReq) (*proto.BoolReply, error) {
	result, err := model.VerifyFinancialOperation(in.Id)
	if err != nil {
		return &proto.BoolReply{
			Result: proto.BoolValue_False,
		}, err
	} else {
		if result {
			return &proto.BoolReply{
				Result: proto.BoolValue_True,
			}, nil
		} else {
			return &proto.BoolReply{
				Result: proto.BoolValue_False,
			}, nil
		}
	}
}

func (ctrl *UserController) BindUserGoogleFactor(ctx context.Context, in *proto.IdReq) (*proto.GoogleQRCodeRelpy, error) {
	fileName, code, err := model.BindUserGoogleFactor(in.Id)

	if err != nil {
		return &proto.GoogleQRCodeRelpy{}, err
	} else {
		return &proto.GoogleQRCodeRelpy{
			ImageUrl: fileName,
			Code: 	  code,
		}, nil
	}
}

func (ctrl *UserController) RemoveBindGoogleFactor(ctx context.Context, in *proto.IdReq) (*proto.Empty, error) {
	err := model.DeleteBindGoogleFactor(in.Id)

	return &proto.Empty{}, err
}
