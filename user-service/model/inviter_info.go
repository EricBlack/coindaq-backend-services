package model

import (
	"time"
	"bx.com/user-service/bxgo"
	"math/rand"
)

type HashKey uint32

type InviterInfo struct {
	Id          int64
	UserId		int64			`xorm:"user_id bigint"`
	InviterId	int64			`xorm:"inviter_id bigint"`
	CreatedAt	time.Time		`xorm:"created_at datetime"`
}

func (ii InviterInfo) TableName() string {
	return "inviter_info"
}

func QueryInvitUserById (id int64) (*InviterInfo, error) {
	var inviteUser *InviterInfo

	_, err := bxgo.OrmEngin.Id(id).Get(&inviteUser)

	return inviteUser, err
}

func QueryInvitUsersInfo (user_id int64) ([]*InviterInfo, error) {
	var invitUsers []*InviterInfo
	err := bxgo.OrmEngin.Where("user_id= ?", user_id).Find(&invitUsers)

	return invitUsers, err
}

func CreateInviter(user_id, inviter_id int64) (int64, error) {
	inviter :=InviterInfo{
		UserId:		user_id,
		InviterId:	inviter_id,
		CreatedAt:	time.Now(),
	}

	_, err := bxgo.OrmEngin.Insert(&inviter)
	if err != nil {
		return inviter.Id, err
	}
	return inviter.Id, nil
}

func GenInviteCode() string {
	baseStr := "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	bytes := []byte(baseStr)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 16; i > 0; i-- {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}
