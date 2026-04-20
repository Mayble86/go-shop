package main

import (
	"backend-online-shop/pkg/auth"
	"log"
	"net/http"

	"backend-online-shop/pkg/database"
)

func main() {
	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbURL := "postgres://admin:admin@localhost:15432/goshopdb?sslmode=disable"
	database.RunMigrations(dbURL)

	userHandler := setupUser(db)
	productHandler := setupProduct(db)
	categoryHandler := setupCategory(db)
	cartHandler := setupCart(db)
	orderHandler := setupOrder(db)

	http.HandleFunc("/orders/checkout", auth.Middleware(orderHandler.Checkout))

	http.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			auth.Middleware(cartHandler.Add)(w, r)
		case http.MethodGet:
			auth.Middleware(cartHandler.Get)(w, r)
		default:
			http.Error(w, "method not allowed", 405)
		}
	})

	http.HandleFunc("/auth/login", userHandler.Login)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/auth/register", userHandler.Register)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			auth.Middleware(userHandler.GetAllUsers)(w, r)
			return
		}

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			productHandler.CreateProduct(w, r)
			return
		}
		if r.Method == http.MethodGet {
			productHandler.GetProducts(w, r)
			return
		}
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			categoryHandler.CreateCategory(w, r)
		case http.MethodGet:
			categoryHandler.GetCategories(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started on :8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
