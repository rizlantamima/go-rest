package orderusecase

import "github.com/rizlantamima/go-rest/domain"

type CustomerOrderUseCase struct{}

func (c *CustomerOrderUseCase) CreateOrder(userID int, order domain.Order) error {

	return nil
}

func (c *CustomerOrderUseCase) CancelOrder(userID int, orderID int) error {

	return nil
}

func (c *CustomerOrderUseCase) ViewOrder(userID int, orderID int) (*domain.Order, error) {

	return nil, nil
}
