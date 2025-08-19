package services

import (
	"errors"
	"restaurant-go-api/internal/core"
	"restaurant-go-api/internal/core/domain"
	"restaurant-go-api/internal/shared/utils"
)

// logica de negocio

type OrderService interface {
	CreateCustomerOrder(customerID string) (*domain.Order, error)

	AddItemToCustomerOrder(orderID string, itemID string, quantity int) (*domain.Order, error)
	RemoveItemFromOrder(orderID string, itemID string) (*domain.Order, error)

	UpdateOrderStatus(orderID string, status domain.OrderStatus) (*domain.Order, error)
	CalculateOrderTotal(orderID string) (int, error)
	GetOrderDetails(orderID string) (*domain.Order, error)
	ProcessPayment(orderID string, paymentMethod string) error
}

type orderService struct {
	orderRepo core.OrderRepository
	userRepo  core.UserRepository
	menuRepo  core.MenuRepository
}

// Constructor para una nueva orden
func NewOrderService(
	orderRepo core.OrderRepository,
	userRepo core.UserRepository,
	menuRepo core.MenuRepository,
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
func (s *orderService) AddItemToCustomerOrder(orderID string, itemID string, quantity int) (*domain.Order, error) {
	order, orderErr := s.orderRepo.FindByID(orderID)
	if orderErr != nil {
		return nil, orderErr
	}
	menuItem, menuItemErr := s.menuRepo.FindByID(itemID)
	if menuItemErr != nil {
		return nil, menuItemErr
	}

	for range quantity {
		order.OrderItems = append(order.OrderItems, *menuItem)
	}

	order.OrderTotal = order.CalculateSubtotal()
	err := s.orderRepo.Update(order)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *orderService) RemoveItemFromOrder(orderID string, itemID string) (*domain.Order, error) {
	order, orderErr := s.orderRepo.FindByID(orderID)
	if orderErr != nil {
		return nil, orderErr
	}

	itemIndex := -1
	for i, orderItem := range order.OrderItems {
		if orderItem.ID == itemID {
			itemIndex = i
			break
		}
	}

	if itemIndex == -1 {
		return nil, errors.New("item not found in customer order")
	}
	order.OrderItems = append(order.OrderItems, order.OrderItems[:itemIndex+1]...)
	order.OrderTotal = order.CalculateSubtotal()

	updateError := s.orderRepo.Update(order)
	if updateError != nil {
		return nil, updateError
	}

	return order, nil
}

func (s *orderService) UpdateOrderStatus(orderID string, status domain.OrderStatus) (*domain.Order, error) {
	order, orderErr := s.orderRepo.FindByID(orderID)
	if orderErr != nil {
		return nil, orderErr
	}

	order.OrderStatus = status

	updateOrderErr := s.orderRepo.Update(order)
	if updateOrderErr != nil {
		return nil, updateOrderErr
	}
	return order, nil
}

func (s *orderService) CalculateOrderTotal(orderID string) (int, error) {
	// encuentra la oden
	order, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		return 0, err
	}
	// utiliza el metodo creado en el "domain"
	subTotal := order.CalculateSubtotal()
	order.OrderTotal = subTotal

	// Guarda los datos en el repositorio
	err = s.orderRepo.Update(order)
	if err != nil {
		return 0, err
	}

	return subTotal, nil
}

func (s *orderService) GetOrderDetails(orderID string) (*domain.Order, error) {
	order, orderErr := s.orderRepo.FindByID(orderID)
	if orderErr != nil {
		return nil, orderErr
	}
	return order, nil
}

func (s *orderService) ProcessPayment(orderID string, paymentMethod string) error {
	// TODO: Implement later
	return errors.New("not implemented yet")
}
