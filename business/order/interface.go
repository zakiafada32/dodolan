package order

type Service interface {
	FindAll() ([]Order, error)
	FindById(orderId uint32) (Order, error)
	Payment(orderId uint32, totalAmount uint64) error
}

type Repository interface {
	FindAll() ([]Order, error)
	FindById(orderId uint32) (Order, error)
	Payment(orderId uint32, totalAmount uint64) error
}
