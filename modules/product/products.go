package product

import (
	"errors"
	"time"

	productBusiness "github.com/zakiafada32/retail/business/product"
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

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) CreateNewProduct(product productBusiness.Product) error {
	var categories []category.Category

	err := repo.db.Where("id IN ?", product.CategoryId).Find(&categories).Error
	if err != nil {
		return err
	}

	err = repo.db.Where("name = ?", product.Name).First(&Product{}).Error
	if err == nil {
		return errors.New("the product name already exist")
	}

	if len(categories) != len(product.CategoryId) {
		return errors.New("category not found")
	}

	productData := convertToProductModel(product)
	productData.Categories = categories

	err = repo.db.Create(&productData).Error
	if err != nil {
		return err
	}

	return nil
}

func convertToProductModel(product productBusiness.Product) Product {
	return Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
