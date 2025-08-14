package core

import "restaurant-go-api/internal/core/domain"

// logica de negocio

type OrderService interface {
	CreateCustomerOrder(customerID string) (*domain.Order, error)

	AddItemToCustomerOrder(orderID string, itemID string, quantity int) error
	RemoveItemFromOrder(orderID string, itemID string) error

	UpdateOrderStatus(orderID string, status domain.OrderStatus) error
	CalculateOrderTotal(orderID string) (int, error)
	GetOrderDetails(orderID string) (*domain.Order, error)
	ProcessPayment(orderID string, paymentMethod string) error
}
