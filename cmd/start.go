package cmd

import (
	"errors"
	"fmt"
	_ "github.com/lifangjunone/go-micro/apps/service_registry"
	"github.com/lifangjunone/go-micro/common/custom_logger"
	"github.com/lifangjunone/go-micro/conf"
	"github.com/lifangjunone/go-micro/protocol"
	"github.com/lifangjunone/go-micro/service_center"
	"github.com/phachon/go-logger"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

type service struct {
	http *protocol.HTTPService
	grpc *protocol.GRPCService
	log  *go_logger.Logger
}

func (s *service) Start() error {
	go s.grpc.Start()
	return s.http.Start()
}

func (s *service) waitSign(sign chan os.Signal) {
	for sg := range sign {
		fmt.Println(sg)
		_ = s.http.Stop()
		_ = s.grpc.Stop()
	}
	return
}

func newService() (*service, error) {
	http := protocol.NewHTTPService()
	grpc := protocol.NewGRPCService()
	svr := &service{
		http: http,
		grpc: grpc,
		log:  custom_logger.CustomLog.LoggerObj,
	}
	return svr, nil
}

func loadServerConfig(configType string) error {
	switch configType {
	case "file":
		err := conf.LoadConfigFromToml(configFile)
		if err != nil {
			return err
		}
	case "env":
		err := conf.LoadConfigFromEnv()
		if err != nil {
			return err
		}
	default:
		return errors.New("unknown config type")
	}
	return nil
}

var runServerCmd = &cobra.Command{
	Use:   "start",
	Short: "启动服务",
	Long:  "启动服务",
	RunE: func(cmd *cobra.Command, args []string) error {

		// init global config
		if err := loadServerConfig(configType); err != nil {
			return err
		}

		// init global logger
		logger := custom_logger.NewLogger(custom_logger.LoggerConsole, conf.ServiceName)
		logger.Config()
		custom_logger.CustomLog = logger

		// init global all server[grpc and http]
		err := service_center.InitAllService()
		if err != nil {
			return err
		}

		// create chan to wait system signal
		ch := make(chan os.Signal, 1)
		defer close(ch)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

		// init server
		svr, err := newService()
		if err != nil {
			return err
		}

		// 等待退出信息
		go svr.waitSign(ch)

		// 启动服务
		if err := svr.Start(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(runServerCmd)
}
