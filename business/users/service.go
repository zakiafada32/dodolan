package users

type NewUserSpec struct {
	Name     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) FindUserByEmail(id string) (*User, error) {
	return s.repository.FindUserByEmail(id)
}

func (s *service) CreateNewUser(newUserSpecs NewUserSpec) error {
	return s.repository.CreateNewUser(newUserSpecs)
}
