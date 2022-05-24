package view_model


type PostmanDeal struct {
	Raw string `bson:"raw"`
	Tag string `bson:"tag"`
	PostDealDate string `bson:"post_deal_date"`
	City string `bson:"city"`
	PostManId string `bson:"post_man_id"`
}


