package rpc

import (
	"context"
	"errors"
	"log"

	"github.com/SpringTY/postrpc/rpc_sdk/post_data_collect"
	"google.golang.org/grpc"
)

const (
	data_collect_address = "211.71.76.189:50051"
)

var data_collect_client post_data_collect.PostDataCollectClient

func init() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(data_collect_address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	data_collect_client = post_data_collect.NewPostDataCollectClient(conn)
}

func UploadDeal(ctx context.Context, postDealId, postDealContent string) error {
	req := &post_data_collect.CollectPostDealRequest{
		PostDealId:      postDealId,
		PostDealContent: postDealContent,
	}
	resp, err := data_collect_client.CollectPostDeal(ctx, req)
	if err != nil {
		return err
	} else if resp.Status != 0 {
		return errors.New(resp.Message)
	}
	return nil
}
