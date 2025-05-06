package routes

import (
	"github.com/labstack/echo/v4"
	"lab4/controllers"
)

const ProductPrefix = "/product"

func ProductRouter(productController *controllers.ProductController, router *echo.Echo) {
	router.POST(ProductPrefix, productController.AddProduct)

	router.GET(ProductPrefix, productController.GetAllProducts)
	router.GET("/product/:id", productController.GetOneProduct)

	router.PUT(ProductPrefix, productController.EditProduct)

	router.DELETE("/product/:id", productController.RemoveProduct)
}
