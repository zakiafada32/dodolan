package modules

import (
	"github.com/zakiafada32/retail/modules/category"
	"github.com/zakiafada32/retail/modules/product"
	"github.com/zakiafada32/retail/modules/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &product.Product{}, &category.Category{})
}
