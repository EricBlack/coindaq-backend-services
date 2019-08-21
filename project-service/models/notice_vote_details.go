package models

import (
	"time"
	"bx.com/project-service/bxgo"
	"github.com/go-xorm/xorm"
)

type NoticeVoteDetails struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	NoticeId   int64     `xorm:"comment('提案id') index BIGINT(20)"`
	UserId     int64     `xorm:"comment('用户id') BIGINT(20)"`
	VoteType   int       `xorm:"comment('投票类型') INT(11)"`
	VoteVolumn int64     `xorm:"comment('投票量') BIGINT(20)"`
	CreateTime time.Time `xorm:"comment('投票时间') DATETIME"`
}

func QueryNoticeDetailsByNotice(noticeId int64) ([]*NoticeVoteDetails, error) {
	var noticeList []*NoticeVoteDetails

	err := bxgo.OrmEngin.Where("notice_id= ?", noticeId).Find(&noticeList)
	return noticeList, err
}

func QueryNoticeVoteDetailsByFilter(noticeId, userId int64) (*NoticeVoteDetails, error) {
	var voteDetails NoticeVoteDetails

	_, err := bxgo.OrmEngin.Where("notice_id=? ", noticeId).
		Where("user_id=? ", userId).Get(&voteDetails)

	return &voteDetails, err
}

func AddUserVoteDetails(session *xorm.Session, userId, noticeId, voteVolumn int64, voteType int) error {
	noticeDetails := NoticeVoteDetails{
		NoticeId:	noticeId,
		UserId:		userId,
		VoteType:	voteType,
		VoteVolumn:	voteVolumn,
		CreateTime: time.Now(),
	}
	_, err := session.Insert(&noticeDetails)

	return err
}
