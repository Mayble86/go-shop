package main

import (
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

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/auth/register", userHandler.Register)

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

	log.Println("Server started on :8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
