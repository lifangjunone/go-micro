package protocol

import (
	"context"
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/lifangjunone/go-micro/common/custom_logger"
	"github.com/lifangjunone/go-micro/conf"
	"github.com/lifangjunone/go-micro/service_center"
	"github.com/phachon/go-logger"
	"net/http"
	"time"
)

type HTTPService struct {
	r      *restful.Container
	l      *go_logger.Logger
	c      *conf.Config
	server *http.Server
}

func NewHTTPService() *HTTPService {
	r := restful.DefaultContainer
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"*"},
		CookiesAllowed: false,
		Container:      r}
	r.Filter(cors.Filter)

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.GetConfig().App.HTTP.GetAddr(),
		Handler:           r,
	}
	return &HTTPService{
		r:      r,
		server: server,
		l:      custom_logger.CustomLog.LoggerObj,
		c:      conf.GetConfig(),
	}
}

func (s *HTTPService) PathPrefix() string {
	return fmt.Sprintf("/%s/api", s.c.App.Name)
}

// Start 启动server
func (s *HTTPService) Start() error {
	// 装在服务路由
	service_center.RegistryAllRestfulServices(s.PathPrefix(), s.r)
	s.l.Infof("HTTP服务启动成功，监听地址:%s", s.server.Addr)
	err := s.server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			s.l.Infof("service is stopped")
		}
		return err
	}
	return nil
}

// Stop 停止server
func (s *HTTPService) Stop() error {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	err := s.server.Shutdown(ctx)
	if err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit")
	}
	return nil
}
