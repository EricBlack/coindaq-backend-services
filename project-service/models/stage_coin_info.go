package models

import (
	"time"
	"bx.com/project-service/bxgo"
	"github.com/go-xorm/xorm"
)

type StageCoinInfo struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	StageId       int64     `xorm:"comment('关联阶段id') index BIGINT(20)"`
	CoinId        string    `xorm:"comment('币种id') VARCHAR(20)"`
	TargetValue   int64     `xorm:"comment('募集目标') BIGINT(20)"`
	CompleteValue int64     `xorm:"comment('完成金额') BIGINT(20)"`
	MinValue      int64     `xorm:"BIGINT(20)"`
	MaxValue      int64     `xorm:"BIGINT(20)"`
	Price         int64     `xorm:"BIGINT(20)"`
	CreateTime    time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateTime    time.Time `xorm:"comment('更新时间') DATETIME"`
}

func QueryStageCoinInfoById(id int64) (*StageCoinInfo, error) {
	var stageInfo StageCoinInfo

	_, err := bxgo.OrmEngin.Id(id).Get(&stageInfo)

	return &stageInfo, err
}

func QueryStageCoinInfoByStageId(stageId int64) ([]*StageCoinInfo, error) {
	var stageCoinList []*StageCoinInfo

	err := bxgo.OrmEngin.Where("stage_id=? ", stageId).Find(&stageCoinList)

	return stageCoinList, err
}

func QueryStageCoinByFilter(stageId int64, coinId string) (*StageCoinInfo, error) {
	var stageInfo StageCoinInfo

	_, err := bxgo.OrmEngin.Where("stage_id=? ", stageId).
		Where("coin_id=? ", coinId).Get(&stageInfo)

	return &stageInfo, err
}

func (coinInfo *StageCoinInfo) UpdateStageCoinInfo(session *xorm.Session, completeValue int64) (error) {
	coinInfo.CompleteValue += completeValue
	coinInfo.UpdateTime = time.Now()

	_, err := session.Where("id=? ", coinInfo.Id).
		Cols("complete_value", "update_time").
		Update(coinInfo)

	return err
}
