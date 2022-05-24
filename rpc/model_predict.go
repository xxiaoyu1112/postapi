package rpc

import (
	"context"
	"errors"
	"log"

	"github.com/SpringTY/postrpc/rpc_sdk/post_model_predict"
	"google.golang.org/grpc"
)

const (
	model_predict_address = "211.71.76.189:50054"
)

var model_predict_client post_model_predict.PostModelPredictClient

func init() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(model_predict_address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	model_predict_client = post_model_predict.NewPostModelPredictClient(conn)
}

func Predict(ctx context.Context, modelName, modelVersion string, input *post_model_predict.PredictData) ([]int32, error) {
	req := &post_model_predict.PredictRequest{
		ModelName:    modelName,
		ModelVersion: modelVersion,
		Data:         input,
	}
	resp, err := model_predict_client.Predict(ctx, req)
	if err != nil {
		return nil, err
	} else if resp.Status != 0 {
		return nil, errors.New(resp.Message)
	}
	return resp.Data.Order, nil
}
