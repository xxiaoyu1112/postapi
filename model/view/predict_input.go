package view_model

type PredictInput struct {
	Points [][]float32 `json:"feature"`
	Start  int32       `json:"start"`
}
