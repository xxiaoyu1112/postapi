package handler

import (
	"context"
	"net/http"
	view_model "postapi/model/view"
	"postapi/rpc"

	"github.com/gin-gonic/gin"
)

type GetRawDealBody struct {
	Tag      string `json:"tag"` // deal id
	Source   string `json:"source"`
	PageSize int64  `json:"page_size"`
	PageNum  int64  `json:"page_num"`
}

func GetRawDealHandler(c *gin.Context) {
	// define io
	body := GetRawDealBody{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	dealTag := body.Tag
	dealSource := body.Source
	pageSize := body.PageSize
	pageNum := body.PageNum
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	// call handler
	GetRawDeal(c, resp, dealTag, dealSource, int32(pageSize), int32(pageNum))
	// give resp
	c.JSON(http.StatusOK, resp)
}

func GetRawDeal(ctx context.Context, resp *view_model.CommonResult, dealTag, dealSource string, pageSize, pageNum int32) {
	data, err := rpc.GetRawData(ctx, dealTag, dealSource, pageNum, pageSize)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		return
	}
	resp.Data = data
}
