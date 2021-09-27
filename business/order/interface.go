package order

type Service interface {
	FindAll() ([]Order, error)
	FindById(orderId uint32) (Order, error)
}

type Repository interface {
	FindAll() ([]Order, error)
	FindById(orderId uint32) (Order, error)
}
