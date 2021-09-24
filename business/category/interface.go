package category

type Service interface {
	CreateNewCategory(category Category) error
}

type Repository interface {
	CreateNewCategory(category Category) error
}
