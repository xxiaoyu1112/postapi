package handler

import (
	"net/http"
	view_model "postapi/model/view"
	"postapi/rpc"

	"github.com/gin-gonic/gin"
)

type RegisterModelBody struct {
	ModelUrl string `json:"model_url"`
}

func RegisterModelHandler(c *gin.Context) {
	body := RegisterModelBody{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	toast, err := RegisterModel(c, &body)
	resp := &view_model.CommonResult{
		Message: toast,
		Status:  0,
	}
	if err != nil {
		resp.Message = err.Error()
		resp.Status = 1
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Message = toast
	c.JSON(http.StatusOK, resp)
}
func RegisterModel(ctx *gin.Context, body *RegisterModelBody) (string, error) {
	//调用post_model_manage 逻辑
	toast, err := rpc.RegisterModel(ctx, body.ModelUrl)
	if err != nil {
		return "", err
	}
	return toast, nil
}
