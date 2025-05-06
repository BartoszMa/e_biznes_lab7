package service

import (
	"gorm.io/gorm"
	"lab4/models"
	"lab4/scopes"
)

type CartService struct {
	DB *gorm.DB
}

func (s *CartService) GetCartByID(cartID uint) (models.Cart, error) {
	var cart models.Cart
	err := s.DB.Scopes(scopes.CartWithItemsAndProducts()).First(&cart, cartID).Error
	return cart, err
}

func (s *CartService) CreateCart(cart models.Cart) (models.Cart, error) {
	err := s.DB.Create(&cart).Error
	if err != nil {
		return models.Cart{}, err
	}
	return cart, nil
}

func (s *CartService) DeleteCartByID(cartID uint) error {
	err := s.DB.Delete(&models.Cart{}, cartID).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *CartService) AddProductToCart(cartID uint, productID uint) error {
	newItem := models.CartItem{
		CartID:    cartID,
		ProductID: productID,
	}
	return s.DB.Create(&newItem).Error
}

func (s *CartService) RemoveProductFromCart(cartID uint, productID uint) error {
	return s.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&models.CartItem{}).Error
}
