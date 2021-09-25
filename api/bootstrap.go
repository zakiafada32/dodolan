package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/zakiafada32/retail/api/utils"
	"github.com/zakiafada32/retail/api/v1/category"
	"github.com/zakiafada32/retail/api/v1/courier"
	"github.com/zakiafada32/retail/api/v1/payment"
	"github.com/zakiafada32/retail/api/v1/product"
	"github.com/zakiafada32/retail/api/v1/user"
)

func Bootstrap(
	e *echo.Echo,
	userController *user.UserController,
	categoryController *category.CategoryController,
	productController *product.ProductController,
	courierController *courier.CourierController,
	paymentController *payment.PaymentController,
) {
	if userController == nil {
		panic("user controller cannot be nil")
	}

	if categoryController == nil {
		panic("category controller cannot be nil")
	}

	if productController == nil {
		panic("product controller cannot be nil")
	}

	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	userV1 := e.Group("api/v1/users")
	userV1.POST("", userController.CreateNewUser)
	userV1.POST("/login", userController.Login)

	productV1 := e.Group("api/v1/products")
	productV1.POST("", productController.CreateNewProduct)

	categoryV1 := e.Group("api/v1/categories")
	categoryV1.POST("", categoryController.CreateNewCategory)

	paymentV1 := e.Group("api/v1/payment-providers")
	paymentV1.POST("", paymentController.CreateNewPaymentProvider)

	courierV1 := e.Group("api/v1/courier-providers")
	courierV1.POST("", courierController.CreateNewCourierProvider)

}
