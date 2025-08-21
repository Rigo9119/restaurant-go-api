package memory

import (
	"errors"
	"restaurant-go-api/internal/core"
	"restaurant-go-api/internal/core/domain"
)

type InMemoryDealRepository struct {
	deals map[string]*domain.Deal
}

func NewDealRepository() core.DealRepository {
	return &InMemoryDealRepository{
		deals: make(map[string]*domain.Deal),
	}
}

func (r *InMemoryDealRepository) Save(deal *domain.Deal) error {
	r.deals[deal.ID] = deal
	return nil
}

func (r *InMemoryDealRepository) FindByID(id string) (*domain.Deal, error) {
	deal, exists := r.deals[id]
	if !exists {
		return nil, errors.New("deal not found")
	}
	return deal, nil
}

func (r *InMemoryDealRepository) GetAll() ([]domain.Deal, error) {
	deals := make([]domain.Deal, 0, len(r.deals))
	for _, deal := range r.deals {
		deals = append(deals, *deal)
	}
	return deals, nil
}

func (r *InMemoryDealRepository) Update(deal *domain.Deal) error {
	if _, exists := r.deals[deal.ID]; !exists {
		return errors.New("deal not found")
	}
	r.deals[deal.ID] = deal
	return nil
}

func (r *InMemoryDealRepository) Delete(id string) error {
	if _, exists := r.deals[id]; !exists {
		return errors.New("deal not found")
	}
	delete(r.deals, id)
	return nil
}

func (r *InMemoryDealRepository) FindActiveDeals() ([]domain.Deal, error) {
	activeDeals := make([]domain.Deal, 0)
	for _, deal := range r.deals {
		if deal.IsActive() {
			activeDeals = append(activeDeals, *deal)
		}
	}
	return activeDeals, nil
}

