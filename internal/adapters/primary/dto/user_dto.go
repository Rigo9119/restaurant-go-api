package dto

// User Request DTOs - what clients send to the server
//  ✅ Use Pointers (*string, *int) for:
//- Update requests - To distinguish between "don't change" vs "set to empty"
// ❌ Don't Use Pointers for:
// - Create requests - All fields are required/explicit

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	Role string `json:"role" validate:"required,oneof=admin customer manager"`
}

type UpdateUserRequest struct {
	Name *string `json:"name" validate:"required,min=2,max=100"`
	Role *string `json:"role" validate:"required,oneof=admin customer manager"`
}

// User Response DTOs - what server sends to clients

type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
	Count int            `json:"count"`
}
