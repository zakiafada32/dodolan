package product

import (
	"time"

	"github.com/zakiafada32/retail/business/category"
)

type Product struct {
	ID           uint32    `json:"id"`
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description"`
	Stock        uint32    `json:"stock" validate:"required"`
	Price        uint64    `json:"price" validate:"required"`
	CategoriesId []uint32  `json:"categories_id" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ProductAtt struct {
	ID          uint32              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Stock       uint32              `json:"stock" `
	Price       uint64              `json:"price"`
	Categories  []category.Category `json:"categories"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}
