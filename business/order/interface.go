package order

type Service interface {
	FindById(orderId uint32) (Order, error)
}

type Repository interface {
	FindById(orderId uint32) (Order, error)
}
