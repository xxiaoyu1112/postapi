package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetpostmanId(ctx context.Context) ([]string,error){
	var findoptions = &options.FindOptions{}
	//findoptions.SetLimit(10)
	results, err := postRawDeal.Distinct(context.TODO(), "post_man_id", findoptions)
	if err != nil {
		return nil, err
	}
	var postman []string
	for _, result := range results {
		postman = append(postman, result.(string))
	}
	return postman,nil
}
