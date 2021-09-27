package order

type service struct {
	repository Repository
}

func NewOrderService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}
