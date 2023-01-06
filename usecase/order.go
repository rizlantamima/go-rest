package usecase

import "github.com/rizlantamima/go-rest/domain"

type OrderUseCase interface {
	CreateOrder(userID uint32, order domain.Order) error
	CancelOrder(userID uint32, orderID uint32) error
	ViewOrder(userID uint32, orderID uint32) (*domain.Order, error)
}
