package controllers

import (
	"github.com/labstack/echo/v4"
	"lab4/models"
	"lab4/service"
	"net/http"
	"strconv"
)

type CartController struct {
	CartService *service.CartService
}

func (cc *CartController) CreateCart(c echo.Context) error {
	var cart models.Cart
	if err := c.Bind(&cart); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	newCart, err := cc.CartService.CreateCart(cart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, newCart)
}

func (cc *CartController) GetCartByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid cart ID"})
	}

	cart, err := cc.CartService.GetCartByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Cart not found"})
	}

	return c.JSON(http.StatusOK, cart)
}

func (cc *CartController) DeleteCart(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid cart ID"})
	}

	err = cc.CartService.DeleteCartByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Cart deleted successfully"})
}

func (cc *CartController) AddProductToCart(c echo.Context) error {
	cartID, err1 := strconv.Atoi(c.Param("id"))
	productID, err2 := strconv.Atoi(c.Param("productID"))
	if err1 != nil || err2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid cart or product ID"})
	}

	err := cc.CartService.AddProductToCart(uint(cartID), uint(productID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Product added to cart"})
}

func (cc *CartController) RemoveProductFromCart(c echo.Context) error {
	cartID, err1 := strconv.Atoi(c.Param("id"))
	productID, err2 := strconv.Atoi(c.Param("productID"))
	if err1 != nil || err2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid cart or product ID"})
	}

	err := cc.CartService.RemoveProductFromCart(uint(cartID), uint(productID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Product removed from cart"})
}
