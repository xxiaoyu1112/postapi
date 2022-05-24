package rpc

import (
	"context"
	"errors"
	"log"

	"github.com/SpringTY/postrpc/rpc_sdk/post_model_manage"
	"google.golang.org/grpc"
)

const (
	model_manage_address = "211.71.76.189:50053"
)

var model_manage_client post_model_manage.PostModelManageClient

func init() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(model_manage_address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	model_manage_client = post_model_manage.NewPostModelManageClient(conn)
}

func GetModelStatus(ctx context.Context) (*post_model_manage.GetModelStatesData, error) {
	req := &post_model_manage.GetModelStatesRequest{}
	resp, err := model_manage_client.GetModelStates(ctx, req)
	if err != nil {
		return nil, err
	} else if resp.Status != 0 {
		return nil, errors.New(resp.Message)
	}
	return resp.Data, nil
}
func RegisterModel(ctx context.Context, modelUrl string) (string, error) {
	req := &post_model_manage.RegisterModelRequest{
		Url: modelUrl,
	}
	resp, err := model_manage_client.RegisterModel(ctx, req)
	if err != nil {
		return "", err
	} else if resp.Status != 0 {
		return "", errors.New(resp.Message)
	}
	return resp.Message, nil
}
func RemoveModel(ctx context.Context, modelName, modelVersion string) error {
	req := &post_model_manage.RemoveModelRequest{
		ModelIdentity: &post_model_manage.ModelIdentity{
			ModelName:    modelName,
			ModelVersion: modelVersion,
		},
	}
	resp, err := model_manage_client.RemoveModel(ctx, req)
	if err != nil {
		return err
	} else if resp.Status != 0 {
		return errors.New(resp.Message)
	}
	return nil
}

func UpdateModel(ctx context.Context, modelName, modelVersion string, maxWorkers, minWokers int32) error {
	req := &post_model_manage.UpdateModelConfigRequest{
		ModelIdentity: &post_model_manage.ModelIdentity{ModelName: modelName, ModelVersion: modelVersion},
		ModelConfig: &post_model_manage.ModelConfig{
			MinWorkers: minWokers,
			MaxWorkers: maxWorkers,
		},
	}
	resp, err := model_manage_client.UpdateModelConfig(ctx, req)
	if err != nil {
		return err
	} else if resp.Status != 0 {
		return errors.New(resp.Message)
	}
	return nil
}
