package user

import (
	"database/sql"
	"errors"
	"log"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	if db == nil {
		log.Fatal("db is nil!")
	}
	return &Repository{db: db}
}

func (r *Repository) Create(user *User) error {
	query := `
	INSERT INTO users (email, password, role_id) 
	VALUES ($1, $2, $3)
	RETURNING id, created_at
	`

	err := r.db.QueryRow(query, user.Email, user.Password, user.RoleID).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetByEmail(email string) (*User, error) {
	var user User
	query := `
	SELECT id, email, password, role_id, created_at
	FROM users
	WHERE email=$1
	`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.RoleID, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
