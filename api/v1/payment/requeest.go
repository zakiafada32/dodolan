package payment

import "github.com/zakiafada32/retail/business/payment"

type paymentProviderRequestBody struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (req *paymentProviderRequestBody) convertToPaymentProviderBusiness() payment.PaymentProvider {
	return payment.PaymentProvider{
		Name:        req.Name,
		Description: req.Description,
	}
}
