package category

type Service interface {
	CreateNewCategory(category Category) error
	FindAllCategory() ([]Category, error)
	FindCategoryById(id uint32) (Category, error)
	UpdateCategory(id uint32, name string, description string) (Category, error)
}

type Repository interface {
	CreateNewCategory(category Category) error
	FindAllCategory() ([]Category, error)
	FindCategoryById(id uint32) (Category, error)
	UpdateCategory(categoryId uint32, name string, description string) (Category, error)
}
