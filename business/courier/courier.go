package courier

import "time"

type CourierProvider struct {
	ID          uint32    `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
