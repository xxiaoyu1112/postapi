package handler

import (
	"context"
	"net/http"
	view_model "postapi/model/view"
	"postapi/rpc"

	"github.com/gin-gonic/gin"
)

func GetRawDataTreeHandler(c *gin.Context) {
	// define io
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	// call handler
	GetRawDealTree(c, resp)
	// give resp
	c.JSON(http.StatusOK, resp)
}

func GetRawDealTree(ctx context.Context, resp *view_model.CommonResult) {
	data, err := rpc.GetRawDataTree(ctx)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		return
	}
	resp.Data = data
}
