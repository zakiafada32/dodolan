package payment

type Service interface {
	CreateNewPaymentProvider(paymentProvider PaymentProvider) error
}

type Repository interface {
	CreateNewPaymentProvider(paymentProvider PaymentProvider) error
}
