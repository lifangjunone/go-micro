package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/lifangjunone/go-micro/apps/category"
	"github.com/lifangjunone/go-micro/common/custom_logger"
	"github.com/lifangjunone/go-micro/service_center"
	"github.com/phachon/go-logger"
)

var (
	svr = &handler{}
)

type handler struct {
	log         *go_logger.Logger
	categorySvr category.ServiceServer
}

func (h *handler) Name() string {
	return category.AppName
}

func (h *handler) Version() string {
	return category.Version
}

func (h *handler) Config() error {
	log := custom_logger.CustomLog.LoggerObj
	h.log = log
	svr := service_center.GetGrpcService(h.Name())
	h.categorySvr = svr.(category.ServiceServer)
	return nil
}

func (h *handler) Registry(ws *restful.WebService) {
	ws.Route(ws.POST("/").To(h.CreateCategory))
	ws.Route(ws.GET("/").To(h.QueryCategory))
}

func init() {
	service_center.RegistryRestfulService(svr)
}
