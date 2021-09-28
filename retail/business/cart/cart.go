package cart

type Cart struct {
	Items       []CartItemAtt `json:"item"`
	TotalAmount uint64        `json:"total_amount"`
}

type CartItem struct {
	ProductID   uint32 `json:"product_id" validate:"required"`
	Quantity    uint32 `json:"quantity" validate:"required"`
	TotalAmount uint64 `json:"total_amount"`
}

type CartItemAtt struct {
	Product     CartProduct `json:"product"`
	Quantity    uint32      `json:"quantity"`
	TotalAmount uint64      `json:"total_amount"`
}

type CartProduct struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64 `json:"price"`
}
