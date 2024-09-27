package blockData

type BlockData struct {
	TransactionHash    string `json:"transaction_hash" gorm:"primaryKey;index"  validate:"required"`
	UserAddress        string `json:"user_address" gorm:"primaryKey" validate:"required"`
	MarketId           string `json:"market_id" validate:"required"`
	OrderId            string `json:"order_id" validate:"required"`
	OpenOrderSizeUsd   string `json:"open_order_size_usd" validate:"required"`
	ClosedOrderSizeUsd string `json:"closed_order_size_usd" validate:"required"`
	OrderType          string `json:"order_type" validate:"required,oneof='vote order_matched cancelled redemption'"`
}
