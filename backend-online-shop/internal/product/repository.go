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
		SELECT 
			p.id, p.name, p.description, p.price, p.category_id, p.created_at,
			c.id, c.name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
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
			&p.Category.ID,
			&p.Category.Name,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
