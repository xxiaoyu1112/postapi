package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"postapi/db"
	view_model "postapi/model/view"
)

type GetPostmanIdBody struct {
	PostmanID  		string 	`json:"postman_id"`
}
func GetPostmanDateHandler(c *gin.Context) {
	// define io
	body := GetPostmanIdBody{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	postmanId := body.PostmanID
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	// call handler
	GetPostmanDate(c,resp,postmanId)
	// give resp
	c.JSON(http.StatusOK, resp)
}

func GetPostmanDate(ctx context.Context, resp *view_model.CommonResult,postmanId string) {
	data, err := GetpostmanDate(ctx,postmanId)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		return
	}
	resp.Data = data
}

func GetpostmanDate(ctx context.Context, postmanId string) ([]*view_model.PostmanWorkDate, error) {
	workDate, err := db.GetpostmanDate(ctx,postmanId)
	if err != nil {
		logrus.WithContext(ctx).Errorf("[GetpostmanDate] call GetAllPostmanDeal error,err:%+v", err)
		return nil, err
	}
	var workDates []*view_model.PostmanWorkDate
	for _,value := range workDate {
		item := workDateFromDBToHttp(value)
		workDates = append(workDates,item)
	}
	return workDates,nil
}

func workDateFromDBToHttp(value string) *view_model.PostmanWorkDate {
	PostmanWorkDate := &view_model.PostmanWorkDate{
		PostmanWorkDate: value,
		Value: value,
	}
	return PostmanWorkDate
}
