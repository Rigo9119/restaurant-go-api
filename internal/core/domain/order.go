package domain

type Order struct {
	ID          string
	CustomerID  string
	OrderItems  []MenuItem
	OrderTotal  int
	OrderStatus OrderStatus
}

func (o *Order) IsEmpty() bool {
	return true
}

func (o *Order) ItemCount() int {
	// TODO: necesita devolver el total de los items que tiene una orden
	return 1
}

func (o *Order) HasStatus(status OrderStatus) bool {
	// TODO: verifica el estado de la ordern
	return false
}

func (o *Order) CalculateSubtotal() int {
	// TODO: Calcual el subtotal de la orden
	return 1
}
