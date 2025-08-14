package domain

type Order struct {
	ID          string
	CustomerID  string
	OrderItems  []MenuItem
	OrderTotal  int
	OrderStatus OrderStatus
}

func GetTotal() {
	// ayuda a calcular el total de la orden
}

func AddItem() {
	// ayuda a añadir una item a la orden
}

func ChangeItem() {
	// ayuda a cambiar items de la orden
}
