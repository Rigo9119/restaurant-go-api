package services

import (
	"restaurant-go-api/internal/core"
	"restaurant-go-api/internal/core/domain"
)

type UserService interface {
	GetUserDetails(userID string) (*domain.User, error)
	CreateUser(name string, role domain.Role) (*domain.User, error)
	ValidateUserPermissions(userID string, action string) (bool, error)
	UpdateUserProfile(userID string, update UserUpdateReq) (*domain.User, error)
}

type userService struct {
	userRepo core.UserRepository
}

func NewUserService(userRepo core.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetUserDetails(userID string) (*domain.User, error) {
	user, userErr := s.userRepo.FindByID(userID)
	if userErr != nil {
		return nil, userErr
	}
	return user, nil
}

func (s *userService) CreateUser(name string, role domain.Role) (*domain.User, error) {
	return nil, nil
}

func (s *userService) ValidateUserPermissions(userID string, action string) (bool, error) {
	user, userErr := s.userRepo.FindByID(userID)
	if userErr != nil {
		return false, userErr
	}

	permissions := user.GetPermissions()
	for _, permission := range permissions {
		if string(permission) == action {
			return true, nil
		}
	}
	return false, nil
}

// UserUpdateReq -> este objeto ayuda a manejar las propiedades
// que van a ser actulizadas, tambien puede cercer a medida que la entidad
// maneje mas propiedades
type UserUpdateReq struct {
	Name *string
	Role *domain.Role
}

func (s *userService) UpdateUserProfile(userID string, updates UserUpdateReq) (*domain.User, error) {
	user, userErr := s.userRepo.FindByID(userID)
	if userErr != nil {
		return nil, userErr
	}

	if updates.Name != nil {
		user.Name = *updates.Name
	}
	if updates.Role != nil {
		user.Role = *updates.Role
	}

	err := s.userRepo.Update(user)
	return nil, err
}
