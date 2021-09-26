package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/api/middlewares"
	"github.com/zakiafada32/retail/api/utils"
	"github.com/zakiafada32/retail/api/v1/category"
	"github.com/zakiafada32/retail/api/v1/courier"
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
		panic("product controller cannot be nil")
	}

	if c.Courier == nil {
		panic("product controller cannot be nil")
	}

	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	userV1 := e.Group("api/v1/users")
	userV1.GET("", c.User.GetCurrentUser, middlewares.Authorized())
	userV1.POST("", c.User.CreateNewUser)
	userV1.POST("/login", c.User.Login)
	userV1.PUT("/", c.User.UpdateUser, middlewares.Authorized())

	categoryV1 := e.Group("api/v1/categories")
	categoryV1.GET("", c.Category.FindAllCategory)
	categoryV1.GET("/:id", c.Category.FindCategoryById)
	categoryV1.POST("", c.Category.CreateNewCategory, middlewares.Authorized(), middlewares.IsAdmin)
	categoryV1.PUT("/:id", c.Category.UpdateCategory, middlewares.Authorized(), middlewares.IsAdmin)

	productV1 := e.Group("api/v1/products")
	productV1.GET("", func(c echo.Context) error { return nil })
	productV1.GET("/:id", func(c echo.Context) error { return nil })
	productV1.GET("/categories/:id", func(c echo.Context) error { return nil })
	productV1.POST("", c.Product.CreateNewProduct, middlewares.Authorized(), middlewares.IsAdmin)
	productV1.PUT("/:id", func(c echo.Context) error { return nil }, middlewares.Authorized(), middlewares.IsAdmin)

	paymentV1 := e.Group("api/v1/payment-providers")
	paymentV1.GET("", func(c echo.Context) error { return nil })
	paymentV1.POST("", c.Payment.CreateNewPaymentProvider)
	paymentV1.PUT("/:id", func(c echo.Context) error { return nil }, middlewares.Authorized(), middlewares.IsAdmin)

	courierV1 := e.Group("api/v1/courier-providers")
	courierV1.GET("", func(c echo.Context) error { return nil })
	courierV1.POST("", c.Courier.CreateNewCourierProvider)
	courierV1.PUT("/:id", func(c echo.Context) error { return nil }, middlewares.Authorized(), middlewares.IsAdmin)

	cartV1 := e.Group("api/v1/cart")
	cartV1.GET("", func(c echo.Context) error { return nil }, middlewares.Authorized())
	cartV1.POST("", func(c echo.Context) error { return nil }, middlewares.Authorized())
	cartV1.POST("/checkout", func(c echo.Context) error { return nil }, middlewares.Authorized())

	orderV1 := e.Group("api/v1/order")
	orderV1.GET("", func(c echo.Context) error { return nil }, middlewares.Authorized())
	orderV1.GET("/:id", func(c echo.Context) error { return nil }, middlewares.Authorized())
	orderV1.POST("", func(c echo.Context) error { return nil }, middlewares.Authorized())
	orderV1.POST("/payment", func(c echo.Context) error { return nil }, middlewares.Authorized())
	orderV1.POST("/courier", func(c echo.Context) error { return nil }, middlewares.Authorized())
}
