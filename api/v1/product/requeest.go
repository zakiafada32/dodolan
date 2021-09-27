package product

import "github.com/zakiafada32/retail/business/product"

type productRequestBody struct {
	Name         string   `json:"name" validate:"required"`
	Description  string   `json:"description"`
	Stock        uint32   `json:"stock" validate:"required"`
	Price        uint64   `json:"price" validate:"required"`
	CategoriesId []uint32 `json:"category_id" validate:"gt=0,dive,required"`
}

func (req *productRequestBody) convertToProductBusiness() product.Product {
	return product.Product{
		Name:         req.Name,
		Description:  req.Description,
		Stock:        req.Stock,
		Price:        req.Price,
		CategoriesId: req.CategoriesId,
	}
}
