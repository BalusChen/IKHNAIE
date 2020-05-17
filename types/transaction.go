package types

import "time"

type Transaction struct {
	// FoodID  int64   `json:"food_id"` /* key */
	TradeTime  time.Time `json:"trade_time"`  // 交易时间
	TradePlace string    `json:"trade_place"` // 交易地点
	SellerName string    `json:"seller_name"` // 售卖者名字
	SellerID   string    `json:"seller_id"`   // 售卖者身份证号码
	BuyerName  string    `json:"buyer_name"`  // 购买者名字
	BuyerID    string    `json:"buyer_id"`    // 购买者身份证号码
	Number     int64     `json:"number"`      // 交易数目
	Price      float64   `json:"price"`       // 单价
}
