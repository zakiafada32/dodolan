package user

type Service interface {
	GetCurrentUser(id string) (User, error)
	CreateNewUser(user User) error
	Login(email string, password string) (token string, err error)
	UpdateUser(userId string, updateData UpdateUser) (User, error)
}

type Repository interface {
	CreateNewUser(user User) error
	FindById(email string) (User, error)
	FindByEmail(email string) (User, error)
	UpdateUser(userId string, updateData UpdateUser) (User, error)
}
