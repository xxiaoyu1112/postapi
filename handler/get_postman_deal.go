package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/issue9/conv"
	"github.com/sirupsen/logrus"
	"net/http"
	"postapi/db"
	view_model "postapi/model/view"
	"strings"
)

type GetPostmanDealBody struct {
	DealRegion 		string 	`json:"deal_region"`
	PostmanID  		string 	`json:"postman_id"`
	DealDate   		string 	`json:"deal_date"`
	DealStartTime   string 	`json:"deal_start_time"`
	DealEndTime   	string 	`json:"deal_end_time"`
}

func GetPostmanDealHandler(c *gin.Context) {
	// define io
	body := GetPostmanDealBody{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	dealRegion := body.DealRegion
	dealDate := body.DealDate
	postmanId := body.PostmanID
	dealStartTime := body.DealStartTime
	dealEndTime := body.DealEndTime
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	logrus.Info("postmanId:%v", postmanId)
	// call handler
	GetPostmanDeal(c, resp, dealRegion, dealDate,postmanId,dealStartTime,dealEndTime)
	// give resp
	c.JSON(http.StatusOK, resp)
}

func GetPostmanDeal(ctx context.Context, resp *view_model.CommonResult,
	dealRegion, dealDate,postmanId,dealStartTime,dealEndTime string) {
	data,err := GetAllPostmanDeal(ctx,dealRegion, dealDate,postmanId,dealStartTime,dealEndTime)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		return
	}
	resp.Data = data
}

func GetAllPostmanDeal(ctx context.Context, dealRegion, dealDate, postmanId,
	dealStartTime, dealEndTime string) (*view_model.FullPostmanDeal,error) {
	fullPostmanData,total, err := db.GetPostmanDeal(ctx,dealRegion, dealDate,postmanId,dealStartTime,dealEndTime)
	if err != nil {
		logrus.WithContext(ctx).Errorf("[GetAllPostmanDeal] call GetAllPostmanDeal error,err:%+v", err)
		return nil, err
	}
	fullPostmanDeal := new(view_model.FullPostmanDeal)
	fullPostmanDeal.PostmanDeal = BuildData(fullPostmanData)
	fullPostmanDeal.Total = conv.MustString(total)
	return fullPostmanDeal,nil
}

func BuildData(postmanData []*view_model.PostmanDeal) []*view_model.RawData{
	var RawDatas []*view_model.RawData
	for _, rawData := range postmanData {
		res := strings.Split(rawData.Raw, ",")
		RawDatas = append(RawDatas, &view_model.RawData{
			DealDate:                  res[0],
			RegionId:                  res[1],
			City:                      res[2],
			PostManId:                 res[3],
			GetDealTime:               res[4],
			AppointmentTimeOne:        res[5],
			AppointmentTimeTwo:        res[6],
			DealLongitude:             res[7],
			DealLatitude:              res[8],
			DealRegionId:              res[9],
			DealBlockTypeId:           res[10],
			DealBlockType:             res[11],
			DealFinishTime:            res[12],
			RecentFinishDealTime:      res[13],
			RecentFinishDealLongitude: res[14],
			RecentFinishDealLatitude:  res[15],
			FinishDealPrecision:       res[16],
			RecentGetDealTime:         res[17],
			RecentGetDealLongitude:    res[18],
			RecentGetDealLatitude:     res[19],
			GetDealPrecision:          res[20],
		})
	}
	return RawDatas
}

