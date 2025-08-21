package services

import (
	"restaurant-go-api/internal/core"
	"restaurant-go-api/internal/core/domain"
)

type MenuService interface {
	// Menu
	CreateMenu(menuItem *domain.MenuItem) (*domain.Menu, error)
	UpdateMenu(itemID string, updates MenuItemUpdatesReq) (*domain.Menu, error)
	GetAllMenuItem()
	GetMenuByCategory()
	// menu items
	AddMenuItem(itemID string) (*domain.Menu, error)
	UpdateMenuItem(itemID string) (*domain.Menu, error)
	RemoveMenuItem(itemID string) error
	GetMenuItem(itemID string) (*domain.Menu, error)
	// Deals
	CreateMenuDeals() (*domain.Menu, error)
	UpdateMenuDeals() (*domain.Menu, error)
	RemoveMenuDeal() (*domain.Menu, error)
}

type MenuItemUpdatesReq struct {
	Name        *string
	Price       *domain.Price
	Description *string
	Category    *domain.Category
}

type menuService struct {
	userRepo core.UserRepository
	menuRepo core.MenuRepository
}

func NewMenuService(
	userRepo core.UserRepository,
	menurepo core.MenuRepository,
) MenuService {
	return &menuService{
		userRepo: userRepo,
		menuRepo: menurepo,
	}
}

func (s *menuService) CreateMenu() (*domain.Menu, error) {
	return nil, nil
}

func (s *menuService) RemoveItemFromMenu() (*domain.Menu, error) {
	return nil, nil
}

func (s *menuService) UpdateMenu() (*domain.Menu, error) {
	return nil, nil
}

func (s *menuService) AddMenuItem(itemID string) (*domain.Menu, error) {
	return nil, nil
}

func (s *menuService) CreateMenuDeals() (*domain.Menu, error) {
	return nil, nil
}

func (s *menuService) UpdateMenuDeals() (*domain.Menu, error) {
	return nil, nil
}
