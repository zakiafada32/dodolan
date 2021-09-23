package users

type Service interface {
	FindUserByEmail(email string) (*User, error)
	CreateNewUser(newUserSpec NewUserSpec) error
}

type Repository interface {
	FindUserByEmail(id string) (*User, error)
	CreateNewUser(newUserSpec NewUserSpec) error
}
