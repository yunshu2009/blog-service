package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yunshu2009/blog-service/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response {
		Ctx : ctx
	}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	r.Cxt.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(List interface{], totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"List" : List,
		"pager":Pager {
			Page:GetPage(r.ctx),
			PageSize:GetPageSize(r.ctx),
			TotalRows:totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code":err.Code(), "msg":err.Msg()}

	details := err.Details()

	if len(details) > 0 {
		response["details"] = details
	}

	r.Cxt.JSON(err.StatusCode(), response)
}
