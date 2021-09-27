package cart

import (
	"errors"
	"time"

	cartBusiness "github.com/zakiafada32/retail/business/cart"
	"github.com/zakiafada32/retail/modules/category"
	"github.com/zakiafada32/retail/modules/order"
	"github.com/zakiafada32/retail/modules/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (repo *CartRepository) Update(userId string, productId, quantity uint32) error {

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

func (repo *CartRepository) FindAll(userId string) ([]cartBusiness.CartItemAtt, error) {
	var cartItem []CartItem
	err := repo.db.Preload(clause.Associations).Where("user_id = ?", userId).Find(&cartItem).Error
	if err != nil {
		return []cartBusiness.CartItemAtt{}, err
	}

	items := make([]cartBusiness.CartItemAtt, len(cartItem))
	for i, item := range cartItem {
		items[i] = convertToCartAttBusiness(item)
	}
	return items, nil
}

func (repo *CartRepository) DeleteCartItem(userId string, productsId []uint32) error {
	err := repo.db.Where("product_id IN ? AND user_id = ?", productsId, userId).Delete(&CartItem{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *CartRepository) Checkout(userId string, paymentId uint32, courierId uint32, cart cartBusiness.Cart) error {
	yet := false
	orderData := order.Order{
		UserID:            userId,
		PaymentProviderID: paymentId,
		PaymentStatus:     &yet,
		CourierProviderID: courierId,
		CourierStatus:     &yet,
		TotalAmount:       cart.TotalAmount,
	}

	err := repo.db.Create(&orderData).Error
	if err != nil {
		return err
	}

	orderItem := make([]order.OrderItem, len(cart.Items))
	for i, item := range cart.Items {
		err = updateProduct(repo.db, item.Product.ID, item.Quantity)
		if err != nil {
			return err
		}
		orderItem[i].OrderID = orderData.ID
		orderItem[i].ProductID = item.Product.ID
		orderItem[i].Quantity = item.Quantity
		orderItem[i].TotalAmount = item.TotalAmount
	}

	err = repo.db.Create(&orderItem).Error
	if err != nil {
		return err
	}

	return nil
}

func convertToCartAttBusiness(cartItem CartItem) cartBusiness.CartItemAtt {
	return cartBusiness.CartItemAtt{
		Quantity:    cartItem.Quantity,
		TotalAmount: uint64(cartItem.Quantity) * cartItem.Product.Price,
		Product: cartBusiness.CartProduct{
			ID:          cartItem.Product.ID,
			Name:        cartItem.Product.Name,
			Description: cartItem.Product.Description,
			Price:       cartItem.Product.Price,
		},
	}
}

func updateProduct(db *gorm.DB, id uint32, quantity uint32) error {
	var product category.Product
	err := db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return err
	}

	if product.Stock < quantity {
		return errors.New("stock of the product not enough")
	}

	product.Stock -= quantity

	err = db.Model(&product).Updates(&category.Product{Stock: product.Stock}).Error
	if err != nil {
		return err
	}

	return nil

}
