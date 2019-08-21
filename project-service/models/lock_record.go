package models

import (
	"time"
	"github.com/go-xorm/xorm"
)

type LockRecord struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	UserId     int64     `xorm:"not null BIGINT(20)"`
	CoinId     string    `xorm:"VARCHAR(20)"`
	Count      int64     `xorm:"comment('锁仓数量') BIGINT(20)"`
	LockType   int       `xorm:"comment('锁仓类型') INT(11)"`
	LockNote   string    `xorm:"comment('锁仓原因') TEXT"`
	Status     int       `xorm:"comment('锁仓状态') INT(11)"`
	LockTime   time.Time `xorm:"comment('锁仓时间') TIMESTAMP"`
	UnlockTime time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('解锁时间') TIMESTAMP"`
}

func AddLockRecord(session *xorm.Session, record LockRecord) error {
	record.LockTime = time.Now()
	record.Status = Locked
	_, err := session.Insert(&record)

	return err
}

func UpdateUnlockRecord(session *xorm.Session, lr LockRecord) error {
	lr.UnlockTime = time.Now()
	lr.Status = Unlocked

	_, err := session.Id(lr.Id).
		Cols("status", "unlock_time").
		Update(&lr)
	return err
}