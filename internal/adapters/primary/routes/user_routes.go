package routes

import (
	"restaurant-go-api/internal/adapters/primary/handlers"

	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Echo, userHandler *handlers.UserHandler) {
	userRoutes := e.Group("/user")

	userRoutes.POST("", userHandler.CreateUser)
	userRoutes.GET("/:id", userHandler.GetUserDetails)
	userRoutes.PUT("/:id", userHandler.UpdateUser)
	userRoutes.DELETE("/:id", userHandler.DeleteUser)
}
