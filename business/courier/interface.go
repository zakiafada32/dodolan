package courier

type Service interface {
	FindAll() ([]CourierProvider, error)
	CreateNew(courierProvider CourierProvider) error
	Update(id uint32, name string, description string) (CourierProvider, error)
}

type Repository interface {
	FindAll() ([]CourierProvider, error)
	CreateNew(courierProvider CourierProvider) error
	Update(id uint32, name string, description string) (CourierProvider, error)
}
