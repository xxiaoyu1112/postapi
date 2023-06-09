package handler

import (
	"net/http"
	view_model "postapi/model/view"
	"postapi/rpc"

	"github.com/gin-gonic/gin"
)

type UpdateModelBody struct {
	MinWorkers int32 `json:"min_workers"`
	MaxWorkers int32 `json:"max_workers"`
}

func UpdateModelHandler(c *gin.Context) {
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	modelName := c.Param("modelName")
	modelVersion := c.Param("modelVersion")
	body := UpdateModelBody{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	minWorkers := body.MinWorkers
	maxWorkers := body.MaxWorkers
	UpdateModel(c, resp, modelName, modelVersion, int32(maxWorkers), int32(minWorkers))
	// give resp
	c.JSON(http.StatusOK, resp)
}

func UpdateModel(ctx *gin.Context, resp *view_model.CommonResult, modelName, modelVersion string, maxWorkers, minWorkers int32) {
	err := rpc.UpdateModel(ctx, modelName, modelVersion, maxWorkers, minWorkers)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
	}
}
