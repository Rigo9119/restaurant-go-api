package memory

import (
	"errors"
	"restaurant-go-api/internal/core"
	"restaurant-go-api/internal/core/domain"
)

type InMemoryUserRepository struct {
	user map[string]*domain.User
}

func NewUserRepository() core.UserRepository {
	return &InMemoryUserRepository{
		user: make(map[string]*domain.User),
	}
}

func (r *InMemoryUserRepository) Save(user *domain.User) error {
	r.user[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) FindByID(id string) (*domain.User, error) {
	user, exists := r.user[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *InMemoryUserRepository) GetAll() ([]domain.User, error) {
	users := make([]domain.User, 0, len(r.user))
	for _, user := range r.user {
		users = append(users, *user)
	}

	return users, nil
}

func (r *InMemoryUserRepository) Update(user *domain.User) error {
	if _, exists := r.user[user.ID]; !exists {
		return errors.New("user not found")
	}
	r.user[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(id string) error {
	if _, exists := r.user[id]; !exists {
		return errors.New("user not found")
	}
	delete(r.user, id)
	return nil
}
