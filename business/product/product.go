package product

import (
	"time"
)

type Product struct {
	ID          uint32
	Name        string `validate:"required"`
	Description string
	Stock       uint32   `validate:"required"`
	Price       uint64   `validate:"required"`
	CategoryId  []uint32 `validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
