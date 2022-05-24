package handler

import (
	"context"
	"net/http"
	view_model "postapi/model/view"
	"postapi/rpc"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GetPredictDealBody struct {
	DealRegion string `json:"deal_region"`
	DealDate   string `json:"deal_date"`
	PageSize   int64  `json:"page_size"`
	PageNum    int64  `json:"page_num"`
}

func GetPredictDealHandler(c *gin.Context) {
	// define io
	body := GetPredictDealBody{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	dealRegion := body.DealRegion
	dealDate := body.DealDate
	pageSize := body.PageSize
	pageNum := body.PageNum
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	logrus.Info("dealRegion:%v", dealRegion)
	// call handler
	GetPredictDeal(c, resp, dealRegion, dealDate, int32(pageSize), int32(pageNum))
	// give resp
	c.JSON(http.StatusOK, resp)
}

func GetPredictDeal(ctx context.Context, resp *view_model.CommonResult, dealRegion, dealDate string, pageSize, pageNum int32) {
	data, err := rpc.GetPredictData(ctx, dealRegion, dealDate, pageNum, pageSize)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		return
	}
	resp.Data = data
}
