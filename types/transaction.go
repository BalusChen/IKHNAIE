package types

import "time"

type Transaction struct {
	// FoodID  int64   `json:"food_id"` /* key */
	TradeAt    time.Time `json:"trade_at"`    // 交易时间
	SellerName string    `json:"seller_name"` // 售卖者名字
	SellerID   string    `json:"seller_id"`   // 售卖者身份证号码
	BuyerName  string    `json:"buyer_name"`  // 购买者名字
	BuyerID    string    `json:"buyer_id"`    // 购买者身份证号码
	Address    string    `json:"address"`     // 交易地点
	Number     int64     `json:"number"`      // 交易数目
	Price      float64   `json:"price"`       // 单价
}

type Product struct {
	FoodID      int64         `json:"food_id"`     // 农产品 ID（唯一标识）
	FoodName    string        `json:"food_name"`   // 农产品名
	Birthday    time.Time     `json:"birthday"`    // 生产日期
	ShelfLife   time.Duration `json:"shelf_life"`  // 保质期
	Ingredients []Product     `json:"ingredients"` // 组成材料
}
