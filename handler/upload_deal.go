package handler

import (
	"net/http"
	view_model "postapi/model/view"
	"postapi/rpc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UploadDealBody struct {
	Content string `json:"post_deal_content"`
}

func UploadDealHanlder(c *gin.Context) {
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	postDealId := uuid.New().String()
	body := UploadDealBody{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	postDealContent := body.Content
	UploadDeal(c, resp, postDealId, postDealContent)
	// give resp
	c.JSON(http.StatusOK, resp)
}

func UploadDeal(c *gin.Context, resp *view_model.CommonResult, postDealId, postDealContent string) {
	// middleware.Logger.Info("content :%v", postDealContent)
	err := rpc.UploadDeal(c, postDealId, postDealContent)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		return
	}
	resp.Data = postDealId
}
