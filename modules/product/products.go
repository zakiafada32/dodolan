package product

import (
	"errors"

	business "github.com/zakiafada32/retail/business/product"
	module "github.com/zakiafada32/retail/modules/category"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) FindById(id uint32) (business.ProductAtt, error) {
	var product module.Product
	err := repo.db.Preload("Categories").Where("id = ?", id).First(&product).Error
	if err != nil {
		return business.ProductAtt{}, err
	}

	productData := convertToProductAttBusiness(product)
	for _, categoryData := range product.Categories {
		productData.Categories = append(productData.Categories, module.ConvertToCategoryBusiness(*categoryData))
	}

	return productData, nil
}

func (repo *ProductRepository) FindAll() ([]business.ProductAtt, error) {
	var products []module.Product
	err := repo.db.Preload("Categories").Find(&products).Error
	if err != nil {
		return []business.ProductAtt{}, err
	}

	productsData := make([]business.ProductAtt, len(products))
	for i, product := range products {
		productsData[i] = convertToProductAttBusiness(product)
		for _, category := range product.Categories {
			productsData[i].Categories = append(productsData[i].Categories, module.ConvertToCategoryBusiness(*category))
		}
	}

	return productsData, nil
}

func (repo *ProductRepository) CreateNew(product business.Product) error {
	var categories []*module.Category

	err := repo.db.Where("id IN ?", product.CategoriesId).Find(&categories).Error
	if err != nil {
		return err
	}

	err = repo.db.Where("name = ?", product.Name).First(&module.Product{}).Error
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

func (repo *ProductRepository) Update(id uint32, updateData business.Product) (business.ProductAtt, error) {
	var product module.Product
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

	if len(updateData.CategoriesId) > 0 {
		var categories []module.Category
		err = repo.db.Where("id IN ?", updateData.CategoriesId).Find(&categories).Error
		if err != nil {
			return business.ProductAtt{}, err
		}
		if len(categories) != len(updateData.CategoriesId) {
			return business.ProductAtt{}, errors.New("category not found")
		}
		err = repo.db.Model(&product).Association("Categories").Delete(product.Categories)
		if err != nil {
			return business.ProductAtt{}, err
		}
		err = repo.db.Model(&product).Association("Categories").Replace(categories)
		if err != nil {
			return business.ProductAtt{}, err
		}
	}

	err = repo.db.Model(&product).Updates(&module.Product{
		Name:        updateData.Name,
		Description: updateData.Description,
		Stock:       updateData.Stock,
		Price:       updateData.Price,
	}).Error
	if err != nil {
		return business.ProductAtt{}, err
	}

	productData := convertToProductAttBusiness(product)
	for _, category := range product.Categories {
		productData.Categories = append(productData.Categories, module.ConvertToCategoryBusiness(*category))
	}
	return productData, nil

}

func (repo *ProductRepository) FindByCategory(categoryId uint32) ([]business.ProductAtt, error) {
	var category module.Category
	err := repo.db.Where("id = ?", categoryId).First(&category).Error
	if err != nil {
		return []business.ProductAtt{}, err
	}

	var products []module.Product
	err = repo.db.Model(&category).Preload("Categories").Association("Products").Find(&products)
	if err != nil {
		return []business.ProductAtt{}, err
	}

	productsData := make([]business.ProductAtt, len(products))
	for i, product := range products {
		productsData[i] = convertToProductAttBusiness(product)
		for _, category := range product.Categories {
			productsData[i].Categories = append(productsData[i].Categories, module.ConvertToCategoryBusiness(*category))
		}
	}

	return productsData, nil
}

func convertToProductModel(product business.Product) module.Product {
	return module.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func convertToProductAttBusiness(product module.Product) business.ProductAtt {
	return business.ProductAtt{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
