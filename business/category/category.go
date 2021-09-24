package category

import "time"

type Category struct {
	ID          uint32
	Name        string `validate:"required"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
