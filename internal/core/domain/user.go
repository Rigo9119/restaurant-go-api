package domain

import "restaurant-go-api/internal/shared/utils"

// Core -> contiene todas las entidades y reglas del negocio / proyecto
type User struct {
	ID   string
	Name string
	Role Role
}

func NewUser(role Role, name string) *User {
	switch role {
	case UserAdmin:
		return &User{
			ID:   utils.GenerateRandomID("user"),
			Name: name,
			Role: UserAdmin,
		}
	case UserCustomer:
		return &User{
			ID:   utils.GenerateRandomID("user"),
			Name: name,
			Role: UserCustomer,
		}
	case UserManager:
		return &User{
			ID:   utils.GenerateRandomID("user"),
			Name: name,
			Role: UserManager,
		}
	default:
		return nil
	}
}

func (u *User) GetPermissions() []string {
	switch u.Role {
	case UserCustomer:
		return []string{
			"create_order", "view_menu", "modify_own_order",
		}
	case UserManager:
		return []string{
			"manage_menu", "view_orders", "update_prices",
		}
	case UserAdmin:
		return []string{
			"manage_menu", "manage_users", "view_all_orders", "system_admin",
		}
	default:
		return []string{}
	}
}
