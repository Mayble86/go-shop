package product

import (
	"time"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CategoryID  int       `json:"-"`
	Category    Category  `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
}
