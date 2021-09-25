package courier

type Service interface {
	CreateNewCourierProvider(courierProvider CourierProvider) error
}

type Repository interface {
	CreateNewCourierProvider(courierProvider CourierProvider) error
}
