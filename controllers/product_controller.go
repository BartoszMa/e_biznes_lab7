package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"lab4/models"
	"lab4/service"
)

type ProductController struct {
	DbService *service.Service
}

func (pc *ProductController) GetAllProducts(c echo.Context) error {
	products, err := pc.DbService.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetOneProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	product, err := pc.DbService.GetOneProduct(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, product)
}

func (pc *ProductController) AddProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	err := pc.DbService.AddProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "Product added successfully"})
}

func (pc *ProductController) RemoveProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	err = pc.DbService.RemoveProduct(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Product removed successfully"})
}

func (pc *ProductController) EditProduct(c echo.Context) error {
	var editedProduct models.Product
	if err := c.Bind(&editedProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	err := pc.DbService.EditProduct(editedProduct)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Product updated successfully"})
}
