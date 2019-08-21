package models

import (
	"time"
	"errors"
	"bx.com/project-service/bxgo"
	"github.com/go-xorm/xorm"
)

type IcoOrder struct {
	Id             int64
	ProjectId      int64     `xorm:"comment('项目id') index BIGINT(20)"`
	StageNumber    int       `xorm:"comment('购买阶段') INT(11)"`
	Price          int64     `xorm:"comment('价格') BIGINT(20)"`
	Status         int       `xorm:"comment('状态，0是 到期以后发buycoin 1是已经支付了锁定。2是在1基础上，解锁完成。3是当场支付，无需解锁') INT(11)"`
	BuyCoin        string    `xorm:"comment('购买的币id') VARCHAR(20)"`
	BuyCount       int64     `xorm:"comment('购买数量') BIGINT(20)"`
	PayCoin        string    `xorm:"comment('支付的币id') VARCHAR(20)"`
	PayCount       int64     `xorm:"comment('支付数量') BIGINT(20)"`
	OrderTime      time.Time `xorm:"comment('下单时间') DATETIME"`
	UserId         int64     `xorm:"comment('用户id') BIGINT(20)"`
	IsSettlement   int       `xorm:"comment('是否已经结算0是未结算，1是已结算。') INT(11)"`
	SettlementTime time.Time `xorm:"comment('结算时间') DATETIME"`
	SettlementNote string    `xorm:"comment('结算备注') TEXT"`
}

func QueryUserIcoInfo(userId int64) ([]*IcoOrder, error) {
	var icoList []*IcoOrder

	err := bxgo.OrmEngin.Where("user_id=? ", userId).Find(&icoList)

	return icoList, err
}

func QueryUserIcoProject(userId int64) ([]int64, error) {
	var projectIds []int64
	var icoOrders []IcoOrder

	err := bxgo.OrmEngin.Where("user_id= ?", userId).
		GroupBy("project_id").Find(&icoOrders)

	if err != nil || len(icoOrders) == 0 {
		return projectIds, err
	}else {
		for _, order := range icoOrders {
			projectIds = append(projectIds, order.ProjectId)
		}

		return projectIds, err
	}
}

func QueryProjectIcoInfo(projectId int64) ([]*IcoOrder, error) {
	var icoList []*IcoOrder

	err := bxgo.OrmEngin.Where("project_id=? ", projectId).Find(&icoList)

	return icoList, err
}

func (icoOrder *IcoOrder) AddIcoOrder(session *xorm.Session) error {
	icoOrder.OrderTime = time.Now()

	_, err := session.Insert(icoOrder)
	return err
}

func JoinProjectIco (projectId, userId, price, payCount int64, stageNo int, currencyId string) error {
	//项目信息
	projectInfo, err := QueryProjectById(projectId)
	if err != nil {
		return err
	}
	if projectInfo.Status == NotStart {
		return errors.New("Project is not start.")
	}
	if projectInfo.Status == Completed {
		return errors.New("Project is completed.")
	}

	//阶段信息
	stageInfo, err := QueryProjectStageByFilter(projectId, stageNo)
	if err != nil {
		return err
	}
	if stageInfo.Id == 0 {
		return errors.New("Project stage number not correct.")
	}
	if stageInfo.StageStatus != Started {
		return errors.New("Current project stage not started.")
	}

	//募集币种信息
	coinInfo, err := QueryStageCoinByFilter(stageInfo.Id, currencyId)
	if err != nil {
		return err
	}
	if coinInfo.Id == 0 {
		return errors.New("Current stage not raise currency you specified.")
	}

	return err
}
