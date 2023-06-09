package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"postapi/db"
	view_model "postapi/model/view"
)

type PostmanPrefixStruct struct {
	Prefix string  `json:"prefix"`
}

func GetPostmanIdHandler(c *gin.Context) {
	// define io
	body := PostmanPrefixStruct{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	prefix := body.Prefix
	// call handler
	GetPostmanId(c,prefix, resp)
	// give resp
	c.JSON(http.StatusOK, resp)
}

func GetPostmanId(ctx context.Context,prefix string, resp *view_model.CommonResult) {
	data,err := GetPostmanID(ctx,prefix)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		return
	}
	resp.Data = data
}

func GetPostmanID(ctx context.Context,prefix string) ([]*view_model.PostmanID, error) {
	postmanId, err := db.GetpostmanId(ctx,prefix)
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


