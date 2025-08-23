package services

import (
	"errors"
	"restaurant-go-api/internal/core"
	"restaurant-go-api/internal/core/domain"
)

type MenuService interface {
	// MenuItem
	AddMenuItem(item *domain.MenuItem) (*domain.MenuItem, error)
	UpdateMenuItem(itemID string, updates MenuItemUpdatesReq) (*domain.MenuItem, error)
	RemoveMenuItem(itemID string) error
	GetMenuItem(itemID string) (*domain.MenuItem, error)

	// Menu
	GetAllMenuItems() ([]domain.MenuItem, error)
	GetMenuByCategory(category domain.Category) ([]domain.MenuItem, error)

	// Deal
	CreateDeal(deal *domain.Deal) (*domain.Deal, error)
	UpdateDeal(dealID string, updates DealUpdatesReq) (*domain.Deal, error)
	RemoveDeal(dealID string) error
	GetActiveDeals() ([]domain.Deal, error)
}

type MenuItemUpdatesReq struct {
	Name        *string
	Price       *domain.Price
	Description *string
	Category    *domain.Category
}

type DealUpdatesReq struct {
	Name            *string
	Description     *string
	DiscountPercent *int
	DealPrice       *domain.Price
}

type menuService struct {
	userRepo core.UserRepository
	menuRepo core.MenuRepository
	dealRepo core.DealRepository
	// Note: Deal repository would be needed when you implement deal functionality
}

func NewMenuService(
	userRepo core.UserRepository,
	menuRepo core.MenuRepository,
	dealRepo core.DealRepository,
) MenuService {
	return &menuService{
		userRepo: userRepo,
		menuRepo: menuRepo,
		dealRepo: dealRepo,
	}
}

func (s *menuService) AddMenuItem(item *domain.MenuItem) (*domain.MenuItem, error) {
	if item == nil {
		return nil, errors.New("item field should not be empty")
	}

	if item.Name == "" {
		return nil, errors.New("item name is required")
	}
	if item.Price.Amount <= 0 {
		return nil, errors.New("item price must be greater than 0")
	}

	saveErr := s.menuRepo.Save(item)
	if saveErr != nil {
		return nil, saveErr
	}

	return item, nil
}

func (s *menuService) UpdateMenuItem(itemID string, updates MenuItemUpdatesReq) (*domain.MenuItem, error) {
	if itemID == "" {
		return nil, errors.New("the item id field should not be empty")
	}

	menuItem, menuItemErr := s.menuRepo.FindByID(itemID)
	if menuItemErr != nil {
		return nil, menuItemErr
	}

	if updates.Name != nil {
		menuItem.Name = *updates.Name
	}
	if updates.Price != nil {
		menuItem.Price = *updates.Price
	}
	if updates.Description != nil {
		menuItem.Description = *updates.Description
	}
	if updates.Category != nil {
		menuItem.Category = *updates.Category
	}

	saveErr := s.menuRepo.Update(menuItem)
	if saveErr != nil {
		return nil, saveErr
	}

	return menuItem, nil
}

func (s *menuService) RemoveMenuItem(itemID string) error {
	if itemID == "" {
		return errors.New("itemID should not be emtpy")
	}

	_, err := s.menuRepo.FindByID(itemID)
	if err != nil {
		return err
	}
	deleteErr := s.menuRepo.Delete(itemID)
	return deleteErr
}

func (s *menuService) GetMenuItem(itemID string) (*domain.MenuItem, error) {
	if itemID == "" {
		return nil, errors.New("itemID should not be empty")
	}
	menuItem, menuItemErr := s.menuRepo.FindByID(itemID)
	if menuItemErr != nil {
		return nil, menuItemErr
	}
	return menuItem, nil
}

func (s *menuService) GetAllMenuItems() ([]domain.MenuItem, error) {
	items, err := s.menuRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *menuService) GetMenuByCategory(category domain.Category) ([]domain.MenuItem, error) {
	menuCategory, menuCategoryErr := s.menuRepo.FindByCategory(category)
	if menuCategoryErr != nil {
		return nil, menuCategoryErr
	}
	return menuCategory, nil
}

func (s *menuService) CreateDeal(deal *domain.Deal) (*domain.Deal, error) {
	return nil, nil
}

func (s *menuService) UpdateDeal(dealID string, updates DealUpdatesReq) (*domain.Deal, error) {
	return nil, nil
}

func (s *menuService) RemoveDeal(dealID string) error {
	if dealID == "" {
		return errors.New("dealID fiedl should not be empty")
	}

	removeErr := s.dealRepo.Delete(dealID)
	return removeErr
}

func (s *menuService) GetActiveDeals() ([]domain.Deal, error) {
	activeDeals, activeDealsErr := s.dealRepo.FindActiveDeals()
	if activeDealsErr != nil {
		return nil, activeDealsErr
	}

	return activeDeals, nil
}
