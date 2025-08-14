package core

import (
	"errors"
	"restaurant-go-api/internal/core/domain"
	"restaurant-go-api/internal/shared/utils"
)

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

type orderService struct {
	orderRepo OrderRepository
	userRepo  UserRepository
	menuRepo  MenuRepository
}

// Constructor para una nueva orden
func NewOrderService(
	orderRepo OrderRepository,
	userRepo UserRepository,
	menuRepo MenuRepository,
) OrderService {
	return &orderService{
		orderRepo: orderRepo,
		userRepo:  userRepo,
		menuRepo:  menuRepo,
	}
}

func (s *orderService) CreateCustomerOrder(userID string) (*domain.Order, error) {
	// 1. Valida que tengamos un usuariok
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// 2. Crea una nueva orden
	order := &domain.Order{
		ID:          utils.GenerateRandomID("order"),
		UserID:      userID,
		OrderItems:  []domain.MenuItem{},
		OrderTotal:  0,
		OrderStatus: domain.Ordered,
	}

	// 3. Guarda la orden
	err = s.orderRepo.Save(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// Implement remaining methods from OrderService interface
func (s *orderService) AddItemToCustomerOrder(orderID string, itemID string, quantity int) error {
	// TODO: Implement later
	return errors.New("not implemented yet")
}

func (s *orderService) RemoveItemFromOrder(orderID string, itemID string) error {
	// TODO: Implement later
	return errors.New("not implemented yet")
}

func (s *orderService) UpdateOrderStatus(orderID string, status domain.OrderStatus) error {
	// TODO: Implement later
	return errors.New("not implemented yet")
}

func (s *orderService) CalculateOrderTotal(orderID string) (int, error) {
	// TODO: Implement later
	return 0, errors.New("not implemented yet")
}

func (s *orderService) GetOrderDetails(orderID string) (*domain.Order, error) {
	// TODO: Implement later
	return nil, errors.New("not implemented yet")
}

func (s *orderService) ProcessPayment(orderID string, paymentMethod string) error {
	// TODO: Implement later
	return errors.New("not implemented yet")
}

// ID generation is now handled by utils.GenerateRandomID()
