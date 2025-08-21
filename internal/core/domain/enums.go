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

type Sauce string

const (
	Ketchup    Sauce = "ketchup"
	Mustard    Sauce = "mustard"
	Mayonnaise Sauce = "mayonnaise"
	Especial   Sauce = "Especial"
	BBQ        Sauce = "bbq"
)

type AddOn string

const (
	Bacon     AddOn = "bacon"
	Cheese    AddOn = "cheese"
	Onion     AddOn = "onion"
	Tomato    AddOn = "tomato"
	Lettuce   AddOn = "lettuce"
	Pickles   AddOn = "pickles"
	Peppers   AddOn = "peppers"
	Jalapenos AddOn = "jalapenos"
)

type Price struct {
	Amount int
}

type OrderStatus string

const (
	Ready     OrderStatus = "ready"
	Preparing OrderStatus = "preparing"
	Cancelled OrderStatus = "cancelled"
	Ordered   OrderStatus = "ordered"
	Delivered OrderStatus = "delivered"
)

type Category string

const (
	Burguers Category = "burguers"
	Sides    Category = "sides"
	Drinks   Category = "drinks"
)

type Permission string

const (
	PermissionCreateOrder    Permission = "create_order"
	PermissionViewMenu       Permission = "view_menu"
	PermissionModifyOwnOrder Permission = "modify_own_order"
	PermissionManageMenu     Permission = "manage_menu"
	PermissionViewOrders     Permission = "view_orders"
	PermissionUpdatePrices   Permission = "update_prices"
	PermissionManageUsers    Permission = "manage_users"
	PermissionViewAllOrders  Permission = "view_all_orders"
	PermissionSystemAdmin    Permission = "system_admin"
)
