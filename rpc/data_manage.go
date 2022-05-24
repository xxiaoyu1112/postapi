package rpc

import (
	"context"
	"errors"
	"log"

	"github.com/SpringTY/postrpc/rpc_sdk/post_data_manage"
	"google.golang.org/grpc"
)

const (
	data_manage_address = "211.71.76.189:50052"
)

var data_manage_client post_data_manage.PostDataManageClient

func init() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(data_manage_address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	data_manage_client = post_data_manage.NewPostDataManageClient(conn)
}
func GetPredictData(ctx context.Context, dealRegion, dealDate string, pageNum, pageSize int32) (*post_data_manage.PostPredictData, error) {
	req := &post_data_manage.GetPostPredictDataRequest{
		PageSize:   pageSize,
		PageNum:    pageNum,
		DealDate:   dealDate,
		DealRegion: dealRegion,
	}
	resp, err := data_manage_client.GetPostPredictData(ctx, req)
	if err != nil {
		return nil, err
	} else if resp.Status != 0 {
		return nil, errors.New(resp.Message)
	}
	return resp.PostPredictData, nil
}
func GetRawData(ctx context.Context, tag, source string, pageNum, pageSize int32) (*post_data_manage.RawDataPage, error) {
	req := &post_data_manage.GetRawDataRequest{
		Source:   source,
		Id:       tag,
		PageSize: pageSize,
		PageNum:  pageNum,
	}
	resp, err := data_manage_client.GetRawData(ctx, req)
	if err != nil {
		return nil, err
	} else if resp.Status != 0 {
		return nil, errors.New(resp.Message)
	}
	return resp.Data, nil
}
func GetRawDataTree(ctx context.Context) (*post_data_manage.RawDataTree, error) {
	req := &post_data_manage.GetRawDataTreeRequest{}
	resp, err := data_manage_client.GetRawDataTree(ctx, req)
	if err != nil {
		return nil, err
	} else if resp.Status != 0 {
		return nil, errors.New(resp.Message)
	}
	return resp.Tree, nil
}

func GetTaskInfo(ctx context.Context, taskId string) (*post_data_manage.PredictTaskStatus, error) {
	req := &post_data_manage.GetPostPredictTaskStatusRequest{
		TaskId: taskId,
	}
	resp, err := data_manage_client.GetPostPredictTaskStatus(ctx, req)
	if err != nil {
		return nil, err
	} else if resp.Status != 0 {
		return nil, errors.New(resp.Message)
	}
	return resp.PredictTaskStatus, nil
}

func GeneratePostPredictData(ctx context.Context, tag string) (string, error) {
	req := &post_data_manage.GeneratePostPredictDataRequest{
		Tag: tag,
	}
	resp, err := data_manage_client.GeneratePostPredictData(ctx, req)
	if err != nil {
		return "nil", err
	} else if resp.Status != 0 {
		return "nil", errors.New(resp.Message)
	}
	return resp.TaskId, nil
}
