package handlers

import (
	"net/http"
	"restaurant-go-api/internal/adapters/primary/dto"
	"restaurant-go-api/internal/core/domain"
	"restaurant-go-api/internal/core/services"

	"github.com/labstack/echo/v4"
)

type MenuHandler struct {
	menuService services.MenuService
}

func NewMenuHandler(menuService services.MenuService) *MenuHandler {
	return &MenuHandler{
		menuService: menuService,
	}
}

func (h *MenuHandler) CreateMenuItem(c echo.Context) error {
	var req dto.CreateMenuItemRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request format",
			Code:    http.StatusBadRequest,
		})
	}

	if req.Name == "" {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Name field should not be empty",
				Code:    http.StatusBadRequest,
			},
		)
	}

	if req.Price == 0 {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Price field should not be equal to 0",
				Code:    http.StatusBadRequest,
			},
		)
	}

	if req.Category == "" {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Category field should not be empty",
				Code:    http.StatusBadRequest,
			})
	}

	reqMenuItem := domain.NewMenuItem(
		req.Name,
		req.Description,
		"menu",
		domain.Category(req.Category),
		domain.Price{Amount: req.Price},
		req.Available,
	)

	menuItem, err := h.menuService.AddMenuItem(reqMenuItem)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
	}

	response := dto.MenuItemResponse{
		ID:        menuItem.ID,
		Name:      menuItem.Name,
		Price:     menuItem.Price.Amount,
		Category:  string(menuItem.Category),
		Available: menuItem.Available,
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *MenuHandler) GetMenuItem(c echo.Context) error {
	itemID := c.Param("id")
	if itemID == "" {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "id field should not be empty",
				Code:    http.StatusBadRequest,
			})
	}

	menuItem, err := h.menuService.GetMenuItem(itemID)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
	}

	response := dto.MenuItemResponse{
		ID:        menuItem.ID,
		Name:      menuItem.Name,
		Price:     menuItem.Price.Amount,
		Category:  string(menuItem.Category),
		Available: menuItem.Available,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *MenuHandler) GetAllMenuItems(c echo.Context) error {
	serviceItems, err := h.menuService.GetAllMenuItems()
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Bad response",
				Code:    http.StatusBadRequest,
			})
	}

	resSlice := make([]dto.MenuItemResponse, len(serviceItems))
	for i, menuItem := range serviceItems {
		resSlice[i] = dto.MenuItemResponse{
			ID:        menuItem.ID,
			Name:      menuItem.Name,
			Price:     menuItem.Price.Amount,
			Category:  string(menuItem.Category),
			Available: menuItem.Available,
		}
	}

	return c.JSON(http.StatusOK, resSlice)
}

func (h *MenuHandler) GetMenuByCategory(c echo.Context) error {
	category := c.Param("category")
	if category == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Category field must not be empty",
			Code:    http.StatusBadRequest,
		})
	}

	items, err := h.menuService.GetMenuByCategory(domain.Category(category))
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		})
	}

	resSlice := make([]dto.MenuItemResponse, len(items))
	for i, item := range items {
		resSlice[i] = dto.MenuItemResponse{
			ID:        item.ID,
			Name:      item.Name,
			Price:     item.Price.Amount,
			Category:  string(item.Category),
			Available: item.Available,
		}
	}

	return c.JSON(http.StatusOK, resSlice)
}

func (h *MenuHandler) UpdateMenuItem(c echo.Context) error {
	itemID := c.Param("id")
	if itemID == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "ID field must not be empty",
			Code:    http.StatusBadRequest,
		})
	}
	return nil
}

func (h *MenuHandler) DeleteMenuItem(c echo.Context) error {
	return nil
}

// deal handlers
func (h *MenuHandler) CreateDeal(c echo.Context) error {
	return nil
}

func (h *MenuHandler) GetActiveDeals(c echo.Context) error {
	return nil
}

func (h *MenuHandler) UpdateDeal(c echo.Context) error {
	return nil
}

func (h *MenuHandler) DeleteDeal(c echo.Context) error {
	return nil
}
