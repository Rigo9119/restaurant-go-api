// Package domain
// contiene las entidades y reglas del negocio
package domain

import "restaurant-go-api/internal/shared/utils"

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

func (u *User) GetPermissions() []Permission {
	switch u.Role {
	case UserCustomer:
		return []Permission{
			PermissionCreateOrder, PermissionViewMenu, PermissionModifyOwnOrder,
		}
	case UserManager:
		return []Permission{
			PermissionManageMenu, PermissionViewOrders, PermissionUpdatePrices,
		}
	case UserAdmin:
		return []Permission{
			PermissionManageMenu, PermissionManageUsers, PermissionViewAllOrders, PermissionSystemAdmin,
		}
	default:
		return []Permission{}
	}
}
