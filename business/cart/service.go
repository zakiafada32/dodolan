package cart

type service struct {
	repository Repository
}

func NewCartService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) AddCartItem(userId string, productId uint32, quantity uint32) error {
	return nil
}
