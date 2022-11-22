package impl

import (
	"database/sql"
	"github.com/lifangjunone/go-micro/apps/category"
	"github.com/lifangjunone/go-micro/common"
	"github.com/phachon/go-logger"
)

type impl struct {
	db  sql.DB
	log *go_logger.Logger
	category.UnimplementedServiceServer
}

func (i *impl) Name() string {
	return category.AppName
}

func (i *impl) Config() {
	log := common.NewLogger(common.LoggerConsole, i.Name())
	log.Config()
	i.log = log.LoggerObj
}
