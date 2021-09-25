package cart

type service struct {
	repository Repository
}

func NewCartService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) AddCartItem(userId string, productId uint32, quantity uint32) error {
	return nil
}
