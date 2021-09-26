package product

type Service interface {
	FindById(id uint32) (ProductAtt, error)
	FindAll() ([]ProductAtt, error)
	CreateNew(product Product) error
	Update(id uint32, updateData ProductUpdate) (ProductAtt, error)
}

type Repository interface {
	FindById(id uint32) (ProductAtt, error)
	FindAll() ([]ProductAtt, error)
	CreateNew(product Product) error
	Update(id uint32, updateData ProductUpdate) (ProductAtt, error)
}
