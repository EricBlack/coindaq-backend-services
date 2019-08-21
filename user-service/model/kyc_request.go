package model

import (
	"time"
	"bx.com/user-service/bxgo"
	"errors"
)

const (
	_ = iota
	KycStageOne
	KycStageTwo
	KycStageTree
)

/*
const (
	KycRequestCreated = iota
	KycRequestPassed
	KycRequestReject
)
*/

type KycRequest struct {
	Id         int64
	UserId     int64             `xorm:"user_id bigint"`
	Stage      int32             `xorm:"stage int"`
	State      int32             `xorm:"state int"`
	Kind       int32             `xorm:"kind int"`
	Resource   map[string]string `xorm:"resource jsonb default '{}'"`
	Reason     string            `xorm:"reason text"`
	CreatedAt  time.Time         `xorm:"created_at datetime"`
	UpdatedAt  time.Time         `xorm:"updated_at datetime"`
	RejectedAt time.Time         `xorm:"rejected_at datetime"`
	PassedAt   time.Time         `xorm:"passed_at datetime"`
}

func (k KycRequest) TableName() string {
	return "kyc_requests"
}

type KycQueryOptions struct {
	State int32 `json:"state"`
	Stage int32 `json:"stage"`
}

func (kycR *KycRequest) validateKycRrequestCreate() error {
	if kycR.UserId == 0 {
		return errors.New("Error user id")
	}
	if kycR.Stage == 0 {
		return errors.New("Invalid kyc stage")
	}
	if len(kycR.Resource) == 0 {
		return errors.New("Invalid kyc resource")
	}

	kycList := []KycRequest{}
	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	session.Where("user_id=? ", kycR.UserId)
	session.Where("stage=? ", kycR.Stage)
	session.In("state", []int32{KycRequestCreated, KycRequestPassed})

	if err := session.Find(&kycList); err != nil {
		return err
	}

	if len(kycList) > 0 {
		return errors.New("Can't create kyc request when it has existed")
	}

	return nil
}

func (kycR *KycRequest) Create() (int64, error) {
	err := kycR.validateKycRrequestCreate()
	if err != nil {
		return 0, err
	}

	_, err = bxgo.OrmEngin.Insert(kycR)
	if err != nil {
		return 0, err
	}

	return kycR.Id, nil
}

func RejectKyc(id int64, reason string) error {
	_, err := bxgo.OrmEngin.Id(id).Cols("state", "reason", "updated_at", "rejected_at").
		Update(&KycRequest{State: KycRequestReject, Reason: reason, UpdatedAt: time.Now(), RejectedAt: time.Now()})
	if err != nil {
		return err
	}
	return nil
}
/*
func PassKyc(id int64, reason string) error {
	kyc, err := GetKycById(id)
	if err != nil {
		return err
	}

	session := bxgo.OrmEngin.NewSession()
	// add Begin() before any action
	defer session.Close()

	err = session.Begin()
	_, err = session.Where("id=? ", id).
		Cols("state", "updated_at", "passed_at").
		Update(&KycRequest{State: KycRequestPassed, Reason: reason, UpdatedAt: time.Now(), PassedAt: time.Now()})
	if err != nil {
		session.Rollback()
		return err
	}

	var resourceJSON []byte
	if resourceJSON, err = utils.Map2Json(kyc.Resource); err != nil {
		return err
	}
	logrus.Infof(string(resourceJSON))

	switch kyc.Stage {
	case KycStageOne:
		type stageOneResourceT struct {
			IdentityID string
			RealName   string
		}
		var stageOneResource stageOneResourceT
		json.Unmarshal(resourceJSON, &stageOneResource)

		_, err = session.Where("id=? ", kyc.UserId).
			Cols("identity_id", "real_name", "updated_at").
			Update(&User{
				IdentityId: stageOneResource.IdentityID,
				RealName:   stageOneResource.RealName,
				UpdatedAt:  time.Now()})

		if err != nil {
			session.Rollback()
			return err
		}
		break
	case KycStageTwo:
		mapInfo, err := utils.Json2Map(resourceJSON)
		if err != nil {
			return err
		}
		_, err = session.Where("id=? ", kyc.UserId).
			Cols("kyc_stage", "kyc_photos").
			Update(&User{KycStage: KycStageTwo, KycPhotos: mapInfo})

		if err != nil {
			session.Rollback()
			return err
		}
		break
	case KycStageTree:
		mapInfo, err := utils.Json2Map(resourceJSON)
		if err != nil {
			return err
		}
		_, err = session.Where("id=? ", kyc.UserId).
			Cols("kyc_stage", "kyc_vedios").
			Update(&User{KycStage: KycStageTree, KycVedios: mapInfo})

		if err != nil {
			session.Rollback()
			return err
		}
		break
	default:
		session.Rollback()
		return errors.New("Invalid stage value")
	}

	// add Commit() after all actions
	if err = session.Commit(); err != nil {
		return err
	}

	return nil
}
*/

func GetKycById(id int64) (KycRequest, error) {
	kyc := KycRequest{}
	_, err := bxgo.OrmEngin.Id(id).Get(&kyc)
	if err != nil {
		return kyc, err
	}
	return kyc, err
}

func GetKycByUserId(userId int64) ([]KycRequest, error) {
	kycList := []KycRequest{}
	if err := bxgo.OrmEngin.Where("user_id=? ", userId).Find(&kycList); err != nil {
		return kycList, err
	}
	return kycList, nil
}

func GetKycLastInfo(userId int64, kycStage int32) (KycRequest, error) {
	kyc := KycRequest{}
	_, err := bxgo.OrmEngin.Where("user_id=? ", userId).
		Where("stage=? ", kycStage).
		Desc("id").Get(&kyc)

	return kyc, err
}

func GetKycByFilter(userId int64, queryOption KycQueryOptions) ([]KycRequest, error) {
	kycList := []KycRequest{}
	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	session.Where("user_id=? ", userId)
	if queryOption.State != 0 {
		session.Where("state=? ", queryOption.State)
	}
	if queryOption.Stage != 0 {
		session.Where("stage=? ", queryOption.Stage)
	}
	if err := session.Find(&kycList); err != nil {
		return kycList, err
	}

	return kycList, nil
}
