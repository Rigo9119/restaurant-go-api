package services

import (
	"errors"
	"restaurant-go-api/internal/core"
	"restaurant-go-api/internal/core/domain"
)

type UserService interface {
	GetUserDetails(userID string) (*domain.User, error)
	CreateUser(name string, role domain.Role) (*domain.User, error)
	ValidateUserPermissions(userID string, action string) (bool, error)
	UpdateUserProfile(userID string, update UserUpdateReq) (*domain.User, error)
	DeleteUser(userID string) error
}

type userService struct {
	userRepo core.UserRepository
}

// UserUpdateReq -> este objeto ayuda a manejar las propiedades
// que van a ser actulizadas, tambien puede cercer a medida que la entidad
// maneje mas propiedades
type UserUpdateReq struct {
	Name *string
	Role *domain.Role
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
	if name == "" {
		return nil, errors.New("name field should not be empty")
	}

	if role == "" {
		return nil, errors.New("role field should not be empty")
	}

	newUser := domain.NewUser(role, name)
	if newUser == nil {
		return nil, errors.New("invalid role provided")
	}

	saveErr := s.userRepo.Save(newUser)
	if saveErr != nil {
		return nil, saveErr
	}

	return newUser, nil
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
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(userID string) error {
	if userID == "" {
		return errors.New("userID should not be empty")
	}

	// SIEMPRE VERIFICAR SI LA ENTIDAD EXISTE
	_, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	return s.userRepo.Delete(userID)
}
