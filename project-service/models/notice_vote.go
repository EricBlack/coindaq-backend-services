package models

import (
	"time"
	"bx.com/project-service/bxgo"
	"bx.com/project-service/utils"
	"errors"
	"github.com/go-xorm/xorm"
	"strconv"
)

type NoticeVote struct {
	Id                 int64     `xorm:"pk autoincr BIGINT(20)"`
	NoticeId           int64     `xorm:"index BIGINT(20)"`
	ApproveVote        int64     `xorm:"comment('赞成票') BIGINT(20)"`
	DisapproveVote     int64     `xorm:"comment('反对票') BIGINT(20)"`
	AbstentionVote     int64     `xorm:"comment('弃权票') BIGINT(20)"`
	PlatformVoteMax    int64     `xorm:"comment('平台票（最大值）') BIGINT(20)"`
	PlatformVoteVolumn int64     `xorm:"comment('平台投票数量') BIGINT(20)"`
	PlatformVoteType   int       `xorm:"comment('平台投票类型（赞成，反对）') INT(11)"`
	PlatformVoteReason string    `xorm:"comment('投票原因') TEXT"`
	PlatformVoteTime   time.Time `xorm:"comment('投票时间') DATETIME"`
	VoteResult         int       `xorm:"comment('投票结果') INT(11)"`
	VoteNote           string    `xorm:"TEXT"`
	CreateTime         time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateTime         time.Time `xorm:"comment('更新时间') DATETIME"`
	BeginTime          time.Time `xorm:"comment('投票开始时间') DATETIME"`
	EndTime            time.Time `xorm:"comment('投票结束时间') DATETIME"`
}


func QueryNoticeVote(noticeId int64) (*NoticeVote, error) {
	var noticeVote NoticeVote

	_, err := bxgo.OrmEngin.Where("notice_id=? ", noticeId).Get(&noticeVote)

	return &noticeVote, err
}

func QueryVoteVolumnInfo(noticeId int64) (approve, disapprove, abstention int64, err error){
	var noticeVote NoticeVote

	_, err = bxgo.OrmEngin.Where("notice_id=? ", noticeId).Get(&noticeVote)

	if err != nil {
		return 0, 0, 0, err
	}

	approve = noticeVote.ApproveVote
	disapprove = noticeVote.DisapproveVote
	abstention = noticeVote.AbstentionVote

	//添加平台投票结果
	switch noticeVote.PlatformVoteType {
	case Approve:
		approve += noticeVote.PlatformVoteVolumn
		break
	case DisApprove:
		disapprove += noticeVote.PlatformVoteVolumn
		break
	case Abstention:
		abstention += noticeVote.AbstentionVote
		break
	default:
		break
	}

	return approve, disapprove, abstention, nil
}

func UpdateNoticeInfo(session *xorm.Session, noticeInfo *NoticeVote, voteVolumn int64, voteType int) error {
	switch voteType {
	case Approve:
		noticeInfo.ApproveVote += voteVolumn
		break
	case DisApprove:
		noticeInfo.DisapproveVote += voteVolumn
		break
	case Abstention:
		noticeInfo.AbstentionVote += voteVolumn
		break
	default:
		break
	}

	noticeInfo.UpdateTime = time.Now()
	_, err := session.Where("id=? ", noticeInfo.Id).Cols("approve_vote", "disapprove_vote", "abstention_vote", "update_time").
		Update(noticeInfo)

	return err
}

func UpdateNoticeVoteTimeStatus(noticeId int64, beginTime, endTime string) error {
	begin, err := utils.String2TimeWithLocation(beginTime)
	if err != nil {
		return err
	}

	end, err := utils.String2TimeWithLocation(endTime)
	if err != nil {
		return err
	}

	if begin.After(end) {
		return errors.New("Begin time should be before After time.")
	}

	noticeVote, err := QueryNoticeVote(noticeId)
	if err != nil {
		return err
	}
	if noticeVote.Id == 0 {
		return errors.New("No such notice vote information.")
	}

	noticeVote.BeginTime = begin.Local()
	noticeVote.EndTime = end.Local()
	_,err = bxgo.OrmEngin.Id(noticeVote.Id).
		Cols("begin_time", "end_time").
		Update(noticeVote)

	return err
}

func CheckUserCanVoteNotice(userId, noticeId int64) (bool, error) {
	//查询提案信息
	noticeInfo, err := QueryNoticeNewsById(noticeId)
	if err != nil {
		return false, err
	}

	//查询投票提案信息
	noticeVoteInfo, err := QueryNoticeVote(noticeId)
	if err != nil {
		return false, err
	}
	now := time.Now().Local().Unix()
	if noticeVoteInfo.BeginTime.Unix() > now || noticeVoteInfo.EndTime.Unix() < now {		 //是否已开始
		return false, errors.New("Notice vote not started or completed.")
	}

	//判断是否已投
	voteDetails, err := QueryNoticeVoteDetailsByFilter(noticeId, userId)
	if err != nil {
		return false, err
	}
	if voteDetails.Id !=0 {
		return false, errors.New("You have voted for this notice aleardy.")
	}

	//判断用户是否可投票项目
	projectVote, err := CheckUserCanVoteProject(noticeInfo.ProjectId, userId)
	if err != nil {
		return false, err
	}
	if !projectVote {
		return false, errors.New("You have no right to vote.")
	}

	return true, nil
}

func UserVoteNotice(userId, noticeId int64, voteType int) (error) {
	//查询提案信息
	noticeInfo, err := QueryNoticeNewsById(noticeId)
	if err != nil {
		return err
	}

	//查询投票提案信息
	noticeVoteInfo, err := QueryNoticeVote(noticeId)
	if err != nil {
		return err
	}
	if noticeVoteInfo.BeginTime.After(time.Now()) || noticeVoteInfo.EndTime.Before(time.Now()) { //是否已开始
		return errors.New("Notice vote not started or completed.")
	}

	//判断是否已投
	voteDetails, err := QueryNoticeVoteDetailsByFilter(noticeId, userId)
	if err != nil {
		return err
	}
	if voteDetails.Id !=0 {
		return errors.New("You have voted for this notice aleardy.")
	}

	//判断用户是否可投票项目
	projectVote, err := CheckUserCanVoteProject(noticeInfo.ProjectId, userId)
	if err != nil {
		return err
	}
	if !projectVote {
		return errors.New("You have no right to vote.")
	}



	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	userBalance, err := QueryUserProjectBalance(userId, noticeInfo.ProjectId)
	if err != nil {
		return err
	}
	//事物处理投票
	err = session.Begin()

	//1.锁仓操作
	err = LockUserBalance(session, userId, userBalance.BalanceValue, userBalance.CurrencyId)
	if err != nil {
		session.Rollback()
		return err
	}
	//2.添加锁仓记录
	lockRecord := LockRecord{
		UserId:		userId,
		CoinId:		userBalance.CurrencyId,
		Count:		userBalance.BalanceValue,
		LockType:	VoteLock,
		LockNote:	"lock balance due to notice: " + strconv.FormatInt(noticeId, 10),
		Status:		Locked,
		LockTime:	time.Now(),
	}
	err = AddLockRecord(session, lockRecord)
	if err != nil {
		session.Rollback()
		return err
	}

	//3.添加提案投票总记录
	err = UpdateNoticeInfo(session, noticeVoteInfo, userBalance.BalanceValue, voteType)
	if err != nil {
		session.Rollback()
		return err
	}

	//4.添加用户投票记录
	err = AddUserVoteDetails(session, userId, noticeId, userBalance.BalanceValue, voteType)
	if err != nil {
		session.Rollback()
		return err
	}

	//事物提交
	err = session.Commit()

	return err
}