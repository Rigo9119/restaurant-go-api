package handlers

import (
	"net/http"
	"restaurant-go-api/internal/adapters/primary/dto"
	"restaurant-go-api/internal/core/domain"
	"restaurant-go-api/internal/core/services"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Invalid request format",
				Code:    http.StatusBadRequest,
			})
	}

	if req.Name == "" {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Name is required",
				Code:    http.StatusBadRequest,
			})
	}

	if req.Role == "" {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Role is required",
				Code:    http.StatusBadRequest,
			})
	}

	user, err := h.userService.CreateUser(
		req.Name, domain.Role(req.Role),
	)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.ErrorResponse{
				Message: "Failed to create user",
				Code:    http.StatusInternalServerError,
			})
	}

	response := dto.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Role: string(user.Role),
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) GetUserDetails(c echo.Context) error {
	userID := c.Param("id")
	if userID == "" {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Bad request",
				Code:    http.StatusBadRequest,
			})
	}

	user, err := h.userService.GetUserDetails(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Failed to get user details",
			Code:    http.StatusBadRequest,
		})
	}

	response := dto.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Role: string(user.Role),
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	userID := c.Param("id")
	if userID == "" {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Bad request",
				Code:    http.StatusBadRequest,
			})
	}

	var req dto.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Failed to bind request",
			Code:    http.StatusBadRequest,
		})
	}

	updates := services.UserUpdateReq{}

	updates.Name = req.Name

	if req.Role != nil {
		roleValue := domain.Role(*req.Role)
		updates.Role = &roleValue
	}

	updateUser, err := h.userService.UpdateUserProfile(userID, updates)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Failed to update user",
			Code:    http.StatusInternalServerError,
		})
	}

	response := dto.UserResponse{
		ID:   updateUser.ID,
		Name: updateUser.Name,
		Role: string(updateUser.Role),
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	userID := c.Param("id")
	if userID == "" {
		return c.JSON(
			http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "Bad request",
				Code:    http.StatusBadRequest,
			})
	}

	err := h.userService.DeleteUser(userID)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.ErrorResponse{
				Message: "Internal server error",
				Code:    http.StatusInternalServerError,
			})
	}

	return c.JSON(http.StatusNoContent, nil)
}
