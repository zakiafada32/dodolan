package category

type Service interface {
	CreateNew(category Category) error
	FindAll() ([]Category, error)
	FindById(id uint32) (Category, error)
	Update(id uint32, name string, description string) (Category, error)
}

type Repository interface {
	CreateNew(category Category) error
	FindAll() ([]Category, error)
	FindById(id uint32) (Category, error)
	Update(id uint32, name string, description string) (Category, error)
}
