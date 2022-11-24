package service_center

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"strings"
)

var (
	restfulServices = map[string]RestfulService{}
)

type RestfulService interface {
	Registry(service *restful.WebService)
	Config() error
	Name() string
	Version() string
}

// RegistryRestfulService 注册服务到服务中心
func RegistryRestfulService(service RestfulService) {
	_, ok := restfulServices[service.Name()]
	if !ok {
		restfulServices[service.Name()] = service
	}
}

// RegisteredRestfulServices  已经注册的服务
func RegisteredRestfulServices() (services []string) {
	for k := range restfulServices {
		services = append(services, k)
	}
	return
}

// GetRestfulService 注册中心发现服务
func GetRestfulService(name string) RestfulService {
	service, ok := restfulServices[name]
	if !ok {
		return nil
	}
	return service
}

func RegistryAllRestfulServices(pathPrefix string, root *restful.Container) {
	for _, srv := range restfulServices {
		pathPrefix = strings.TrimSuffix(pathPrefix, "/")
		ws := new(restful.WebService)
		ws.Path(fmt.Sprintf("%s/%s/%s", pathPrefix, srv.Version(), srv.Name())).
			Consumes(restful.MIME_JSON, restful.MIME_XML).
			Produces(restful.MIME_JSON, restful.MIME_XML)
		srv.Registry(ws)
		root.Add(ws)
	}
}
