package cart

import "time"

type CartItem struct {
	ID        int
	UserID    int
	ProductID int
	Quantity  int
	CreatedAt time.Time
}
