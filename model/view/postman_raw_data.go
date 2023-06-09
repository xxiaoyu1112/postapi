package view_model

type PostmanRawData struct {
	DealDate                  string `json:"DealDate,omitempty"`                 // "日期",
	GetDealTime               string `json:"GetDealTime,omitempty"`              // "接单时间",
	AppointmentTimeOne        string `json:"AppointmentTimeOne,omitempty"`       // "预约时间1",
	AppointmentTimeTwo        string `json:"AppointmentTimeTwo,omitempty"`       // "预约时间2",
	DealLongitude             string `json:"DealLongitude,omitempty"`            // "订单经度",
	DealLatitude              string `json:"DealLatitude,omitempty"`             // "订单纬度",
	DealFinishTime            string `json:"DealFinishTime,omitempty"`           // "订单揽收时间",
}

type FullPostmanRawData struct {
	PostmanRawData []*PostmanRawData `json:"postman_raw_data"`
	Total       string                `json:"total"`
}

type PostmanRawDataDeal struct {
	Raw string `bson:"raw"`
	Tag string `bson:"tag"`
}
