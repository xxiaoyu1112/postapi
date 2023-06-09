package handler

import (
	"context"
	"net/http"
	view_model "postapi/model/view"
	"postapi/rpc"

	"github.com/gin-gonic/gin"
)

func GetModelStatesHandler(c *gin.Context) {
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	GetModelStates(c, resp)
	// give resp
	c.JSON(http.StatusOK, resp)
}

func GetModelStates(ctx context.Context, resp *view_model.CommonResult) {
	modelStates, err := rpc.GetModelStatus(ctx)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		return
	}
	resp.Data = modelStates
}
