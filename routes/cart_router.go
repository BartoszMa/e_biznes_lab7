package routes

import (
	"github.com/labstack/echo/v4"
	"lab4/controllers"
)

func CartRouter(cartController *controllers.CartController, router *echo.Echo) {
	router.POST("/cart", cartController.CreateCart)
	router.GET("/cart/:id", cartController.GetCartByID)
	router.DELETE("/cart/:id", cartController.DeleteCart)
	router.POST("/cart/:id/add/:productID", cartController.AddProductToCart)
	router.DELETE("/cart/:id/remove/:productID", cartController.RemoveProductFromCart)
}
