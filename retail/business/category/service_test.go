package category_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/category"
	"github.com/zakiafada32/retail/business/category/mocks"
)

const (
	id          uint32 = 1
	name        string = "category"
	description string = "description"
)

var (
	categoryService    category.Service
	categoryRepository mocks.Repository
	categoryData       category.Category
	categorysData      []category.Category
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindAll(t *testing.T) {
	t.Run("Expect found all category", func(t *testing.T) {
		categoryRepository.On("FindAll").Return(categorysData, nil).Once()
		categorys, err := categoryService.FindAll()
		assert.Nil(t, err)
		assert.IsType(t, []category.Category{}, categorys)
	})

	t.Run("Expect internal server error when cannot fetch category from database", func(t *testing.T) {
		categoryRepository.On("FindAll").Return([]category.Category{}, errors.New("error")).Once()
		_, err := categoryService.FindAll()
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.InternalServerError)
	})
}

func TestFindById(t *testing.T) {
	t.Run("Expect find the category by the id", func(t *testing.T) {
		categoryRepository.On("FindById", id).Return(categoryData, nil).Once()
		category, err := categoryService.FindById(id)
		assert.Nil(t, err)
		assert.Equal(t, name, category.Name)
		assert.Equal(t, description, category.Description)
	})

	t.Run("Expect not found when cannot find category id", func(t *testing.T) {
		categoryRepository.On("FindById", id).Return(category.Category{}, errors.New("error")).Once()
		_, err := categoryService.FindById(id)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.NotFound)
	})
}

func TestCreateNew(t *testing.T) {
	t.Run("Expect create new category", func(t *testing.T) {
		categoryRepository.On("CreateNew", categoryData).Return(nil).Once()
		err := categoryService.CreateNew(categoryData)
		assert.Nil(t, err)
	})

	t.Run("Expect bad request error when category name already exist", func(t *testing.T) {
		categoryRepository.On("CreateNew", categoryData).Return(errors.New("error")).Once()
		err := categoryService.CreateNew(categoryData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Expect update the category by the id", func(t *testing.T) {
		categoryRepository.On("Update", id, name, description).Return(categoryData, nil).Once()
		category, err := categoryService.Update(id, name, description)
		assert.Nil(t, err)
		assert.Equal(t, name, category.Name)
		assert.Equal(t, description, category.Description)
	})

	t.Run("Expect not found when cannot find category id", func(t *testing.T) {
		categoryRepository.On("Update", id, name, description).Return(category.Category{}, errors.New("error")).Once()
		_, err := categoryService.Update(id, name, description)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func setup() {
	categoryData = category.Category{
		ID:          id,
		Name:        name,
		Description: description,
	}

	categorysData = []category.Category{
		{
			ID:          1,
			Name:        "category 1",
			Description: "description 1",
		},
		{
			ID:          2,
			Name:        "category 2",
			Description: "description 2",
		},
	}

	categoryService = category.NewCategoryService(&categoryRepository)
}
