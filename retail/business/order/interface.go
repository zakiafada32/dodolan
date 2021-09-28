package order

type Service interface {
	FindAll(userId string) ([]Order, error)
	FindById(userId string, orderId uint32) (Order, error)
	Payment(userId string, orderId uint32, totalAmount uint64) error
	Courier(userId string, orderId uint32) error
}

type Repository interface {
	FindAll(userId string) ([]Order, error)
	FindById(userId string, orderId uint32) (Order, error)
	Payment(userId string, orderId uint32, totalAmount uint64) error
	Courier(userId string, orderId uint32) error
}
