package handler

import (
	"context"
	"net/http"
	view_model "postapi/model/view"
	"postapi/rpc"

	"github.com/gin-gonic/gin"
)

func GetGenTaskInfoHandler(c *gin.Context) {
	// define io
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	taskId := c.Param("taskId")
	// call handler
	GetGenTaskInfo(c, resp, taskId)
	// give resp
	c.JSON(http.StatusOK, resp)
}

func GetGenTaskInfo(ctx context.Context, resp *view_model.CommonResult, taskId string) {
	data, err := rpc.GetTaskInfo(ctx, taskId)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		return
	}
	resp.Data = data
}
