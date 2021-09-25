package payment

import "github.com/zakiafada32/retail/business/utils"

type service struct {
	repository Repository
}

func NewPaymentService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateNewPaymentProvider(paymentProvider PaymentProvider) error {
	err := utils.GetValidator().Struct(paymentProvider)
	if err != nil {
		return err
	}
	return s.repository.CreateNewPaymentProvider(paymentProvider)
}
