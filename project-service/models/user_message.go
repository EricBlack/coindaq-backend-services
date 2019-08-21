package models

import (
	"time"
	"github.com/pkg/errors"
	"bx.com/project-service/bxgo"
)

type UserMessage struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	UserFrom   int64     `xorm:"comment('消息来源') BIGINT(20)"`
	UserTo     int64     `xorm:"comment('消息对象') BIGINT(20)"`
	Message    string    `xorm:"TEXT"`
	CreateTime time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
}

func SendMessage(fromId, toId int64, message string) (error) {
	if fromId == 0 || toId ==0 {
		return errors.New("User from or to id should not be zero.")
	}

	sendMsg := UserMessage{
		UserFrom:			fromId,
		UserTo:				toId,
		Message:			message,
		CreateTime:			time.Now(),
	}

	_, err := bxgo.OrmEngin.Insert(&sendMsg)

	return err
}

func QueryUserMessage(fromId, toId, lastId int64) ([]*UserMessage, error) {
	var messageList []*UserMessage

	if fromId == 0 || toId == 0 {
		return messageList, errors.New("User from or to id should not be zero.")
	}

	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	session.Where("user_from=? or user_from=? ", fromId, toId)
	session.Where("user_to=? or user_to=? ", fromId, toId)
	if lastId != 0 {
		session.Where("id >?", lastId)
	}

	err := session.Desc("create_time").Find(&messageList)

	return messageList, err
}