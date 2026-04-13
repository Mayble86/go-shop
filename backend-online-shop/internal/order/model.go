package order

import (
	"time"
)

type Order struct {
	ID        int
	UserID    int
	Total     float64
	Status    string
	CreatedAt time.Time
}
