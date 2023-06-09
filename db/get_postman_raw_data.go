package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	view_model "postapi/model/view"
)

func GetPostmanRawData(ctx context.Context,postmanId,postmanWordDay string) ([]*view_model.PostmanRawDataDeal,error){
	opts := options.Find().SetSort(bson.D{{"raw", 1}})
	cursor,err := postRawDeal.Find(context.TODO(),
		bson.M{"post_man_id":postmanId,"post_deal_date":postmanWordDay},opts)
	if err != nil {
		log.Printf("[GetPostmanRawData] error call mongo find,err: %v", err)
		return nil, err
	}
	var deals []*view_model.PostmanRawDataDeal
	if err = cursor.All(context.TODO(), &deals); err != nil {
		log.Printf("[GetPostmanRawData] error call mongo curosr,err: %v", err)
		return nil, err
	}
	return deals,nil
}
