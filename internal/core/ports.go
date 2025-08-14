package core

import "restaurant-go-api/internal/core/domain"

// interfaces y contratos
// Secondary port - repository -> Accede a los datos y manipula la persistencia de estos

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
	FindByCategory(category string) ([]domain.MenuItem, error) // Metodo especifico al menu
}

type CustomerRepository interface {
	Repository[domain.Customer]
}

type OrderRepository interface {
	Repository[domain.Order]
	FindByCustomerID(customerID string) ([]domain.Order, error)
}
