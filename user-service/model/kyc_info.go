package model

import (
	"bx.com/user-service/bxgo"
	"errors"
	"time"
)

const (
	KycRequestCreated = iota
	KycRequestPassed
	KycRequestReject
)

type KycInfo struct {
	Id         		int64
	UserId     		int64             `xorm:"user_id bigint"`
	State      		int32             `xorm:"state int"`
	Kind       		int32             `xorm:"kind int"`
	RealName   		string            `xorm:"real_name text"`
	CountryCode 	int32			  `xorm:"country_code int"`
	IdentityType   	int32 			  `xorm:"identity_type int"`
	IdentityId   	string 			  `xorm:"identity_id text"`
	PhotoFront	 	string			  `xorm:"photo_front text"`
	PhotoBack   	string 			  `xorm:"photo_back text"`
	PhotoHand   	string 			  `xorm:"photo_hand text"`
	Reason     		string            `xorm:"reason text"`
	CreatedAt  		time.Time         `xorm:"created_at datetime"`
	UpdatedAt  		time.Time         `xorm:"updated_at datetime"`
	RejectedAt 		time.Time         `xorm:"rejected_at datetime"`
	PassedAt   		time.Time         `xorm:"passed_at datetime"`
}

func (ki KycInfo) TableName() string {
	return "kyc_info"
}

func (ki *KycInfo) ValidateKycRrequestCreate() error {
	if ki.UserId == 0 {
		return errors.New("Error user id")
	}
	if ki.Kind == 0 {
		ki.Kind = UserKindPerson
	}
	if ki.RealName == ""{
		return errors.New("User real name should not be blank.")
	}
	if ki.CountryCode == 0 {
		return errors.New("Country code should not be zero.")
	}
	if ki.IdentityType == 0 {
		return errors.New("Identity type should not be zero.")
	}
	if ki.IdentityId == "" {
		return errors.New("Identity id should not be blank.")
	}
	if ki.PhotoFront =="" || ki.PhotoBack == "" || ki.PhotoHand =="" {
		return errors.New("Kyc upload photo not meet requirement.")
	}

	kycList := []KycInfo{}
	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	session.Where("user_id=? ", ki.UserId)
	session.In("state", []int32{KycRequestCreated, KycRequestPassed})

	if err := session.Find(&kycList); err != nil {
		return err
	}

	if len(kycList) > 0 {
		return errors.New("Can't create kyc request when it has existed")
	}

	return nil
}

func (ki *KycInfo) CreateKycInfo() (error) {
	err := ki.ValidateKycRrequestCreate()
	if err != nil {
		return err
	}

	_, err = bxgo.OrmEngin.Insert(ki)

	return err
}

func (ki * KycInfo) UpdateKycInfo() (error) {
	kiResult, err := QueryKycById(ki.Id)
	if err != nil {
		return err
	}
	if ki.Kind != 0 {
		kiResult.Kind = ki.Kind
	}
	if ki.RealName != ""{
		kiResult.RealName = ki.RealName
	}
	if ki.CountryCode != 0 {
		kiResult.CountryCode = ki.CountryCode
	}
	if ki.IdentityType != 0 {
		kiResult.IdentityType = ki.IdentityType
	}
	if ki.IdentityId != "" {
		kiResult.IdentityId = ki.IdentityId
	}
	if ki.PhotoFront !="" {
		kiResult.PhotoFront = ki.PhotoFront
	}
	if ki.PhotoBack != "" {
		kiResult.PhotoBack = ki.PhotoBack
	}
	if ki.PhotoHand !="" {
		kiResult.PhotoHand = ki.PhotoHand
	}
	kiResult.UpdatedAt = time.Now()

	_,err = bxgo.OrmEngin.Id(ki.Id).Update(&kiResult)

	return err
}

func QueryKycById(id int64) (KycInfo, error) {
	kyc := KycInfo{}
	_, err := bxgo.OrmEngin.Id(id).Get(&kyc)
	if err != nil {
		return kyc, err
	}
	return kyc, err
}

func QueryKycByUserId(userId int64) ([]KycInfo, error) {
	kycList := []KycInfo{}
	if err := bxgo.OrmEngin.Where("user_id=? ", userId).Find(&kycList); err != nil {
		return kycList, err
	}
	return kycList, nil
}

func QueryLastKycInfo(userId int64) (KycInfo, error) {
	kyc := KycInfo{}
	_, err := bxgo.OrmEngin.Where("user_id=? ", userId).
		Desc("id").Limit(1).Get(&kyc)

	return kyc, err
}

func QueryKycByFilter(userId int64, state int32) ([]KycInfo, error){
	kycList := []KycInfo{}
	session := bxgo.OrmEngin.NewSession()
	defer session.Close()
	if state != 0 {
		session.Where("state=? ", state)
	}
	session.Where("user_id=? ", userId)

	err := session.Find(&kycList)

	return kycList, err
}