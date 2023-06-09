package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"postapi/db"
	view_model "postapi/model/view"
	"strings"
)

type GetPostmanRawDataBody struct {
	PostmanID      string `json:"postman_id"`
	PostmanWorkday string `json:"postman_workday"`
}
func GetPostmanRawDataHandler(c *gin.Context) {
	// define io
	body := GetPostmanRawDataBody{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	postmanId := body.PostmanID
	postmanWordDay := body.PostmanWorkday
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	// call handler
	GetPostmanRawData(c,resp,postmanId,postmanWordDay)
	// give resp
	c.JSON(http.StatusOK, resp)
}

func GetPostmanRawData(ctx context.Context, resp *view_model.CommonResult,postmanId,postmanWordDay string) {
	data, err := NewGetPostmanRawData(ctx,postmanId,postmanWordDay)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		return
	}
	resp.Data = data
}

func NewGetPostmanRawData(ctx context.Context, postmanId,postmanWordDay string) ([]*view_model.PostmanRawData, error) {
	PostmanRawData, err := db.GetPostmanRawData(ctx,postmanId,postmanWordDay)
	if err != nil {
		logrus.WithContext(ctx).Errorf("[NewGetPostmanRawData] call NewGetPostmanRawData error,err:%+v", err)
		return nil, err
	}
	PostmanRawDatas := PostmanRawDataFromDBToHttp(PostmanRawData)
	return PostmanRawDatas,nil
}

func PostmanRawDataFromDBToHttp(PostmanRawData []*view_model.PostmanRawDataDeal) []*view_model.PostmanRawData {
	var FullPostmanRawData []*view_model.PostmanRawData
	for _, rawData := range PostmanRawData {
		res := strings.Split(rawData.Raw, ",")
		FullPostmanRawData = append(FullPostmanRawData, &view_model.PostmanRawData{
			DealDate:                  res[0],
			GetDealTime:               res[4],
			AppointmentTimeOne:        res[5],
			AppointmentTimeTwo:        res[6],
			DealLongitude:             res[7],
			DealLatitude:              res[8],
			DealFinishTime:            res[12],
		})
	}
	return FullPostmanRawData
}
