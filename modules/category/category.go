package category

import (
	"errors"
	"time"

	"github.com/zakiafada32/retail/business/category"
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

func (cr *CategoryRepository) CreateNewCategory(category category.Category) error {
	if err := cr.db.Where("name = ?", category.Name).First(&Category{}).Error; err == nil {
		return errors.New("the category name already exist")
	}

	categoryData := convertToModulesCategory(category)
	if err := cr.db.Create(&categoryData).Error; err != nil {
		return err
	}

	return nil
}

func convertToModulesCategory(category category.Category) Category {
	return Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}
