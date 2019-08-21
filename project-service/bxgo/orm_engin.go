package bxgo

import (
	"sync"

	"bx.com/project-service/config"
	"github.com/go-xorm/xorm"
	_"github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var OrmEngin *xorm.Engine
var setOrmEnginOnce sync.Once

func SetOrmEngin(e *xorm.Engine) {
	if OrmEngin != nil {
		return
	}
	setOrmEnginOnce.Do(func() {
		OrmEngin = e
	})
}

func CreateOrmEngin(cfg config.DataSource) {
	if OrmEngin != nil {
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
