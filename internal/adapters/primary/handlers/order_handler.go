package handlers

import (
	"net/http"
	"restaurant-go-api/internal/adapters/primary/dto"
	"restaurant-go-api/internal/core"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService core.OrderService
}

func NewOrderHandler(orderService core.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var req dto.CreateOrderRequest
	// c.Bind(&req) hace el parce de JSON
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request format",
			Code:    http.StatusBadRequest,
		})
	} // esto hace el Parse JSON req

	if req.UserID == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "user_id is required",
			Code:    http.StatusBadRequest,
		})
	}

	order, err := h.orderService.CreateCustomerOrder(req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	response := dto.OrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		OrderItems:  []dto.ItemDTO{},
		OrderTotal:  order.OrderTotal,
		OrderStatus: string(order.OrderStatus),
	}

	return c.JSON(http.StatusCreated, response)
}
