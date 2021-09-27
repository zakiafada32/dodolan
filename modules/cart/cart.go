package cart

import (
	"errors"
	"time"

	"github.com/zakiafada32/retail/modules/category"
	"github.com/zakiafada32/retail/modules/user"
	"gorm.io/gorm"
)

type CartItem struct {
	UserID    string
	User      user.User
	ProductID uint32
	Product   category.Product
	Quantity  uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (repo *CartRepository) UpdateCartItem(userId string, productId, quantity uint32) error {

	var product category.Product
	err := repo.db.Where("id = ?", productId).First(&product).Error
	if err != nil {
		return err
	}

	if product.Stock < quantity {
		return errors.New("quantity is more than stock of the product")
	}

	var cartItem CartItem
	repo.db.Where("user_id = ? AND product_id = ?", userId, productId).Limit(1).Find(&cartItem)
	if cartItem.UserID == userId {
		err = repo.db.Model(&cartItem).Where("user_id = ? AND product_id = ?", userId, productId).Updates(CartItem{Quantity: quantity}).Error
		if err != nil {
			return err
		}
		return nil
	}

	err = repo.db.Create(&CartItem{UserID: userId, ProductID: productId, Quantity: quantity}).Error
	if err != nil {
		return err
	}

	return nil
}
