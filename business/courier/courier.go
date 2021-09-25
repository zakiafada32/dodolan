package courier

import "time"

type CourierProvider struct {
	ID          uint32
	Name        string `validate:"required"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
