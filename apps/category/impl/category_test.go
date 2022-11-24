package impl_test

import (
	"context"
	"github.com/lifangjunone/go-micro/apps/category"
	"github.com/lifangjunone/go-micro/conf"
	"testing"
)

var (
	svr category.ServiceServer
)

func TestCreateCategory(t *testing.T) {
	req := &category.CreateCategoryRequest{
		Name:       "风景",
		KeyPicture: "https://img2.baidu.com/it/u=867579726,2670217964&fm=253&app=120&size=w931&n=0&f=JPEG&fmt=auto?sec=1669482000&t=5ac0a1632cc1f5e1bcff8e52576199b8",
	}
	ins, err := svr.CreateCategory(context.Background(), req)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(ins.Data.Name)
}

func TestQueryCategory(t *testing.T) {
	query := &category.QueryCategoryRequest{
		Keyword: "风景",
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
}
