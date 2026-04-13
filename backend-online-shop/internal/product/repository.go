package product

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(p *Product) error {
	query := `
	INSERT INTO products (name, description, price, category_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at
	`

	return r.db.QueryRow(
		query,
		p.Name,
		p.Description,
		p.Price,
		p.CategoryID,
	).Scan(&p.ID, &p.CreatedAt)
}

func (r *Repository) GetAll() ([]Product, error) {
	rows, err := r.db.Query(`
	SELECT id, name, description, price, category_id, created_at
	FROM products
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.CategoryID,
			&p.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
