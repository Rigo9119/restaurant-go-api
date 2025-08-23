package cmd

import (
	"fmt"
	"restaurant-go-api/internal/adapters/primary/handlers"
	"restaurant-go-api/internal/adapters/primary/routes"
	"restaurant-go-api/internal/adapters/secondary/memory"
	"restaurant-go-api/internal/core/domain"
	"restaurant-go-api/internal/core/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mattn/go-colorable"
)

func StartServer() error {
	// test user
	testUser := &domain.User{
		ID:   "user123",
		Name: "Test Customer",
		Role: domain.UserCustomer,
	}

	// repos - en este momento se guardan en memoria
	orderRepo := memory.NewOrderRepository()
	userRepo := memory.NewUserRepository()
	menuRepo := memory.NewMenuRepository()

	// guarda el test user
	if err := userRepo.Save(testUser); err != nil {
		panic(fmt.Sprintf("failed to save tet user: %v", err))
	}

	// servicios
	orderService := services.NewOrderService(orderRepo, userRepo, menuRepo)

	// handlers
	orderHandler := handlers.NewOrderHandler(orderService)

	e := echo.New()
	// middleware basico de echo
	e.Use(
		middleware.LoggerWithConfig(
			middleware.LoggerConfig{
				Format: `🚀${time_rfc3339} | ${status} | ${latency_human} | ${method} ${uri}`,
				Output: colorable.NewColorableStdout(),
			},
		),
	)
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// rutas especificas
	routes.SetupOrderRoutes(e, orderHandler)

	e.Logger.Info("Server starting on :8080")
	return e.Start(":8080")
}
