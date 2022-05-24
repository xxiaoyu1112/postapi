package view_model

type CommonResult struct {
	Message string      `json:"message"`
	Status  int32       `json:"status"`
	Data    interface{} `json:"data"`
}
