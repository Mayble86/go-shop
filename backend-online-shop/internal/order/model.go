package order

import (
	"time"
)

type Order struct {
	ID        int
	UserID    int
	Status    string
	CreatedAt time.Time
}

type OrderItem struct {
	ID        int
	OrderID   int
	ProductID int
	Quantity  int
	Price     float64
}
