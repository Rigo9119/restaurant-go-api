package domain

import "restaurant-go-api/internal/shared/utils"

const zeroTotal = 0

type Order struct {
	ID          string
	UserID      string
	OrderItems  []MenuItem
	OrderTotal  int
	OrderStatus OrderStatus
}

type OrderBuilder struct {
	userID      string
	orderItems  []MenuItem
	orderStatus OrderStatus
}

func NewOrderBuilder(userID string) *OrderBuilder {
	return &OrderBuilder{
		userID:      userID,
		orderItems:  []MenuItem{},
		orderStatus: Ordered,
	}
}

func (b *OrderBuilder) AddItem(item MenuItem) *OrderBuilder {
	b.orderItems = append(b.orderItems, item)
	return b
}

func (b *OrderBuilder) AddItems(items []MenuItem) *OrderBuilder {
	b.orderItems = append(b.orderItems, items...)
	return b
}

func (b *OrderBuilder) SetStatus(status OrderStatus) *OrderBuilder {
	b.orderStatus = status
	return b
}

func (b *OrderBuilder) Build() *Order {
	order := &Order{
		ID:          utils.GenerateRandomID("order"),
		UserID:      b.userID,
		OrderItems:  b.orderItems,
		OrderTotal:  0,
		OrderStatus: b.orderStatus,
	}
	order.OrderTotal = order.CalculateSubtotal()
	return order
}

func NewOrder(userID string) *Order {
	return NewOrderBuilder(userID).Build()
}

func (o *Order) IsEmpty() bool {
	return len(o.OrderItems) == 0
}

func (o *Order) ItemCount() int {
	return len(o.OrderItems)
}

func (o *Order) HasStatus(status OrderStatus) bool {
	return o.OrderStatus == status
}

func (o *Order) CalculateSubtotal() int {
	total := zeroTotal
	orderItems := o.OrderItems

	for _, orderItem := range orderItems {
		total += orderItem.Price.Amount
	}
	return total
}
