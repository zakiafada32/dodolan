package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/api/middlewares"
	"github.com/zakiafada32/retail/api/utils"
	"github.com/zakiafada32/retail/api/v1/cart"
	"github.com/zakiafada32/retail/api/v1/category"
	"github.com/zakiafada32/retail/api/v1/courier"
	"github.com/zakiafada32/retail/api/v1/order"
	"github.com/zakiafada32/retail/api/v1/payment"
	"github.com/zakiafada32/retail/api/v1/product"
	"github.com/zakiafada32/retail/api/v1/user"
)

type Controller struct {
	User     *user.UserController
	Category *category.CategoryController
	Product  *product.ProductController
	Payment  *payment.PaymentController
	Courier  *courier.CourierController
	Cart     *cart.CartController
	Order    *order.OrderController
}

func Bootstrap(e *echo.Echo, c *Controller) {
	if c.User == nil {
		panic("user controller cannot be nil")
	}

	if c.Category == nil {
		panic("category controller cannot be nil")
	}

	if c.Product == nil {
		panic("product controller cannot be nil")
	}

	if c.Payment == nil {
		panic("payment controller cannot be nil")
	}

	if c.Courier == nil {
		panic("courier controller cannot be nil")
	}

	if c.Order == nil {
		panic("order controller cannot be nil")
	}

	if c.Cart == nil {
		panic("cart controller cannot be nil")
	}

	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	userV1 := e.Group("api/v1/users")
	userV1.GET("", c.User.GetCurrentUser, middlewares.Authorized)
	userV1.POST("", c.User.CreateNewUser)
	userV1.POST("/login", c.User.Login)
	userV1.PUT("/", c.User.UpdateUser, middlewares.Authorized)

	categoryV1 := e.Group("api/v1/categories")
	categoryV1.GET("", c.Category.FindAllCategory)
	categoryV1.GET("/:id", c.Category.FindCategoryById)
	categoryV1.POST("", c.Category.CreateNewCategory, middlewares.Authorized, middlewares.IsAdmin)
	categoryV1.PUT("/:id", c.Category.UpdateCategory, middlewares.Authorized, middlewares.IsAdmin)

	productV1 := e.Group("api/v1/products")
	productV1.GET("", c.Product.FindAll)
	productV1.GET("/:id", c.Product.FindById)
	productV1.GET("/categories/:id", c.Product.FindByCategory)
	productV1.POST("", c.Product.CreateNew, middlewares.Authorized, middlewares.IsAdmin)
	productV1.PUT("/:id", c.Product.Update, middlewares.Authorized, middlewares.IsAdmin)

	paymentV1 := e.Group("api/v1/payment-providers")
	paymentV1.GET("", c.Payment.FindAll)
	paymentV1.POST("", c.Payment.CreateNew, middlewares.Authorized, middlewares.IsAdmin)
	paymentV1.PUT("/:id", c.Payment.Update, middlewares.Authorized, middlewares.IsAdmin)

	courierV1 := e.Group("api/v1/courier-providers")
	courierV1.GET("", c.Courier.FindAll)
	courierV1.POST("", c.Courier.CreateNew, middlewares.Authorized, middlewares.IsAdmin)
	courierV1.PUT("/:id", c.Courier.Update, middlewares.Authorized, middlewares.IsAdmin)

	cartV1 := e.Group("api/v1/cart")
	cartV1.GET("", c.Cart.FindAll, middlewares.Authorized)
	cartV1.POST("", c.Cart.AddCartItem, middlewares.Authorized)
	cartV1.POST("/delete", c.Cart.DeleteCartItem, middlewares.Authorized)
	cartV1.POST("/checkout", c.Cart.Checkout, middlewares.Authorized)

	orderV1 := e.Group("api/v1/orders")
	orderV1.GET("", c.Order.FindAll, middlewares.Authorized)
	orderV1.GET("/:id", c.Order.FindById, middlewares.Authorized)
	orderV1.POST("/payments", c.Order.Payment, middlewares.Authorized)
	orderV1.POST("/couriers", c.Order.Courier, middlewares.Authorized)
}
