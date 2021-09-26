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

func (repo *CategoryRepository) CreateNew(category categoryBusiness.Category) error {
	err := repo.db.Where("name = ?", category.Name).First(&Category{}).Error
	if err == nil {
		return errors.New("the category name already exist")
	}

	categoryData := convertToCategoryModel(category)
	err = repo.db.Create(&categoryData).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *CategoryRepository) FindAll() ([]categoryBusiness.Category, error) {
	var categories []Category
	err := repo.db.Find(&categories).Error
	if err != nil {
		return []categoryBusiness.Category{}, err
	}

	categoriesData := make([]categoryBusiness.Category, len(categories))
	for i, category := range categories {
		categoriesData[i] = ConvertToCategoryBusiness(category)
	}

	return categoriesData, nil

}

func (repo *CategoryRepository) FindById(id uint32) (categoryBusiness.Category, error) {
	var category Category
	err := repo.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return categoryBusiness.Category{}, err
	}

	categoryData := ConvertToCategoryBusiness(category)

	return categoryData, nil

}

func (repo *CategoryRepository) Update(id uint32, name string, description string) (categoryBusiness.Category, error) {
	var category Category
	err := repo.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return categoryBusiness.Category{}, err
	}

	if len(name) > 0 && name != category.Name {
		err = repo.db.Where("name = ?", name).First(&Category{}).Error
		if err == nil {
			return categoryBusiness.Category{}, errors.New("the category name already exist")
		}
	}

	err = repo.db.Model(&category).Updates(&Category{Name: name, Description: description}).Error
	if err != nil {
		return categoryBusiness.Category{}, err
	}

	categoryData := ConvertToCategoryBusiness(category)
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

func ConvertToCategoryBusiness(category Category) categoryBusiness.Category {
	return categoryBusiness.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}
