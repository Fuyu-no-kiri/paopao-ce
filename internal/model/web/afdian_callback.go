package web

// AfdianCallbackReq 对应最外层的响应结构
type AfdianCallbackReq struct {
	EC   int       `json:"ec"`
	EM   string    `json:"em"`
	Data OrderData `json:"data"`
}

type AfdianCallbackResp struct {
	EC int    `json:"ec"`
	EM string `json:"em"`
}

// OrderData 对应 data 对象
type OrderData struct {
	Type  string `json:"type"`
	Order Order  `json:"order"`
}

// Order 对应 data.order
type Order struct {
	OutTradeNo     string      `json:"out_trade_no"`
	UserID         string      `json:"user_id"`
	UserPrivateID  string      `json:"user_private_id"` // 每个用户唯一，相当于微信的 unionid
	PlanID         string      `json:"plan_id"`
	Month          int         `json:"month"`
	TotalAmount    float64     `json:"total_amount,string"` // 保留两位小数的金额
	ShowAmount     float64     `json:"show_amount,string"`  // 展示用金额
	Status         int         `json:"status"`              // 订单状态码
	Remark         string      `json:"remark"`
	RedeemID       string      `json:"redeem_id"`
	ProductType    int         `json:"product_type"`    // 商品类型
	Discount       string      `json:"discount"`        // 折扣金额
	SKUs           []SKUDetail `json:"sku_detail"`      // 商品明细列表
	AddressPerson  string      `json:"address_person"`  // 收件人
	AddressPhone   string      `json:"address_phone"`   // 收件电话
	AddressAddress string      `json:"address_address"` // 收件地址
}

// SKUDetail 对应 sku_detail 数组中的每一项
type SKUDetail struct {
	SKUID   string `json:"sku_id"`   // SKU 编号
	Count   int    `json:"count"`    // 数量
	Name    string `json:"name"`     // 名称
	AlbumID string `json:"album_id"` // 关联相册 ID（可空）
	Pic     string `json:"pic"`      // 图片 URL
}
