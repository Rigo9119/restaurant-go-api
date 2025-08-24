package handlers

import (
	"net/http"
	"restaurant-go-api/internal/adapters/primary/dto"
	"restaurant-go-api/internal/core/domain"
	"restaurant-go-api/internal/core/services"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var req dto.CreateOrderRequest
	// c.Bind(&req) hace el parse de JSON
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request format",
			Code:    http.StatusBadRequest,
		})
	}

	// esto hace el Parse JSON req
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

func (h *OrderHandler) GetOrder(c echo.Context) error {
	orderID := c.Param("id")
	if orderID == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Orde ID is required",
			Code:    http.StatusBadRequest,
		})
	}

	order, getErr := h.orderService.GetOrderDetails(orderID)
	if getErr != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: getErr.Error(),
			Code:    http.StatusNotFound,
		})
	}
	orderItems := make([]dto.ItemDTO, len(order.OrderItems))

	for i, item := range order.OrderItems {
		orderItems[i] = dto.ItemDTO{
			ID:       item.ID,
			Name:     item.Name,
			Price:    item.Price.Amount,
			Category: string(item.Category),
		}
	}

	response := dto.OrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		OrderItems:  orderItems,
		OrderTotal:  order.OrderTotal,
		OrderStatus: string(order.OrderStatus),
	}

	return c.JSON(http.StatusOK, response)
}

func (h *OrderHandler) AddItemToOrder(c echo.Context) error {
	orderID := c.Param("id")
	if orderID == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Bad request, ID should not be empty",
			Code:    http.StatusBadRequest,
		})
	}

	var req dto.AddItemToOrderRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Bad request",
			Code:    http.StatusBadRequest,
		})
	}

	if req.ItemID == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Bad request, ID should not be empty",
			Code:    http.StatusBadRequest,
		})
	}

	if req.Quantity <= 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Bad request, Order queantity should be greater than 0",
			Code:    http.StatusBadRequest,
		})
	}

	updateOrder, err := h.orderService.AddItemToCustomerOrder(orderID, req.ItemID, req.Quantity)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	orderItems := make([]dto.ItemDTO, len(updateOrder.OrderItems))
	for i, item := range updateOrder.OrderItems {
		orderItems[i] = dto.ItemDTO{
			ID:       item.ID,
			Name:     item.Name,
			Price:    item.Price.Amount,
			Category: string(item.Category),
		}
	}

	response := dto.OrderResponse{
		ID:          updateOrder.ID,
		UserID:      updateOrder.UserID,
		OrderItems:  orderItems,
		OrderTotal:  updateOrder.OrderTotal,
		OrderStatus: string(updateOrder.OrderStatus),
	}

	return c.JSON(http.StatusOK, response)
}

func (h *OrderHandler) UpdateOrderStatus(c echo.Context) error {
	orderID := c.Param("id")
	if orderID == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Bad request, orderID should not be empty",
			Code:    http.StatusBadRequest,
		})
	}

	var req dto.UpdateOrderStatusRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	if req.Status == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Bad request, status should be not empty",
			Code:    http.StatusBadRequest,
		})
	}
	updateOrder, err := h.orderService.UpdateOrderStatus(orderID,
		domain.OrderStatus(req.Status))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	orderItems := make([]dto.ItemDTO, len(updateOrder.OrderItems))
	for i, item := range updateOrder.OrderItems {
		orderItems[i] = dto.ItemDTO{
			ID:       item.ID,
			Name:     item.Name,
			Price:    item.Price.Amount,
			Category: string(item.Category),
		}
	}

	response := dto.OrderResponse{
		ID:          updateOrder.ID,
		UserID:      updateOrder.UserID,
		OrderItems:  orderItems,
		OrderTotal:  updateOrder.OrderTotal,
		OrderStatus: string(updateOrder.OrderStatus),
	}

	return c.JSON(http.StatusOK, response)
}
