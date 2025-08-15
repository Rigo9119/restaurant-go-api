package dto

type CreateOrderRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type AddItemToOrderRequest struct {
	ItemID   string `json:"item_id" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,min=1"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

type OrderResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	OrderItems  []ItemDTO `json:"order_items"`
	OrderTotal  int       `json:"order_total"`
	OrderStatus string    `json:"order_status"`
}

type ItemDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

