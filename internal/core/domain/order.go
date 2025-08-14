package domain

const ZERO_TOTAL = 0

type Order struct {
	ID          string
	CustomerID  string
	OrderItems  []MenuItem
	OrderTotal  int
	OrderStatus OrderStatus
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
	total := ZERO_TOTAL
	orderItems := o.OrderItems

	for _, orderItem := range orderItems {
		total += orderItem.Price.Amount
	}
	return total
}
