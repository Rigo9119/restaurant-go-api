package memory

import (
	"errors"
	"restaurant-go-api/internal/core"
	"restaurant-go-api/internal/core/domain"
)

type InMemoryOrderRepository struct {
	orders map[string]*domain.Order
}

func NewOrderRepository() core.OrderRepository {
	return &InMemoryOrderRepository{
		orders: make(map[string]*domain.Order),
	}
}

func (r *InMemoryOrderRepository) Save(order *domain.Order) error {
	r.orders[order.ID] = order
	return nil
}

func (r *InMemoryOrderRepository) FindByID(id string) (*domain.Order, error) {
	order, exists := r.orders[id]
	if !exists {
		return nil, errors.New("order not found")
	}
	return order, nil
}

func (r *InMemoryOrderRepository) GetAll() ([]domain.Order, error) {
	orders := make([]domain.Order, 0, len(r.orders))
	for _, order := range r.orders {
		orders = append(orders, *order)
	}
	return orders, nil
}

func (r *InMemoryOrderRepository) Update(order *domain.Order) error {
	if _, exists := r.orders[order.ID]; !exists {
		return errors.New("order not found")
	}
	r.orders[order.ID] = order
	return nil
}

func (r *InMemoryOrderRepository) Delete(id string) error {
	if _, exists := r.orders[id]; !exists {
		return errors.New("order not found")
	}
	delete(r.orders, id)
	return nil
}

func (r *InMemoryOrderRepository) FindByCustomerID(customerID string) ([]domain.Order, error) {
	orders := make([]domain.Order, 0)
	for _, order := range r.orders {
		if order.UserID == customerID {
			orders = append(orders, *order)
		}
	}
	return orders, nil
}
