package controllers

import (
	"github.com/labstack/echo/v4"
	"lab4/models"
	"lab4/service"
	"net/http"
	"strconv"
)

type PaymentController struct {
	PaymentService *service.PaymentService
}

func (pc *PaymentController) CreatePayment(c echo.Context) error {
	var payment models.Payment
	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	newPayment, err := pc.PaymentService.CreatePayment(payment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, newPayment)
}

func (pc *PaymentController) GetPaymentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payment ID"})
	}

	payment, err := pc.PaymentService.GetPaymentByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Payment not found"})
	}

	return c.JSON(http.StatusOK, payment)
}

func (pc *PaymentController) DeletePayment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payment ID"})
	}

	err = pc.PaymentService.DeletePaymentByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Payment deleted successfully"})
}

func (pc *PaymentController) MarkPaymentAsPaid(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payment ID"})
	}

	err = pc.PaymentService.MarkAsPaid(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Payment marked as paid"})
}
