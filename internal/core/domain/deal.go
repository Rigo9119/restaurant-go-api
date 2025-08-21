package domain

import (
	"restaurant-go-api/internal/shared/utils"
	"time"
)

type Deal struct {
	ID              string
	Name            string
	Description     string
	DiscountPercent int
	StartDate       time.Time
	EndDate         time.Time
	Items           []MenuItem
	DealPrice       Price
}

func NewDeal(
	name, description string,
	items []MenuItem,
	dealPrice Price,
	startDate, endDate time.Time,
) *Deal {
	return &Deal{
		ID:          utils.GenerateRandomID("deal"),
		Name:        name,
		Description: description,
		Items:       items,
		DealPrice:   dealPrice,
		StartDate:   startDate,
		EndDate:     endDate,
	}
}

func (d *Deal) CalculateDealSubTotal() int {
	return d.DealPrice.Amount
}

func (d *Deal) IsActive() bool {
	now := time.Now()
	return now.After(d.StartDate) && now.Before(d.EndDate)
}
