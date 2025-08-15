package dto

// Menu Request DTOs - what clients send to the server

type CreateMenuItemRequest struct {
	Name      string `json:"name" validate:"required,min=2,max=100"`
	Price     int    `json:"price" validate:"required,min=1"`
	Category  string `json:"category" validate:"required"`
	Available bool   `json:"available"`
}

type CreateBurgerRequest struct {
	CreateMenuItemRequest
	PattyType string   `json:"patty_type" validate:"required,oneof=beef chicken veggie"`
	BunType   string   `json:"bun_type" validate:"required,oneof=sesame brioche potato pretzel"`
	Sauces    []string `json:"sauces,omitempty"`
	AddOns    []string `json:"add_ons,omitempty"`
}

// Menu Response DTOs - what server sends to clients

type MenuItemResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Category  string `json:"category"`
	Available bool   `json:"available"`
}

type BurgerResponse struct {
	MenuItemResponse
	PattyType string   `json:"patty_type"`
	BunType   string   `json:"bun_type"`
	Sauces    []string `json:"sauces"`
	AddOns    []string `json:"add_ons"`
}

type MenuResponse struct {
	MenuItems []MenuItemResponse `json:"menu_items"`
	Burgers   []BurgerResponse   `json:"burgers"`
	Count     int                `json:"count"`
}