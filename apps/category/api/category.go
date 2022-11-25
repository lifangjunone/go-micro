package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/lifangjunone/go-micro/apps/category"
	"github.com/lifangjunone/go-micro/common/response"
)

func (h *handler) CreateCategory(request *restful.Request, resp *restful.Response) {
	entity := category.NewDefaultCreateCategoryRequest()
	err := request.ReadEntity(entity)
	if err != nil {
		response.Failed(resp.ResponseWriter, err)
		return
	}
	ins, err := h.categorySvr.CreateCategory(request.Request.Context(), entity)
	if err != nil {
		response.Failed(resp.ResponseWriter, err)
		return
	}
	response.Success(resp.ResponseWriter, ins)
}

func (h *handler) QueryCategory(request *restful.Request, resp *restful.Response) {
	param := request.QueryParameter("keyword")
	query := category.QueryCategoryRequest{
		Keyword: param,
	}
	data, err := h.categorySvr.QueryCategory(request.Request.Context(), &query)
	if err != nil {
		response.Failed(resp.ResponseWriter, err)
		return
	}
	response.Success(resp.ResponseWriter, data)
}
