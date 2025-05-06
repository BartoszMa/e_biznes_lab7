package routes

import (
	"github.com/labstack/echo/v4"
	"lab4/controllers"
)

func ProductRouter(productController *controllers.ProductController, router *echo.Echo) {
	router.POST("/product", productController.AddProduct)

	router.GET("/product", productController.GetAllProducts)
	router.GET("/product/:id", productController.GetOneProduct)

	router.PUT("/product", productController.EditProduct)

	router.DELETE("/product/:id", productController.RemoveProduct)
}
