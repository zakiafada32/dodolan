package product

import (
	"errors"
	"time"

	"github.com/zakiafada32/retail/business/product"
	"github.com/zakiafada32/retail/modules/category"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

type Product struct {
	ID          uint32 `gorm:"primaryKey"`
	Name        string
	Description string
	Stock       uint32
	Price       uint64
	Categories  []category.Category `gorm:"many2many:product_categories;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProduct(product product.Product) *Product {
	return &Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) CreateNewProduct(product product.Product) error {
	var categories []category.Category

	if err := pr.db.Where("id IN ?", product.CategoryId).Find(&categories).Error; err != nil {
		return err
	}

	if err := pr.db.Where("name = ?", product.Name).First(&Product{}).Error; err == nil {
		return errors.New("the product name already exist")
	}

	if len(categories) != len(product.CategoryId) {
		return errors.New("category not found")
	}

	productData := NewProduct(product)
	productData.Categories = categories

	if err := pr.db.Create(productData).Error; err != nil {
		return err
	}

	return nil
}
