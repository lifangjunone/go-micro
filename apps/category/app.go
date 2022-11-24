package category

import "time"

const (
	AppName = "category"
	Version = "v1.0.0"
)

func NewCategorySet() *CategorySet {
	return &CategorySet{
		Items: []*Category{},
	}
}

func NewCategory(req *CreateCategoryRequest) *Category {
	return &Category{
		CreateAt: time.Now().UnixMicro(),
		UpdateAt: time.Now().UnixMicro(),
		Data:     req,
	}
}

func NewDefaultCategory() *Category {
	return &Category{
		Data: &CreateCategoryRequest{},
	}
}

func (c *CategorySet) Add(item *Category) {
	c.Items = append(c.Items, item)
}
