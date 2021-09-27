package cart

import "github.com/zakiafada32/retail/business/cart"

type cartRequestBody struct {
	ProductID uint32 `json:"product_id" validate:"required"`
	Quantity  int32  `json:"quantity" validate:"required"`
}

func (r *cartRequestBody) convertToCartItemBusiness() cart.CartItem {
	return cart.CartItem{
		ProductID: r.ProductID,
		Quantity:  uint32(r.Quantity),
	}
}

type deleteCartItemRequestBody struct {
	ProductsID []uint32 `json:"products_id" validate:"required"`
}
