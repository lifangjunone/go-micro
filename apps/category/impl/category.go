package impl

import (
	"context"
	"github.com/lifangjunone/go-micro/apps/category"
)

func (i *impl) CreateCategory(ctx context.Context, req *category.CreateCategoryRequest) (*category.Category, error) {
	ins := category.NewCategory(req)
	i.save(ins)
	return ins, nil
}
func (i *impl) QueryCategory(ctx context.Context, req *category.QueryCategoryRequest) (*category.CategorySet, error) {
	data, _ := i.query(req)
	return data, nil
}
