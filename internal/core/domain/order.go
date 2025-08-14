package domain

const zeroTotal = 0

type Order struct {
	ID          string
	UserID      string
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
	total := zeroTotal
	orderItems := o.OrderItems

	for _, orderItem := range orderItems {
		total += orderItem.Price.Amount
	}
	return total
}
