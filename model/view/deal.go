package view_model

type Deal struct {
	Id          string      `json:"id"`
	Start       string      `json:"start"`
	GroundTruth []int       `json:"groundTruth"`
	Points      [][]float64 `json:"points"`
	Length      int         `json:"length"`
}

type GetDealsData struct {
	Deals []*Deal `json:"deals"`
}
