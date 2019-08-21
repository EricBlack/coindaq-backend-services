package model

import (
	"sync"

	"bx.com/user-service/config"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var ormEngin *xorm.Engine
var setOrmEnginOnce sync.Once

func SetOrmEngin(e *xorm.Engine) {
	if ormEngin != nil {
		return
	}
	setOrmEnginOnce.Do(func() {
		ormEngin = e
	})
}

func CreateOrmEngin(cfg config.DataSource) {
	if ormEngin != nil {
		return
	}
	e, err := xorm.NewEngine(cfg.DriverName, cfg.URI)
	if err != nil {
		log.Error("data source init error", err.Error())
		return
	}
	e.ShowSQL(cfg.ShowSQL)
	e.SetMaxIdleConns(cfg.MaxIdle)
	e.SetMaxOpenConns(cfg.MaxOpen)
	SetOrmEngin(e)
	log.Info("[ok] init datasource")
}

func OrmEngin() *xorm.Engine {
	return ormEngin
}
