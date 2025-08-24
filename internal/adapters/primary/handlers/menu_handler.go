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

	var req dto.UpdateMenuItemRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Invalid request format",
				Code:    http.StatusBadRequest,
			})
	}

	updates := services.MenuItemUpdatesReq{}

	if req.Price != nil {
		priceValue := domain.Price{Amount: *req.Price}
		updates.Price = &priceValue
	}

	if req.Category != nil {
		categoryValue := domain.Category(*req.Category)
		updates.Category = &categoryValue
	}

	updates.Name = req.Name
	updates.Description = req.Description

	updateItem, err := h.menuService.UpdateMenuItem(itemID, updates)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
	}

	response := dto.MenuItemResponse{
		ID:        updateItem.ID,
		Name:      updateItem.Name,
		Price:     updateItem.Price.Amount,
		Category:  string(updateItem.Category),
		Available: updateItem.Available,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *MenuHandler) DeleteMenuItem(c echo.Context) error {
	itemID := c.Param("id")
	if itemID == "" {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Bad request",
				Code:    http.StatusBadRequest,
			})
	}

	err := h.menuService.RemoveMenuItem(itemID)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResponse{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
	}
	return c.JSON(http.StatusNoContent, nil)
}

// deal handlers -> estos tienen que ir luego en su propio handler
func (h *MenuHandler) CreateDeal(c echo.Context) error {
	var req dto.CreateDealRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request format",
			Code:    http.StatusBadRequest,
		})
	}
	// de momento solo valida el nombre
	if req.Name == "" {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Name field should not be empty",
				Code:    http.StatusBadRequest,
			},
		)
	}

	var menuItems []domain.MenuItem
	for _, itemID := range req.Items {
		item, err := h.menuService.GetMenuItem(itemID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Message: "Invalid item ID: " + itemID,
				Code:    http.StatusBadRequest,
			})
		}

		menuItems = append(menuItems, *item)
	}

	dealPrice := domain.Price{
		Amount: req.DealPrice.Amount,
	}

	deal := domain.NewDeal(
		req.Name,
		req.Description,
		menuItems,
		dealPrice,
		req.StartDate,
		req.EndDate,
	)

	createdDeal, err := h.menuService.CreateDeal(deal)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	itemsDTOs := make([]dto.ItemDTO, len(createdDeal.Items))
	for i, item := range createdDeal.Items {
		itemsDTOs[i] = dto.ItemDTO{
			ID:       item.ID,
			Name:     item.Name,
			Price:    item.Price.Amount,
			Category: string(item.Category),
		}
	}

	response := dto.DealResponse{
		ID:              createdDeal.ID,
		Name:            createdDeal.Name,
		Description:     createdDeal.Description,
		DiscountPercent: createdDeal.DiscountPercent,
		StartDate:       createdDeal.StartDate,
		EndDate:         createdDeal.EndDate,
		Items:           itemsDTOs,
		DealPrice:       dto.PriceDTO{Amount: createdDeal.DealPrice.Amount},
		IsActive:        createdDeal.IsActive(),
	}
	return c.JSON(http.StatusCreated, response)
}

func (h *MenuHandler) GetActiveDeals(c echo.Context) error {
	deals, err := h.menuService.GetActiveDeals()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	dealsResponse := make([]dto.DealResponse, len(deals))
	for i, deal := range deals {
		itemDTOs := make([]dto.ItemDTO, len(deal.Items))
		for j, item := range deal.Items {
			itemDTOs[j] = dto.ItemDTO{
				ID:       item.ID,
				Name:     item.Name,
				Price:    item.Price.Amount,
				Category: string(item.Category),
			}
		}

		dealsResponse[i] = dto.DealResponse{
			ID:              deal.ID,
			Name:            deal.Name,
			Description:     deal.Description,
			DiscountPercent: deal.DiscountPercent,
			StartDate:       deal.StartDate,
			EndDate:         deal.EndDate,
			Items:           itemDTOs,
			DealPrice:       dto.PriceDTO{Amount: deal.DealPrice.Amount},
			IsActive:        deal.IsActive(),
		}
	}

	return c.JSON(http.StatusOK, dealsResponse)
}

func (h *MenuHandler) UpdateDeal(c echo.Context) error {
	dealID := c.Param("id")
	if dealID == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "ID field must not be empty",
			Code:    http.StatusBadRequest,
		})
	}

	var req dto.UpdateDealRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid request format",
			Code:    http.StatusBadRequest,
		})
	}

	updates := services.DealUpdatesReq{}

	// Simple pointer fields
	updates.Name = req.Name
	updates.Description = req.Description
	updates.DiscountPercent = req.DiscountPercent

	// Handle PriceDTO conversion
	if req.DealPrice != nil {
		priceValue := domain.Price{Amount: req.DealPrice.Amount}
		updates.DealPrice = &priceValue
	}

	// Call service
	updatedDeal, err := h.menuService.UpdateDeal(dealID, updates)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	// Convert items for response
	itemDTOs := make([]dto.ItemDTO, len(updatedDeal.Items))
	for i, item := range updatedDeal.Items {
		itemDTOs[i] = dto.ItemDTO{
			ID:       item.ID,
			Name:     item.Name,
			Price:    item.Price.Amount,
			Category: string(item.Category),
		}
	}

	response := dto.DealResponse{
		ID:              updatedDeal.ID,
		Name:            updatedDeal.Name,
		Description:     updatedDeal.Description,
		DiscountPercent: updatedDeal.DiscountPercent,
		StartDate:       updatedDeal.StartDate,
		EndDate:         updatedDeal.EndDate,
		Items:           itemDTOs,
		DealPrice:       dto.PriceDTO{Amount: updatedDeal.DealPrice.Amount},
		IsActive:        updatedDeal.IsActive(),
	}

	return c.JSON(http.StatusOK, response)
}

func (h *MenuHandler) DeleteDeal(c echo.Context) error {
	dealID := c.Param("id")
	if dealID == "" {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Bad request",
				Code:    http.StatusBadRequest,
			})
	}

	err := h.menuService.RemoveDeal(dealID)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResponse{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
	}

	return c.JSON(http.StatusNoContent, nil)
}
