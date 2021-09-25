package cart

type Cart struct {
	UserId      string `validate:"required"`
	Item        []CartItem
	TotalAmount uint64
}

type CartItem struct {
	ProductId   uint32 `validate:"required"`
	Quantity    uint32 `validate:"required"`
	TotalAmount uint64
}
