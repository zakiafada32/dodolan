package product

import (
	"errors"
	"time"

	business "github.com/zakiafada32/retail/business/product"
	moduleCategory "github.com/zakiafada32/retail/modules/category"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

type Product struct {
	ID          uint32 `gorm:"primaryKey"`
	Name        string
	Description string
	Stock       *uint32
	Price       uint64
	Categories  []moduleCategory.Category `gorm:"many2many:product_categories;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) FindById(id uint32) (business.ProductAtt, error) {
	var product Product
	err := repo.db.Preload("Categories").Where("id = ?", id).First(&product).Error
	if err != nil {
		return business.ProductAtt{}, err
	}

	productData := convertToProductAttBusiness(product)
	for _, category := range product.Categories {
		productData.Categories = append(productData.Categories, moduleCategory.ConvertToCategoryBusiness(category))
	}

	return productData, nil
}

func (repo *ProductRepository) FindAll() ([]business.ProductAtt, error) {
	var products []Product
	err := repo.db.Preload("Categories").Find(&products).Error
	if err != nil {
		return []business.ProductAtt{}, err
	}

	productsData := make([]business.ProductAtt, len(products))
	for i, product := range products {
		productsData[i] = convertToProductAttBusiness(product)
		for _, category := range product.Categories {
			productsData[i].Categories = append(productsData[i].Categories, moduleCategory.ConvertToCategoryBusiness(category))
		}
	}

	return productsData, nil
}

func (repo *ProductRepository) CreateNew(product business.Product) error {
	var categories []moduleCategory.Category

	err := repo.db.Where("id IN ?", product.CategoriesId).Find(&categories).Error
	if err != nil {
		return err
	}

	err = repo.db.Where("name = ?", product.Name).First(&Product{}).Error
	if err == nil {
		return errors.New("the product name already exist")
	}

	if len(categories) != len(product.CategoriesId) {
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

func (repo *ProductRepository) Update(id uint32, updateData business.ProductUpdate) (business.ProductAtt, error) {
	var product Product
	err := repo.db.Preload("Categories").Where("id = ?", id).First(&product).Error
	if err != nil {
		return business.ProductAtt{}, err
	}

	if len(updateData.Name) > 0 && product.Name != updateData.Name {
		err := repo.db.Where("name = ?", updateData.Name).First(&product).Error
		if err == nil {
			return business.ProductAtt{}, errors.New("product name already exist")
		}
	}

	// todo: cannot update category
	// var categories []moduleCategory.Category
	// err = repo.db.Where("id IN ?", updateData.CategoriesId).Find(&categories).Error
	// if err != nil {
	// 	return business.ProductAtt{}, err
	// }
	// if len(categories) != len(updateData.CategoriesId) {
	// 	return business.ProductAtt{}, errors.New("category not found")
	// }

	err = repo.db.Model(&product).Updates(&Product{
		Name:        updateData.Name,
		Description: updateData.Description,
		Stock:       &updateData.Stock,
		Price:       updateData.Price,
	}).Error
	if err != nil {
		return business.ProductAtt{}, err
	}

	productData := convertToProductAttBusiness(product)
	for _, category := range product.Categories {
		productData.Categories = append(productData.Categories, moduleCategory.ConvertToCategoryBusiness(category))
	}
	return productData, nil

}

func convertToProductModel(product business.Product) Product {
	return Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Stock:       &product.Stock,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func convertToProductBusiness(product Product) business.Product {
	return business.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Stock:       *product.Stock,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func convertToProductAttBusiness(product Product) business.ProductAtt {
	return business.ProductAtt{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Stock:       *product.Stock,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
