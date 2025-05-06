package routes

import (
	"github.com/labstack/echo/v4"
	"lab4/controllers"
)

func PaymentRouter(paymentController *controllers.PaymentController, router *echo.Echo) {
	router.POST("/payments", paymentController.CreatePayment)
	router.GET("/payments/:id", paymentController.GetPaymentByID)
	router.DELETE("/payments/:id", paymentController.DeletePayment)
	router.PUT("/payments/:id/pay", paymentController.MarkPaymentAsPaid)
}
