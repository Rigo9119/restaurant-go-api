// Package domain burguer: lleva el control de la entidad burguer mantiene
// control sobre los metodos de la entidad
package domain

import "slices"

type Burger struct {
	MenuItem
	PattyType PattyType
	BunType   BunType
	Sauces    []Sauce
	AddOns    []AddOn
}

// Builder pattern -> cuando las funciones constructoras / factories tienen
// muchos argumentos es recomendable user el "Builder pattern", en este caso es
// una struct BurgerBuilder que tiene los argumentos que van usar los metodos de
// la hamburguesa
type BurgerBuilder struct {
	name      string
	price     Price
	available bool
	pattyType PattyType
	bunType   BunType
	sauces    []Sauce
	addOns    []AddOn
}

// Factory -> est funcion acepta parte de los argumentos que se necesitan para
// construir una hamburguesa, setea los etados default del builder.
func NewBurgerBuilder(name string, price Price) *BurgerBuilder {
	return &BurgerBuilder{
		name:      name,
		price:     price,
		available: true,
		sauces:    []Sauce{},
		addOns:    []AddOn{},
	}
}

// Setter methods -> estos nos ayudan a controlar y modificar el estado interno
// del builder, retornan el mismo builder y nos permiten encadenar "chain"
// methods entre si para un mejor manejo de la API
func (b *BurgerBuilder) WithPatty(pattyType PattyType) *BurgerBuilder {
	b.pattyType = pattyType
	return b
}

func (b *BurgerBuilder) WithBun(bunType BunType) *BurgerBuilder {
	b.bunType = bunType
	return b
}

func (b *BurgerBuilder) AddSauces(sauces ...Sauce) *BurgerBuilder {
	b.sauces = append(b.sauces, sauces...)
	return b
}

func (b *BurgerBuilder) AddAddOns(addOns ...AddOn) *BurgerBuilder {
	b.addOns = append(b.addOns, addOns...)
	return b
}

func (b *BurgerBuilder) SetAvailable(available bool) *BurgerBuilder {
	b.available = available
	return b
}

// esta funcion aplica el resto de las funciones factory y crea una nuev
// ahamburguesa, entonces el proceso es:
func (b *BurgerBuilder) Build() *Burger {
	menuItem := NewMenuItem(
		b.name,
		"burger",
		Burguers,
		b.price,
		b.available,
	)
	return &Burger{
		MenuItem:  *menuItem,
		PattyType: b.pattyType,
		BunType:   b.bunType,
		Sauces:    b.sauces,
		AddOns:    b.addOns,
	}
}

func (b *Burger) CalculateBurguerPrice() int {
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
			totalPrice += 0
		case Mayonnaise, BBQ:
			totalPrice += 1
		case Especial:
			totalPrice += 2
		default:
			totalPrice += 1
		}
	}

	for _, addon := range b.AddOns {
		switch addon {
		case Lettuce, Tomato, Pickles, Onion:
			totalPrice += 0
		case Cheese:
			totalPrice += 2
		case Bacon:
			totalPrice += 3
		case Peppers, Jalapenos:
			totalPrice += 1
		default:
			totalPrice += 1
		}
	}

	return totalPrice
}

func (b *Burger) MakeVegetarian() bool {
	if b.PattyType != PattyVeggie {
		return false
	}

	if slices.Contains(b.AddOns, Bacon) {
		return false
	}

	return true
}
