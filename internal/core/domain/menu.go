package domain

import "slices"

type MenuItem struct {
	ID        string
	Name      string
	Price     Price
	Category  string
	Available bool
}

type Burguer struct {
	MenuItem
	PattyType PattyType
	BunType   BunType
	Sauces    []Sauces
	AddOns    []AddOns
}

type Side struct {
	MenuItem
	Portion Size
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
	Sides     []Side
	Drink     Drink
}

func (b *Burguer) CalculateBurguerPrice() int {
	// Calcula el precio de la hamburguesa
	totalPrice := b.MenuItem.Price.Amount
	switch b.PattyType {
	case PattyBeef:
		totalPrice += 4
	case PattyChicken:
		totalPrice += 3
	case PattyVeggie:
		totalPrice += 2

	default:
		return 0
	}

	switch b.BunType {
	case BunSesame, BunPotato:
		totalPrice += 1
	case BunBrioche:
		totalPrice += 2
	case BunPretzel:
		totalPrice += 3
	default:
		return 0
	}

	for _, sauce := range b.Sauces {
		switch sauce {
		case Ketchup, Mustard:
			totalPrice += 0 // Free basic sauces
		case Mayonnaise, BBQ:
			totalPrice += 1 // Premium sauces
		case Especial:
			totalPrice += 2 // Specialty sauce
		default:
			totalPrice += 1
		}
	}

	for _, addon := range b.AddOns {
		switch addon {
		case Lettuce, Tomato, Pickles, Onion:
			totalPrice += 0 // Free vegetables
		case Cheese:
			totalPrice += 2
		case Bacon:
			totalPrice += 3
		case Peppers, Jalapenos:
			totalPrice += 1 // Spicy extras
		default:
			totalPrice += 1
		}
	}

	return totalPrice
}

func (b *Burguer) MakeVegetarian() bool {
	if b.PattyType != PattyVeggie {
		return false
	}

	if slices.Contains(b.AddOns, Bacon) {
		return false
	}

	return true
}

func (s *Side) CalculateSidePrice() int {
	totalPrice := s.MenuItem.Price.Amount
	switch s.Portion {
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
