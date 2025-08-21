// Package core
// secondary port / repo -> acces data and controls data persistency
// access data and controls data persistency
// (interfaces y contratos)
package core

import "restaurant-go-api/internal/core/domain"

// Repository generico - T puede ser any
type Repository[T any] interface {
	Save(entity *T) error
	FindByID(id string) (*T, error)
	GetAll() ([]T, error)
	Update(entity *T) error
	Delete(id string) error
}

// MenuRepository accede al repo generico y añade otros metodos
type MenuRepository interface {
	Repository[domain.MenuItem]                                // Operaciones CRUD genericas
	FindByCategory(category domain.Category) ([]domain.MenuItem, error) // Metodo especifico al menu
}

type UserRepository interface {
	Repository[domain.User]
}

type OrderRepository interface {
	Repository[domain.Order]
	FindByCustomerID(customerID string) ([]domain.Order, error)
}
