package user

type Service interface {
	CreateNewUser(user User) error
	Login(email string, password string) (string, error)
}

type Repository interface {
	CreateNewUser(user User) error
	FindByEmail(email string) (User, error)
}
