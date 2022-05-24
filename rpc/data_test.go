package rpc

import (
	"context"
	"fmt"
	"testing"
)

func TestGetRawData(t *testing.T) {
	ctx := context.Background()

	res, err := GetRawData(ctx, "mongodb", "杭州市-20210330", 1, 10)
	if err != nil {
		fmt.Printf("%+v", err)
	} else {
		fmt.Printf("%+v", res)
	}
}

func TestGetRawDataTree(t *testing.T) {
	ctx := context.Background()
	res, err := GetRawDataTree(ctx)
	if err != nil {
		fmt.Printf("%+v", err)
	} else {
		fmt.Printf("%+v", res)
	}
}
