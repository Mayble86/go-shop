package category

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(name string) (*Category, error) {
	query := `
	INSERT INTO categories (name)
	VALUES ($1)
	RETURNING id
	`

	var c Category
	c.Name = name

	err := r.db.QueryRow(query, name).Scan(&c.ID)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *Repository) GetAll() ([]Category, error) {
	rows, err := r.db.Query(`SELECT id, name FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]Category, 0)

	for rows.Next() {
		var c Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}
