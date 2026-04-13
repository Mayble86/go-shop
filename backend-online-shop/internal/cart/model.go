package cart

type Cart struct {
	ID     int
	UserID int
}

type CartItem struct {
	ID        int
	CartID    int
	ProductID int
	Quantity  int
}
