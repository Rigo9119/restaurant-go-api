package domain

// Core -> contiene todas las entidades y reglas del negocio / proyecto

type Customer struct {
	ID     string
	Name   string
	Orders []Order
}
