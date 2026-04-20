package order

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateOrder(userID int) (int, error) {
	var id int

	err := r.db.QueryRow(`
		INSERT INTO orders (user_id, status)
		VALUES ($1, 'pending')
		RETURNING id
	`, userID).Scan(&id)

	return id, err
}

func (r *Repository) AddItem(item OrderItem) error {
	_, err := r.db.Exec(`
		INSERT INTO order_items (order_id, product_id, quantity, price)
		VALUES ($1, $2, $3, $4)
	`, item.OrderID, item.ProductID, item.Quantity, item.Price)

	return err
}

func (r *Repository) GetCart(userID int) ([]OrderItem, error) {
	rows, err := r.db.Query(`
		SELECT p.id, ci.quantity, p.price
		FROM cart_items ci
		JOIN products p ON p.id = ci.product_id
		WHERE ci.user_id = $1
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []OrderItem

	for rows.Next() {
		var i OrderItem
		err := rows.Scan(&i.ProductID, &i.Quantity, &i.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	return items, nil
}

func (r *Repository) ClearCart(userID int) error {
	_, err := r.db.Exec(`
		DELETE FROM cart_items WHERE user_id = $1
	`, userID)

	return err
}
