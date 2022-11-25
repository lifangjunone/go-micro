package protocol

import (
	"github.com/lifangjunone/go-micro/common/custom_logger"
	"github.com/lifangjunone/go-micro/conf"
	"github.com/lifangjunone/go-micro/service_center"
	"github.com/phachon/go-logger"
	"google.golang.org/grpc"
	"net"
)

type GRPCService struct {
	svr *grpc.Server
	l   *go_logger.Logger
	c   *conf.Config
}

func NewGRPCService() *GRPCService {
	return &GRPCService{
		svr: grpc.NewServer(),
		l:   custom_logger.CustomLog.LoggerObj,
		c:   conf.GetConfig(),
	}
}

func (s *GRPCService) Start() error {
	// 注册所有的服务到grpc服务
	service_center.RegistryAllGrpcServices(s.svr)
	// 启动HTTP服务
	lis, err := net.Listen("tcp", s.c.App.GRPC.GetAddr())
	if err != nil {
		s.l.Errorf("listen grpc tcp conn error , %s", err)
		return err
	}
	s.l.Infof("GRPC 服务监听地址:%s", s.c.App.GRPC.GetAddr())
	err = s.svr.Serve(lis)
	if err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info("service is stopped")
		}
	}
	s.l.Errorf("start grpc service error, %s", err.Error())
	return nil
}

func (s *GRPCService) Stop() error {
	s.svr.GracefulStop()
	return nil
}
