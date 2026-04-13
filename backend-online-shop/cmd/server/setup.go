package main

import (
	"database/sql"

	"backend-online-shop/internal/product"
	"backend-online-shop/internal/user"
)

func setupUser(db *sql.DB) *user.Handler {
	repo := user.NewRepository(db)
	service := user.NewService(repo)
	return user.NewHandler(service)
}

func setupProduct(db *sql.DB) *product.Handler {
	repo := product.NewRepository(db)
	service := product.NewService(repo)
	return product.NewHandler(service)
}
