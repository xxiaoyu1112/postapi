package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"postapi/db"
	view_model "postapi/model/view"
)

func GetPostmanIdHandler(c *gin.Context) {
	// define io
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	// call handler
	GetPostmanId(c, resp)
	// give resp
	c.JSON(http.StatusOK, resp)
}

func GetPostmanId(ctx context.Context, resp *view_model.CommonResult) {
	data,err := GetPostmanID(ctx)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		return
	}
	resp.Data = data
}

func GetPostmanID(ctx context.Context) ([]*view_model.PostmanID, error) {
	postmanId, err := db.GetpostmanId(ctx)
	if err != nil {
		logrus.WithContext(ctx).Errorf("[GetpostmanId] call GetAllPostmanDeal error,err:%+v", err)
		return nil, err
	}
	var postmanID []*view_model.PostmanID
	for _,value := range postmanId{
		item := PostmanIDFromDBtoHttp(value)
		postmanID = append(postmanID,item)
	}
	return postmanID,nil
}

func PostmanIDFromDBtoHttp(value string) *view_model.PostmanID {
	postmanId := &view_model.PostmanID{
		PostmanID: value,
		Value: value,
	}
	return postmanId
}


