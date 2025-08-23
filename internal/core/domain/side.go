package domain

type Side struct {
	MenuItem
	Size Size
}

type SideBuilder struct {
	name        string
	category    Category
	price       Price
	available   bool
	size        Size
	description string
}

func NewSideBuilder(name string, price Price) *SideBuilder {
	return &SideBuilder{
		name:      name,
		category:  Sides,
		price:     price,
		available: true,
		size:      SizeSmall,
	}
}

func (b *SideBuilder) SetSize(size Size) *SideBuilder {
	b.size = size
	return b
}

func (b *SideBuilder) SetAvailable(available bool) *SideBuilder {
	b.available = available
	return b
}

func (b *SideBuilder) Build() *Side {
	menuItem := NewMenuItem(
		b.name,
		b.description,
		"side",
		Sides,
		b.price,
		b.available,
	)
	return &Side{
		MenuItem: *menuItem,
		Size:     b.size,
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
