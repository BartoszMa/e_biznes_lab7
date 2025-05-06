package service

import (
	"gorm.io/gorm"
	"lab4/models"
)

type PaymentService struct {
	DB *gorm.DB
}

func (s *PaymentService) CreatePayment(payment models.Payment) (models.Payment, error) {
	err := s.DB.Create(&payment).Error
	if err != nil {
		return models.Payment{}, err
	}
	return payment, nil
}

func (s *PaymentService) GetPaymentByID(paymentID uint) (models.Payment, error) {
	var payment models.Payment
	err := s.DB.First(&payment, paymentID).Error
	return payment, err
}

func (s *PaymentService) DeletePaymentByID(paymentID uint) error {
	return s.DB.Delete(&models.Payment{}, paymentID).Error
}

func (s *PaymentService) MarkAsPaid(paymentID uint) error {
	return s.DB.Model(&models.Payment{}).
		Where("id = ?", paymentID).
		Update("is_payed", true).Error
}
