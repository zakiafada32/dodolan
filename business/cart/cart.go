package cart

type Cart struct {
	UserId      string     `json:"user_id" validate:"required"`
	Item        []CartItem `json:"item"`
	TotalAmount uint64     `json:"total_amount"`
}

type CartItem struct {
	ProductId   uint32 `json:"product_id" validate:"required"`
	Quantity    uint32 `json:"quantity" validate:"required"`
	TotalAmount uint64 `json:"total_amount"`
}
