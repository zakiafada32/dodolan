package payment

import "github.com/zakiafada32/retail/business/utils"

type service struct {
	repository Repository
}

func NewPaymentService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) CreateNewPaymentProvider(provider PaymentProvider) error {
	err := utils.GetValidator().Struct(provider)
	if err != nil {
		return err
	}
	return s.repository.CreateNewPaymentProvider(provider)
}
