package main

import (
	"backend-online-shop/internal/cart"
	"backend-online-shop/internal/order"
	"database/sql"

	"backend-online-shop/internal/category"
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

func setupCategory(db *sql.DB) *category.Handler {
	repo := category.NewRepository(db)
	service := category.NewService(repo)
	return category.NewHandler(service)
}

func setupCart(db *sql.DB) *cart.Handler {
	repo := cart.NewRepository(db)
	service := cart.NewService(repo)
	return cart.NewHandler(service)
}

func setupOrder(db *sql.DB) *order.Handler {
	orderRepo := order.NewRepository(db)
	orderService := order.NewService(orderRepo)
	return order.NewHandler(orderService)
}
