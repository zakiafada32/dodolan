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

type checkoutCartRequestBody struct {
	PaymentID uint32 `json:"payment_id" validate:"required"`
	CourierID uint32 `json:"courier_id" validate:"required"`
}
