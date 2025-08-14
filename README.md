Suggested Structure:

internal/
├── core/
│ ├── domain/
│ │ ├── customer.go
│ │ ├── menu.go
│ │ ├── order.go
│ │ └── enums.go
│ ├── services/
│ │ ├── menu_service.go
│ │ ├── order_service.go
│ │ └── customer_service.go
│ └── ports/
│ ├── menu_repository.go
│ └── order_repository.go
└── adapters/
├── primary/
│ └── http_handlers.go
└── secondary/
├── memory_repository.go
└── database_repository.go

What Goes Where:

domain/enums.go - All your enums (Size, PattyType, BunType,
etc.)

domain/menu.go - Menu entities and domain logic:
type MenuItem struct { ... }
type Burguer struct { ... }
type Side struct { ... }
type Drink struct { ... }

domain/customer.go - Customer entity and methods

domain/order.go - Order entity and business rules

services/ - Use cases and business orchestration

ports/ - Interface definitions for repositories

This approach gives you:

- Clear separation of concerns
- Easier to find specific domain logic
- Better testability
- Follows Domain-Driven Design principles
- Scales better as your application grows
