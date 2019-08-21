package model

import (
	"math/rand"
	"time"

	"bx.com/user-service/bxgo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	_ = iota
	EmailAuthType
	PhoneAuthType
	GoogleAuthType
)

type TwoFactor struct {
	Id           int64
	UserId       int64     `xorm:"user_id bigint notnull"`
	OtpSecret    string    `xorm:"otp_secret text"`
	Activated    int32     `xorm:"activated int default 0"`
	VerifyType   int32     `xorm:"verify_type int notnull"`
	LastVerifyAt time.Time `xorm:"last_verify_at datetime"`
	RefreshedAt  time.Time `xorm:"refreshed_at datetime"`
}

func (tf TwoFactor) TableName() string {
	return "two_factors"
}

func GenSecret(length int32) string {
	baseStr := "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	bytes := []byte(baseStr)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := length; i > 0; i-- {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}

func GetLastVerifyTime(verifyType int32) (time.Time, error) {
	switch verifyType {
	case EmailAuthType: //Email Verify
		stamp, _ := time.ParseDuration("24h")
		return time.Now().Add(stamp), nil
	case PhoneAuthType: //Phone Verify
		stamp, _ := time.ParseDuration("30m")
		return time.Now().Add(stamp), nil
	case GoogleAuthType: //Google Verify
		return time.Now(), nil
	default:
		return time.Time{}, status.Error(codes.InvalidArgument, "Input verify type is not correct.")
	}
}

//Update Method
func UpdateFactorStatus(id int64, activateStatus int32) (TwoFactor, error) {
	//查询
	factor, err := GetTwoFactorById(id)
	if err != nil {
		return factor, err
	}
	if factor.Id == 0 {
		return factor, status.Error(codes.InvalidArgument, "No factor found")
	}
	//更新
	factor.Activated = activateStatus
	_, err = bxgo.OrmEngin.Id(factor.Id).Update(&factor)
	if err != nil {
		return factor, err
	}

	return factor, nil
}

func RefreshFactor(id int64, verifyType int32) (TwoFactor, error) {
	//查询
	factor, err := GetTwoFactor(id, verifyType)
	if err != nil {
		return factor, err
	}
	if factor.Id == 0 {
		return factor, status.Error(codes.InvalidArgument, "no factor found")
	}
	//更新
	var lastVerify time.Time
	if lastVerify, err = GetLastVerifyTime(factor.VerifyType); err != nil {
		return TwoFactor{}, err
	}

	factor.OtpSecret = RandSalt()
	factor.LastVerifyAt = lastVerify
	factor.Activated = False
	factor.RefreshedAt = time.Now()
	_, err = bxgo.OrmEngin.Id(factor.Id).Update(&factor)
	if err != nil {
		return factor, err
	}

	return factor, nil
}

func CreateTwoFactor(tf TwoFactor) (TwoFactor, error) {
	var lastVerify time.Time
	var err error
	var factor TwoFactor
	if lastVerify, err = GetLastVerifyTime(tf.VerifyType); err != nil {
		return TwoFactor{}, err
	}

	if factor, err = GetTwoFactor(tf.UserId, tf.VerifyType); err != nil {
		return factor, err
	}

	//Create New Factor
	if factor.Id == 0 {
		factor = TwoFactor{
			UserId:       tf.UserId,
			OtpSecret:    RandSalt(),
			VerifyType:   tf.VerifyType,
			LastVerifyAt: lastVerify,
			Activated:    False,
			RefreshedAt:  time.Now(),
		}
		if _, err = bxgo.OrmEngin.Insert(&factor); err != nil {
			return factor, err
		}
	}

	//Update
	factor.OtpSecret = RandSalt()
	factor.Activated = False
	factor.LastVerifyAt = lastVerify
	if _, err = bxgo.OrmEngin.Cols("otp_secret", "activated", "last_verify_at").Update(&factor, &TwoFactor{Id: factor.Id}); err != nil {
		return factor, err
	}

	return factor, nil
}

//Query Method
func GetTwoFactorById(id int64) (TwoFactor, error) {
	factor := TwoFactor{}
	_, err := bxgo.OrmEngin.Id(id).Get(&factor)
	if err != nil {
		return factor, err
	}
	return factor, nil
}

func GetTwoFactor(userId int64, factorType int32) (TwoFactor, error) {
	factor := TwoFactor{}
	_, err := bxgo.OrmEngin.Where("user_id= ?", userId).Where("verify_type = ?", factorType).Get(&factor)
	if err != nil {
		return factor, err
	}
	return factor, nil
}

func GetVaildTwoFactor(id int64, secret string) (TwoFactor, error) {
	factor := TwoFactor{}

	_, err := bxgo.OrmEngin.
		Where("user_id= ?", id).
		Where("otp_secret= ?", secret).
		Where("activated = ?", False).
		And("last_verify_at> ?", time.Now()).
		Get(&factor)
	if err != nil {
		return factor, err
	}
	return factor, nil
}
