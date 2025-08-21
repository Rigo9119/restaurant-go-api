package domain

type Side struct {
	MenuItem
	Size Size
}

func NewSide(name, category string, price Price, available bool, size Size) *Side {
	menuItem := NewMenuItem(name, category, "side", price, available)

	return &Side{
		MenuItem: *menuItem,
		Size:     size,
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
