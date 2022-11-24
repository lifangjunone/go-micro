package service_center

import "google.golang.org/grpc"

var (
	grpcServices = map[string]GrpcService{}
)

type GrpcService interface {
	Registry(*grpc.Server)
	Config() error
	Name() string
	Version() string
}

func RegistryGrpcService(srv GrpcService) {
	_, ok := grpcServices[srv.Name()]
	if !ok {
		grpcServices[srv.Name()] = srv
	}
}

func RegisteredGrpcServices(services []string) {
	for K := range grpcServices {
		services = append(services, K)
	}
	return
}

func GetGrpcService(name string) GrpcService {
	srv, _ := grpcServices[name]
	return srv
}

// RegistryAllGrpcServices 注册时所有的grpc服务到grpc服务中
func RegistryAllGrpcServices(server *grpc.Server) {
	for _, srv := range grpcServices {
		srv.Registry(server)
	}
}
