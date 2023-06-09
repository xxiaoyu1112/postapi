package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	view_model "postapi/model/view"
)

func GetpostmanId(ctx context.Context,prefix string) ([]string,error){
	var findOptions = &options.FindOptions{}
	//findOptions.SetLimit(1000)
	if prefix == "0000"{
		results, err := postRawDeal.Distinct(context.TODO(), "post_man_id", findOptions)
		if err != nil {
			return nil, err
		}

		var postman []string
		for _, result := range results {
			postman = append(postman, result.(string))
		}
		fmt.Println(len(postman))
		return postman,nil
	}

	// 总数为 385027
	// 前缀为 22 347602个

	// 前缀为 220 198103个   前缀为 2200
	// 前缀为 2204 4203个 只有前缀为 22041 的id
	// 前缀为 2206 27469个
	// 前缀为 2207 45671个
	// 前缀为 2208 34287个
	// 前缀为 2209 86458个

	// 前缀为 221 149499个
	// 前缀为 2210 83921个
	// 前缀为 2211 65578个   前缀为 22118 和 22119 没有符合条件的id
	// 前缀为 4398 37425个
	cursor, err := postRawDeal.Find(context.TODO(),bson.M{
		"post_man_id":primitive.Regex{
			Pattern:"^" + prefix,
			Options: "i",
		},},findOptions)
	if err != nil {
		return nil, err
	}
	var deals []*view_model.PostmanDeal
	err = cursor.All(ctx, &deals)
	if err != nil {
		log.Printf("[GetpostmanId] error call mongo curosr,err: %v", err)
		return nil,err
	}
	fmt.Println(len(deals))
	postmanCountMap := make(map[string]int,0)
	var postman []string
	for _,deal := range deals{
		postmanCountMap[deal.PostManId + "-"+ deal.PostDealDate]++
	}
	postmanIdMap := map[string]struct{}{}
	for k,val := range postmanCountMap{
		if val > 10 && val < 50{
			postmanIdMap[k[:13]] = struct{}{}
		}
	}
	for k := range postmanIdMap{
		postman = append(postman, k)
	}
	fmt.Println(len(postman))
	return postman,nil
}
