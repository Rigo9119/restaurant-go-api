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

type Side struct {
	MenuItem
	Size Size
}

type Drink struct {
	MenuItem
	Size         Size
	Temperature  DrinkTemperature
	IsCarbonated bool
	HasIce       bool
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

func NewBurger(name string, price Price, category Category, available bool, pattyType PattyType, bunType BunType, sauces []Sauce, addOns []AddOn) *Burger {
	menuItem := NewMenuItem(name, "burguer", Burguers, price, available)
	return &Burger{
		MenuItem:  *menuItem,
		PattyType: pattyType,
		BunType:   bunType,
		Sauces:    sauces,
		AddOns:    addOns,
	}
}

func NewSide(name, category string, price Price, available bool, size Size) *Side {
	menuItem := NewMenuItem(name, category, "side", price, available)

	return &Side{
		MenuItem: *menuItem,
		Size:     size,
	}
}

func NewDrink(name, category string, price Price, available, isCarbonated, hasIce bool, size Size, temp DrinkTemperature) *Drink {
	menuItem := NewMenuItem(name, category, "drink", price, available)
	return &Drink{
		MenuItem:     *menuItem,
		Size:         size,
		Temperature:  temp,
		IsCarbonated: isCarbonated,
		HasIce:       hasIce,
	}
}

func (s *Side) CalculateSidePrice() int {
	totalPrice := s.MenuItem.Price.Amount
	switch s.Size {
	case SizeSmall:
		totalPrice += 1
	case SizeMedium:
		totalPrice += 2
	case SizeLarge:
		totalPrice += 3
	default:
		totalPrice += 0
	}

	return totalPrice
}

func (d *Drink) CalculateDrinkPrice() int {
	totalPrice := d.MenuItem.Price.Amount
	switch d.Size {
	case SizeSmall:
		totalPrice += 1
	case SizeMedium:
		totalPrice += 2
	case SizeLarge:
		totalPrice += 3
	default:
		totalPrice += 0
	}

	switch d.Temperature {
	case Cold, Room:
		totalPrice += 0
	case Hot:
		totalPrice += 1
	default:
		totalPrice += 0
	}

	return totalPrice
}
