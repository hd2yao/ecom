package types

import "time"

type UserStore interface {
    GetUserByEmail(email string) (*User, error)
    GetUserByID(id int) (*User, error)
    CreateUser(User) error
}

type ProductStore interface {
    GetProducts() ([]Product, error)
}

type OrderStore interface {
    CreateOrder(Order) (int, error)
    CreateOrderItem(OrderItem) error
}

type User struct {
    ID        int       `json:"id"`
    FirstName string    `json:"first_name"`
    LastName  string    `json:"last_name"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    CreatedAt time.Time `json:"created_at"`
}

type Product struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Image       string  `json:"image"`
    Price       float64 `json:"price"`
    // note that this isn't the best way to handle quantity
    // because it's not atomic (in ACID), but it's good enough for this example
    Quantity  int       `json:"quantity"`
    CreatedAt time.Time `json:"created_at"`
}

type Order struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    Total     float64   `json:"total"`
    Status    string    `json:"status"`
    Address   string    `json:"address"`
    CreatedAt time.Time `json:"created_at"`
}

type OrderItem struct {
    ID        int       `json:"id"`
    OrderID   int       `json:"order_id"`
    ProductID int       `json:"product_id"`
    Quantity  int       `json:"quantity"`
    Price     float64   `json:"price"`
    CreatedAt time.Time `json:"created_at"`
}

type CartCheckoutItem struct {
    ProductID int `json:"product_id"`
    Quantity  int `json:"quantity"`
}

type RegisterUserPayload struct {
    FirstName string `json:"first_name" validate:"required"`
    LastName  string `json:"last_name" validate:"required"`
    Email     string `json:"email" validate:"required,email"`
    Password  string `json:"password" validate:"required,min=3,max=13"`
}

type LoginUserPayload struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

type CartCheckoutPayload struct {
    Items []CartCheckoutItem `json:"items" validate:"required"`
}
