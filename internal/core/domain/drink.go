package domain

type Drink struct {
	MenuItem
	Size         Size
	Temperature  DrinkTemperature
	IsCarbonated bool
	HasIce       bool
}

type DrinkBuilder struct {
	name         string
	price        Price
	size         Size
	temp         DrinkTemperature
	isCarbonated bool
	hasIce       bool
	available    bool
}

func NewDrinkBuilder(name string, price Price) *DrinkBuilder {
	return &DrinkBuilder{
		name:         name,
		price:        price,
		size:         SizeMedium,
		temp:         Cold,
		isCarbonated: true,
		hasIce:       true,
		available:    true,
	}
}

func (b *DrinkBuilder) SetSize(size Size) *DrinkBuilder {
	b.size = size
	return b
}

func (b *DrinkBuilder) SetTemperature(temp DrinkTemperature) *DrinkBuilder {
	b.temp = temp
	return b
}

func (b *DrinkBuilder) SetCarbonated(carbonated bool) *DrinkBuilder {
	b.isCarbonated = carbonated
	return b
}

func (b *DrinkBuilder) SetIce(hasIce bool) *DrinkBuilder {
	b.hasIce = hasIce
	return b
}

func (b *DrinkBuilder) SetAvailable(available bool) *DrinkBuilder {
	b.available = available
	return b
}

func (b *DrinkBuilder) Build() *Drink {
	menuItem := NewMenuItem(
		b.name,
		"drink",
		Drinks,
		b.price,
		b.available,
	)
	return &Drink{
		MenuItem:     *menuItem,
		Size:         b.size,
		Temperature:  b.temp,
		IsCarbonated: b.isCarbonated,
		HasIce:       b.hasIce,
	}
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
