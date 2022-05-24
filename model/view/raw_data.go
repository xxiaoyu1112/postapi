package view_model

type RawData struct {
	DealDate                  string `json:"DealDate,omitempty"`                 // "日期",
	RegionId                  string `json:"RegionId,omitempty"`                 // "运营区id",
	City                      string `json:"City,omitempty"`                     // "城市",
	PostManId                 string `json:"PostManId,omitempty"`                // "快递员id",
	GetDealTime               string `json:"GetDealTime,omitempty"`              // "接单时间",
	AppointmentTimeOne        string `json:"AppointmentTimeOne,omitempty"`       // "预约时间1",
	AppointmentTimeTwo        string `json:"AppointmentTimeTwo,omitempty"`       // "预约时间2",
	DealLongitude             string `json:"DealLongitude,omitempty"`            // "订单经度",
	DealLatitude              string `json:"DealLatitude,omitempty"`             // "订单纬度",
	DealRegionId              string `json:"DealRegionId,omitempty"`             // "订单所属区块id",
	DealBlockTypeId           string `json:"DealBlockTypeId,omitempty"`          // "区块类型id",
	DealBlockType             string `json:"DealBlockType,omitempty"`            // "区块类型",
	DealFinishTime            string `json:"DealFinishTime,omitempty"`           // "订单揽收时间",
	RecentFinishDealTime      string `json:"RecentFinishDealTime,omitempty"`     // "揽收最近时间",
	RecentFinishDealLongitude string `json:"RecentFinishDealLongitude,omitempty"`// "揽收最近经度",
	RecentFinishDealLatitude  string `json:"RecentFinishDealLatitude,omitempty"` // "揽收最近纬度",
	FinishDealPrecision       string `json:"FinishDealPrecision,omitempty"`      // "揽收轨迹精度",
	RecentGetDealTime         string `json:"RecentGetDealTime,omitempty"`        // "接单最近时间",
	RecentGetDealLongitude    string `json:"RecentGetDealLongitude,omitempty"`   // "接单最近经度",
	RecentGetDealLatitude     string `json:"RecentGetDealLatitude,omitempty"`    // "接单最近纬度",
	GetDealPrecision          string `json:"GetDealPrecision,omitempty"`         // "接单轨迹精度",
}

type FullPostmanDeal struct {
	PostmanDeal []*RawData `json:"full_Postman_Deal"`
	Total       string                `json:"total"`
}
