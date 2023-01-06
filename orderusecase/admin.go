package orderusecase

import "github.com/rizlantamima/go-rest/domain"

type AdminOrderUseCase struct{}

func (c *AdminOrderUseCase) CreateOrder(userID int, order domain.Order) error {

	return nil
}

func (c *AdminOrderUseCase) CancelOrder(userID int, orderID int) error {

	return nil
}

func (c *AdminOrderUseCase) ViewOrder(userID int, orderID int) (*domain.Order, error) {

	return nil, nil
}
