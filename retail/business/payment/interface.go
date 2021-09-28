package payment

type Service interface {
	FindAll() ([]PaymentProvider, error)
	CreateNew(provider PaymentProvider) error
	Update(id uint32, name, description string) (PaymentProvider, error)
}

type Repository interface {
	FindAll() ([]PaymentProvider, error)
	CreateNew(provider PaymentProvider) error
	Update(id uint32, name, description string) (PaymentProvider, error)
}
