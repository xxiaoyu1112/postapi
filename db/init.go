package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client          *mongo.Client
	err             error
	db              *mongo.Database
	postPredictDeal *mongo.Collection
	postRawDeal     *mongo.Collection
)

func init() {

	//1.建立连接
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://admin:123456@211.71.76.189:27017").SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Print(err)
		return
	}
	//2.选择数据库
	db = client.Database("post_data")

	//3.选择表
	postPredictDeal = db.Collection("post_predict")
	postRawDeal = db.Collection("post_deal")
}
