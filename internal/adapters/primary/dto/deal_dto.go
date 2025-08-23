package dto

import "time"

// Request DTOs for creating deals
type CreateDealRequest struct {
	Name            string    `json:"name" validate:"required"`
	Description     string    `json:"description"`
	DiscountPercent int       `json:"discount_percent" validate:"min=0,max=100"`
	StartDate       time.Time `json:"start_date" validate:"required"`
	EndDate         time.Time `json:"end_date" validate:"required"`
	Items           []string  `json:"items" validate:"required"` // Array of MenuItem IDs
	DealPrice       PriceDTO  `json:"deal_price" validate:"required"`
}

type UpdateDealRequest struct {
	Name            *string    `json:"name,omitempty"`
	Description     *string    `json:"description,omitempty"`
	DiscountPercent *int       `json:"discount_percent,omitempty" validate:"omitempty,min=0,max=100"`
	StartDate       *time.Time `json:"start_date,omitempty"`
	EndDate         *time.Time `json:"end_date,omitempty"`
	DealPrice       *PriceDTO  `json:"deal_price,omitempty"`
}

// Response DTOs
type DealResponse struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	DiscountPercent int       `json:"discount_percent"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	Items           []ItemDTO `json:"items"`
	DealPrice       PriceDTO  `json:"deal_price"`
	IsActive        bool      `json:"is_active"`
}

type PriceDTO struct {
	Amount   int    `json:"amount" validate:"required,min=0"`
	Currency string `json:"currency" validate:"required"`
}

type DealsListResponse struct {
	Deals []DealResponse `json:"deals"`
	Total int            `json:"total"`
}