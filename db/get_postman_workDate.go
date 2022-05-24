package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func GetpostmanDate(ctx context.Context,postmanId string) ([]string,error){
	results, err := postRawDeal.Distinct(context.TODO(), "post_deal_date",
		bson.M{"post_man_id":postmanId})
	if err != nil {
		return nil, err
	}
	var postDate []string
	// tags = append(tags, results...)
	for _, result := range results {
		postDate = append(postDate, result.(string))
	}
	return postDate,nil
}
