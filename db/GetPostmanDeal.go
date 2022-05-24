package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	view_model "postapi/model/view"
)

func GetPostmanDeal(ctx context.Context,region, date,postmanId,
	dealStartTime,dealEndTime string) ([]*view_model.PostmanDeal,int64,error){
	var findoptions = &options.FindOptions{}
	cur, err := postRawDeal.Find(ctx, bson.M{"post_deal_date": date, "city": region,
		"post_man_id":postmanId}, findoptions)
	if err != nil {
		log.Printf("[GetPostPredictDeal] error call mongo find,err: %v", err)
		return nil,0, err
	}
	var total int64
	total,err = postRawDeal.CountDocuments(ctx, bson.M{"post_deal_date": date, "city": region,
		"post_man_id":postmanId})
	deals := []*view_model.PostmanDeal{}
	err = cur.All(ctx, &deals)
	if err != nil {
		log.Printf("[GetPostPredictDeal] error call mongo curosr,err: %v", err)
		return nil,0, err
	}
	return deals,total,nil
}
