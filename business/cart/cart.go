package cart

type Cart struct {
	Item        []CartItem `json:"item"`
	TotalAmount uint64     `json:"total_amount"`
}

type CartItem struct {
	ProductID   uint32 `json:"product_id" validate:"required"`
	Quantity    uint32 `json:"quantity" validate:"required"`
	TotalAmount uint64 `json:"total_amount"`
}
