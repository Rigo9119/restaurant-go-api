package domain

type Size string

const (
	SizeSmall  Size = "small"
	SizeMedium Size = "medium"
	SizeLarge  Size = "large"
)

type PattyType string

const (
	PattyBeef    PattyType = "beef"
	PattyChicken PattyType = "chicken"
	PattyVeggie  PattyType = "veggie"
)

type BunType string

const (
	BunSesame  BunType = "sesame"
	BunBrioche BunType = "brioce"
	BunPotato  BunType = "potato"
	BunPretzel BunType = "pretzel"
)

type DrinkTemperatue string

const (
	Hot  DrinkTemperatue = "hot"
	Cold DrinkTemperatue = "cold"
	Room DrinkTemperatue = "room"
)

type Sauces string

const (
	Ketchup   Sauces = "ketchup"
	Mustard   Sauces = "mustard"
	Mayonaise Sauces = "mayonaise"
	Especial  Sauces = "Espcial"
	BBQ       Sauces = "bbq"
)

type AddOns string

const (
	Bacon     AddOns = "bacon"
	Cheese    AddOns = "cheese"
	Onion     AddOns = "onion"
	Tomato    AddOns = "tomato"
	Letttuce  AddOns = "lettuce"
	Pickles   AddOns = "Pickles"
	Peppers   AddOns = "peppers"
	Jalapenos AddOns = "jalapenos"
)

type Price struct {
	Amount int
}

type OrderStatus string

const (
	Ready     OrderStatus = "ready"
	Preparing OrderStatus = "preparing"
	Canceled  OrderStatus = "canceled"
	Ordered   OrderStatus = "ordered"
)
