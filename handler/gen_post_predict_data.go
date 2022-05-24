package handler

import (
	"context"
	"net/http"
	view_model "postapi/model/view"
	"postapi/rpc"

	"github.com/gin-gonic/gin"
)

func GenPostPredictDataHandler(c *gin.Context) {
	// define io
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	tag := c.Param("tag")
	// call handler
	GenPostPredictData(c, resp, tag)
	// give resp
	c.JSON(http.StatusOK, resp)
}

func GenPostPredictData(ctx context.Context, resp *view_model.CommonResult, tag string) {
	taskId, err := rpc.GeneratePostPredictData(ctx, tag)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		return
	}
	resp.Data = taskId
}
