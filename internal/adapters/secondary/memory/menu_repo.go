// Package memory ayuda a guardar los datos en memoria, de momento es una
// solucion inicial porque luego se conectara todo a una base de datos
package memory

import (
	"errors"
	"restaurant-go-api/internal/core"
	"restaurant-go-api/internal/core/domain"
)

type InMemoryMenuRepository struct {
	menuItems map[string]*domain.MenuItem
}

func NewMenuRepository() core.MenuRepository {
	return &InMemoryMenuRepository{
		menuItems: make(map[string]*domain.MenuItem),
	}
}

func (r *InMemoryMenuRepository) Save(item *domain.MenuItem) error {
	r.menuItems[item.ID] = item
	return nil
}

func (r *InMemoryMenuRepository) FindByID(id string) (*domain.MenuItem, error) {
	item, exists := r.menuItems[id]
	if !exists {
		return nil, errors.New("menu item not found")
	}
	return item, nil
}

func (r *InMemoryMenuRepository) GetAll() ([]domain.MenuItem, error) {
	items := make([]domain.MenuItem, 0, len(r.menuItems))
	for _, item := range r.menuItems {
		items = append(items, *item)
	}
	return items, nil
}

func (r *InMemoryMenuRepository) Update(item *domain.MenuItem) error {
	if _, exists := r.menuItems[item.ID]; !exists {
		return errors.New("menu item not found")
	}
	r.menuItems[item.ID] = item
	return nil
}

func (r *InMemoryMenuRepository) Delete(id string) error {
	if _, exists := r.menuItems[id]; !exists {
		return errors.New("menu item not found")
	}
	delete(r.menuItems, id)
	return nil
}

func (r *InMemoryMenuRepository) FindByCategory(category domain.Category) ([]domain.MenuItem, error) {
	menuItems := make([]domain.MenuItem, 0)
	for _, item := range r.menuItems {
		if item.Category == category {
			menuItems = append(menuItems, *item)
		}
	}
	return menuItems, nil
}
