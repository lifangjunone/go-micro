package impl

import (
	"database/sql"
	"github.com/lifangjunone/go-micro/apps/category"
	"github.com/lifangjunone/go-micro/common/custom_logger"
	"github.com/lifangjunone/go-micro/conf"
	"github.com/lifangjunone/go-micro/service_center"
	"github.com/phachon/go-logger"
	"google.golang.org/grpc"
)

var (
	svr = &impl{}
)

type impl struct {
	db  *sql.DB
	log *go_logger.Logger
	category.UnimplementedServiceServer
}

func (i *impl) Name() string {
	return category.AppName
}

func (i *impl) Config() error {
	log := custom_logger.NewLogger(custom_logger.LoggerConsole, i.Name())
	log.Config()
	i.log = log.LoggerObj
	db, err := conf.GetConfig().MySQL.GetDB()
	if err != nil {
		return err
	}
	i.db = db
	return nil
}

func (i *impl) Version() string {
	return category.Version
}

func (i *impl) Registry(server *grpc.Server) {
	category.RegisterServiceServer(server, svr)
}

func init() {
	service_center.RegistryGrpcService(svr)
}
