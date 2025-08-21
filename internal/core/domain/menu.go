package domain

import (
	"restaurant-go-api/internal/shared/utils"
)

type MenuItem struct {
	ID        string
	Name      string
	Price     Price
	Category  Category
	Available bool
}

type Menu struct {
	ID        string
	MenuItems []MenuItem
}

func NewMenuItem(name, prefix string, category Category, price Price, available bool) *MenuItem {
	return &MenuItem{
		ID:        utils.GenerateRandomID(prefix),
		Name:      name,
		Price:     price,
		Category:  category,
		Available: available,
	}
}
