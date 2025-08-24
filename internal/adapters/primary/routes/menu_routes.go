package routes

import (
	"restaurant-go-api/internal/adapters/primary/handlers"

	"github.com/labstack/echo/v4"
)

func SetupMenuRoutes(e *echo.Echo, menuHandler *handlers.MenuHandler) {
	menuGroup := e.Group("menu")

	menuGroup.POST("/items", menuHandler.CreateMenuItem)
	menuGroup.GET("/items", menuHandler.GetAllMenuItems)
	menuGroup.GET("/items/:id", menuHandler.GetMenuItem)
	menuGroup.PUT("/items/:id", menuHandler.UpdateMenuItem)
	menuGroup.DELETE("/items/:id", menuHandler.DeleteMenuItem)
	menuGroup.GET("/category/:category", menuHandler.GetMenuByCategory)

	// Deal routes
	menuGroup.POST("/deals", menuHandler.CreateDeal)
	menuGroup.GET("/deals", menuHandler.GetActiveDeals)
	menuGroup.PUT("/deals/:id", menuHandler.UpdateDeal)
	menuGroup.DELETE("/deals/:id", menuHandler.DeleteDeal)
}
