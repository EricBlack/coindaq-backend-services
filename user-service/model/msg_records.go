package model

import (
	"time"

	"bx.com/user-service/bxgo"
)

type MsgRecords struct {
	Id            int64     `xorm:"pk autoincr BIGINT(20)"`
	Destination   string    `xorm:"TEXT"`
	Message       string    `xorm:"TEXT"`
	SendStatus    int       `xorm:"INT(11)"`
	ReturnMessage string    `xorm:"TEXT"`
	CreateTime    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

func (mi MsgRecords) TableName() string {
	return "msg_records"
}

func CreateMsgRecord(mi *MsgRecords) (err error){
	mi.CreateTime = time.Now()

	_, err = bxgo.OrmEngin.Insert(mi)
	return err
}

func MessageInHour(destination string) (int32, error){
	timeAfter := time.Now()
	timeBegin := time.Now().Add((-1)*time.Hour)
	var messages []*MsgRecords
	err := bxgo.OrmEngin.Where("destination=? ", destination).
		Where("create_time> ?", timeBegin).
		Where("create_time< ?", timeAfter).Find(&messages)
	if err != nil {
		return -1, err
	}else{
		return int32(len(messages)), err
	}
}
