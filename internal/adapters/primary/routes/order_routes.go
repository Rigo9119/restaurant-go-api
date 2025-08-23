package routes

import (
	"restaurant-go-api/internal/adapters/primary/handlers"

	"github.com/labstack/echo/v4"
)

func SetupOrderRoutes(e *echo.Echo, orderHandler *handlers.OrderHandler) {
	orderGroup := e.Group("/orders")

	orderGroup.POST("", orderHandler.CreateOrder)
	orderGroup.GET("/:id", orderHandler.GetOrder)
	orderGroup.POST("/:id/items", orderHandler.AddItemToOrder)
	orderGroup.PUT("/:id/status", orderHandler.UpdateOrderStatus)
}
