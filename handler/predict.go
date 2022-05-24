package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	view_model "postapi/model/view"
	"postapi/rpc"

	"github.com/SpringTY/postrpc/rpc_sdk/post_model_predict"
	"github.com/gin-gonic/gin"
)

func PredictHandler(c *gin.Context) {
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	modelName := c.Param("modelName")
	modelVersion := c.Param("modelVersion")

	// ReadAll 一次性的读取 io.Reader 当中的数据
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	predictInput := view_model.PredictInput{}
	err = json.Unmarshal(jsonData, &predictInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	Predict(c, resp, modelName, modelVersion, &predictInput)
	// give resp
	c.JSON(http.StatusOK, resp)
}

func Predict(ctx context.Context, resp *view_model.CommonResult, modelName, modelVersion string, input *view_model.PredictInput) {
	rpcInput := ConvertPredictIputFromApiToRpc(input)
	order, err := rpc.Predict(ctx, modelName, modelVersion, rpcInput)
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		return
	}
	resp.Data = &view_model.PredictResult{
		Order: order,
	}
}

func ConvertPredictIputFromApiToRpc(input *view_model.PredictInput) *post_model_predict.PredictData {
	var points []*post_model_predict.Point
	for _, point := range input.Points {
		points = append(points, &post_model_predict.Point{
			Features: point,
		})
	}
	rpcInput := &post_model_predict.PredictData{
		Start:  input.Start,
		Points: points,
	}
	return rpcInput
}
