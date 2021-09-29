package product_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/product"
	"github.com/zakiafada32/retail/business/product/mocks"
)

const (
	id          uint32 = 1
	name        string = "product"
	description string = "description"
	price       uint64 = 1000
	stock       uint32 = 10
	categoryId  uint32 = 1
)

var (
	productService    product.Service
	productRepository mocks.Repository
	productData       product.Product
	productDataRepo   product.ProductAtt
	productsRepo      []product.ProductAtt
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindById(t *testing.T) {
	t.Run("Expect found the product with the input id", func(t *testing.T) {
		productRepository.On("FindById", id).Return(productDataRepo, nil).Once()
		product, err := productService.FindById(id)
		assert.Nil(t, err)
		assert.Equal(t, id, product.ID)
	})

	t.Run("Expect err when the product id not found", func(t *testing.T) {
		productRepository.On("FindById", id).Return(product.ProductAtt{}, errors.New("error")).Once()
		_, err := productService.FindById(id)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.NotFound)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Expect found all product", func(t *testing.T) {
		productRepository.On("FindAll").Return(productsRepo, nil).Once()
		products, err := productService.FindAll()
		assert.Nil(t, err)
		assert.IsType(t, []product.ProductAtt{}, products)
	})

	t.Run("Expect internal server error when cannot fetch products from database", func(t *testing.T) {
		productRepository.On("FindAll").Return([]product.ProductAtt{}, errors.New("error")).Once()
		_, err := productService.FindAll()
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func TestCreateNew(t *testing.T) {
	t.Run("Expect create new product success", func(t *testing.T) {
		productRepository.On("CreateNew", productData).Return(nil).Once()
		err := productService.CreateNew(productData)
		assert.Nil(t, err)
	})

	t.Run("Expect error when new product name already exist", func(t *testing.T) {
		productRepository.On("CreateNew", productData).Return(errors.New("error")).Once()
		err := productService.CreateNew(productData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Expect update the product success", func(t *testing.T) {
		productRepository.On("Update", id, productData).Return(productDataRepo, nil).Once()
		product, err := productService.Update(id, productData)
		assert.Nil(t, err)
		assert.Equal(t, productData.Name, product.Name)
		assert.Equal(t, productData.Stock, product.Stock)
		assert.Equal(t, productData.Price, product.Price)
	})

	t.Run("Expect error when the product id not found or name already exist", func(t *testing.T) {
		productRepository.On("Update", id, productData).Return(product.ProductAtt{}, errors.New("error")).Once()
		_, err := productService.Update(id, productData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func TestFindByCategory(t *testing.T) {
	t.Run("Expect find product by category", func(t *testing.T) {
		productRepository.On("FindByCategory", categoryId).Return(productsRepo, nil).Once()
		products, err := productService.FindByCategory(categoryId)
		assert.Nil(t, err)
		assert.IsType(t, []product.ProductAtt{}, products)
	})

	t.Run("Expect err when category id not found", func(t *testing.T) {
		productRepository.On("FindByCategory", categoryId).Return([]product.ProductAtt{}, errors.New("error")).Once()
		_, err := productService.FindByCategory(categoryId)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func setup() {
	productData = product.Product{
		ID:           id,
		Name:         name,
		Description:  description,
		Price:        price,
		Stock:        stock,
		CategoriesId: []uint32{1, 2},
	}

	productDataRepo = product.ProductAtt{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	productsRepo = []product.ProductAtt{
		{
			ID:          1,
			Name:        "product 1",
			Description: "description 1",
			Price:       100,
			Stock:       10,
		},
		{
			ID:          2,
			Name:        "product 2",
			Description: "description 2",
			Price:       200,
			Stock:       20,
		},
	}

	productService = product.NewProductService(&productRepository)
}
