package category

import (
	"errors"
	"time"

	categoryBusiness "github.com/zakiafada32/retail/business/category"
	"gorm.io/gorm"
)

type Category struct {
	ID          uint32 `gorm:"primaryKey"`
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (repo *CategoryRepository) CreateNewCategory(category categoryBusiness.Category) error {
	if err := repo.db.Where("name = ?", category.Name).First(&Category{}).Error; err == nil {
		return errors.New("the category name already exist")
	}

	categoryData := convertToCategoryModel(category)
	if err := repo.db.Create(&categoryData).Error; err != nil {
		return err
	}

	return nil
}

func (repo *CategoryRepository) FindAllCategory() ([]categoryBusiness.Category, error) {
	var categories []Category
	if err := repo.db.Find(&categories).Error; err != nil {
		return []categoryBusiness.Category{}, err
	}

	categoriesData := make([]categoryBusiness.Category, len(categories))
	for i, category := range categories {
		categoriesData[i] = convertToCategoryBusiness(category)
	}

	return categoriesData, nil

}

func (repo *CategoryRepository) FindCategoryById(id uint32) (categoryBusiness.Category, error) {
	var category Category
	if err := repo.db.Where("id = ?", id).First(&category).Error; err != nil {
		return categoryBusiness.Category{}, err
	}

	categoryData := convertToCategoryBusiness(category)

	return categoryData, nil

}

func (repo *CategoryRepository) UpdateCategory(categoryId uint32, name string, description string) (categoryBusiness.Category, error) {
	var category Category
	if err := repo.db.Where("name = ?", name).First(&category).Error; err == nil {
		return categoryBusiness.Category{}, errors.New("the category name already exist")
	}

	if err := repo.db.Where("id = ?", categoryId).First(&category).Error; err != nil {
		return categoryBusiness.Category{}, err
	}

	if err := repo.db.Model(&category).Updates(&Category{Name: name, Description: description}).Error; err != nil {
		return categoryBusiness.Category{}, err
	}

	categoryData := convertToCategoryBusiness(category)
	return categoryData, nil
}

func convertToCategoryModel(category categoryBusiness.Category) Category {
	return Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}

func convertToCategoryBusiness(category Category) categoryBusiness.Category {
	return categoryBusiness.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}
