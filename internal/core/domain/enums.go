package domain

type Role string

const (
	UserAdmin    Role = "admin"
	UserCustomer Role = "customer"
	UserManager  Role = "manager"
)

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
	BunBrioche BunType = "brioche"
	BunPotato  BunType = "potato"
	BunPretzel BunType = "pretzel"
)

type DrinkTemperature string

const (
	Hot  DrinkTemperature = "hot"
	Cold DrinkTemperature = "cold"
	Room DrinkTemperature = "room"
)

type Sauces string

const (
	Ketchup    Sauces = "ketchup"
	Mustard    Sauces = "mustard"
	Mayonnaise Sauces = "mayonnaise"
	Especial   Sauces = "Especial"
	BBQ        Sauces = "bbq"
)

type AddOns string

const (
	Bacon     AddOns = "bacon"
	Cheese    AddOns = "cheese"
	Onion     AddOns = "onion"
	Tomato    AddOns = "tomato"
	Lettuce   AddOns = "lettuce"
	Pickles   AddOns = "pickles"
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
