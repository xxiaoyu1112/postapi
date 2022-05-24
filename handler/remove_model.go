package handler

import (
	"net/http"
	view_model "postapi/model/view"
	"postapi/rpc"

	"github.com/gin-gonic/gin"
)

func RemoveModelHandler(c *gin.Context) {
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	modelName := c.Param("modelName")
	modelVersion := c.Param("modelVersion")
	RemoveModel(c, resp, modelName, modelVersion)
	// give resp
	c.JSON(http.StatusOK, resp)
}

func RemoveModel(ctx *gin.Context, resp *view_model.CommonResult, modelName, modelVersion string) {
	err := rpc.RemoveModel(ctx, modelName, modelVersion)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
	}
}
