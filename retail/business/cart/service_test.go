package cart_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/cart"
	"github.com/zakiafada32/retail/business/cart/mocks"
)

const (
	userId    string = "1"
	productId uint32 = 1
	quantity  uint32 = 10
	paymentId uint32 = 1
	courierId uint32 = 1
)

var (
	cartService    cart.Service
	cartRepository mocks.Repository
	cartItemData   cart.CartItem
	cartItemsRepo  []cart.CartItemAtt
	cartData       cart.Cart
	productsId     []uint32 = []uint32{1, 2}
)

func TestFindAll(t *testing.T) {
	t.Run("Expect found all cart item", func(t *testing.T) {
		cartRepository.On("FindAll", userId).Return(cartItemsRepo, nil).Once()
		cartData, err := cartService.FindAll(userId)
		assert.Nil(t, err)
		assert.IsType(t, cart.Cart{}, cartData)
	})

	t.Run("Expect internal server error when cannot fetch cart items from database", func(t *testing.T) {
		cartRepository.On("FindAll", userId).Return([]cart.CartItemAtt{}, errors.New("error")).Once()
		_, err := cartService.FindAll(userId)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.InternalServerError)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Expect update cart item success", func(t *testing.T) {
		cartRepository.On("Update", userId, productId, quantity).Return(nil).Once()
		err := cartService.Update(userId, cartItemData)
		assert.Nil(t, err)
	})

	t.Run("Expect bad request when product id not found or quantity is more than stock", func(t *testing.T) {
		cartRepository.On("Update", userId, productId, quantity).Return(errors.New("error")).Once()
		err := cartService.Update(userId, cartItemData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func TestDeleteCartItem(t *testing.T) {
	t.Run("Expect delete cart item success", func(t *testing.T) {
		cartRepository.On("DeleteCartItem", userId, productsId).Return(nil).Once()
		err := cartService.DeleteCartItem(userId, productsId)
		assert.Nil(t, err)
	})

	t.Run("Expect bad request when product id on cart not found", func(t *testing.T) {
		cartRepository.On("DeleteCartItem", userId, productsId).Return(errors.New("error")).Once()
		err := cartService.DeleteCartItem(userId, productsId)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func TestCheckout(t *testing.T) {
	t.Run("Expect checkout cart item success", func(t *testing.T) {
		cartRepository.On("FindAll", userId).Return(cartItemsRepo, nil).Once()
		cartRepository.On("Checkout", userId, paymentId, courierId, cartData).Return(nil).Once()
		cartRepository.On("DeleteCartItem", userId, productsId).Return(nil).Once()
		err := cartService.Checkout(userId, paymentId, courierId)
		assert.Nil(t, err)
	})

	t.Run("Expect bad request when cannot fetch cart item", func(t *testing.T) {
		cartRepository.On("FindAll", userId).Return([]cart.CartItemAtt{}, errors.New("error")).Once()
		err := cartService.Checkout(userId, paymentId, courierId)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.InternalServerError)
	})

	t.Run("Expect bad request when cannot fetch cart item", func(t *testing.T) {
		cartRepository.On("FindAll", userId).Return(cartItemsRepo, nil).Once()
		cartRepository.On("Checkout", userId, paymentId, courierId, cartData).Return(errors.New("error")).Once()
		err := cartService.Checkout(userId, paymentId, courierId)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})

	t.Run("Expect bad request when cannot fetch cart item", func(t *testing.T) {
		cartRepository.On("FindAll", userId).Return(cartItemsRepo, nil).Once()
		cartRepository.On("Checkout", userId, paymentId, courierId, cartData).Return(errors.New("error")).Once()
		cartRepository.On("DeleteCartItem", userId, productsId).Return(errors.New("error")).Once()
		err := cartService.Checkout(userId, paymentId, courierId)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {

	cartItemData = cart.CartItem{
		ProductID:   productId,
		Quantity:    quantity,
		TotalAmount: 10,
	}

	cartItemsRepo = []cart.CartItemAtt{
		{
			Product: cart.CartProduct{
				ID:          1,
				Name:        "product 1",
				Description: "desc 1",
				Price:       10,
			},
			Quantity:    1,
			TotalAmount: 10,
		},
		{
			Product: cart.CartProduct{
				ID:          2,
				Name:        "product 2",
				Description: "desc 2",
				Price:       20,
			},
			Quantity:    2,
			TotalAmount: 40,
		},
	}

	cartData = cart.Cart{
		Items:       cartItemsRepo,
		TotalAmount: 50,
	}

	cartService = cart.NewCartService(&cartRepository)
}
