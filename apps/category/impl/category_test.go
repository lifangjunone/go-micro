package impl_test

import (
	"context"
	"github.com/lifangjunone/go-micro/apps/category"
	"github.com/lifangjunone/go-micro/conf"
	"github.com/lifangjunone/go-micro/service_center"
	"testing"
)

var (
	svr category.ServiceServer
)

func TestCreateCategory(t *testing.T) {
	req := &category.CreateCategoryRequest{
		Name:       "动物",
		KeyPicture: "https://img2.baidu.com/it/u=3202947311,1179654885&fm=253&fmt=auto&app=138&f=JPEG?w=800&h=500",
	}
	ins, err := svr.CreateCategory(context.Background(), req)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(ins.Data.Name)
}

func TestQueryCategory(t *testing.T) {
	query := &category.QueryCategoryRequest{
		Keyword: "动物",
	}
	data, err := svr.QueryCategory(context.Background(), query)
	if err != nil {
		t.Error(err)
	}
	for _, item := range data.Items {
		t.Log(item.Data.Name)
	}
}

func init() {
	conf.LoadConfigFromToml("/Users/lifangjun/go-project/go-micro/etc/config.toml")
	service_center.InitAllService()
	svr = service_center.GetGrpcService(category.AppName).(category.ServiceServer)
}
