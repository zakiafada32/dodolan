package user

type Service interface {
	CreateNewUser(user User) error
}

type Repository interface {
	CreateNewUser(user User) error
}
