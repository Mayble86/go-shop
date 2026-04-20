package cart

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Add(userID, productID, quantity int) error {
	query := `
	INSERT INTO cart_items (user_id, product_id, quantity)
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id, product_id)
	DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity
	`

	_, err := r.db.Exec(query, userID, productID, quantity)
	return err
}

func (r *Repository) GetByUser(userID int) ([]CartItem, error) {
	rows, err := r.db.Query(`
	SELECT id, user_id, product_id, quantity, created_at
	FROM cart_items
	WHERE user_id = $1
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []CartItem

	for rows.Next() {
		var c CartItem
		err := rows.Scan(&c.ID, &c.UserID, &c.ProductID, &c.Quantity, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, c)
	}

	return items, nil
}
