package order

type paymentRequestBody struct {
	OrderId     uint32 `json:"order_id" validate:"required"`
	TotalAmount uint64 `json:"total_amount" validate:"required"`
}
