package user

type service struct {
	repository Repository
}

func NewUserService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateNewUser(user User) error {
	return s.repository.CreateNewUser(user)
}
