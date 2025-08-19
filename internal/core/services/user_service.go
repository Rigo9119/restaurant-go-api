package services

import (
	"restaurant-go-api/internal/core"
	"restaurant-go-api/internal/core/domain"
)

type UserService interface {
	GetUserDetails(userID string) (*domain.User, error)
	CreateUser(name string, role domain.Role) (*domain.User, error)
	ValidateUserPermissions(userID string, action string) (bool, error)
	UpdateUserProfile(userID string, name string) (*domain.User, error)
}

// TODO: implement later
type userService struct {
	userRepo core.UserRepository
}

func NewUserService(userRepo core.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetUserDetails(userID string) (*domain.User, error) {
	return nil, nil
}

func (s *userService) CreateUser(name string, role domain.Role) (*domain.User, error) {
	return nil, nil
}

func (s *userService) ValidateUserPermissions(userID string, action string) (bool, error) {
	return false, nil
}

func (s *userService) UpdateUserProfile(userID, name string) (*domain.User, error) {
	return nil, nil
}
