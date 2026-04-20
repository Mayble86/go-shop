package cart

import (
	"backend-online-shop/pkg/auth"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

type AddRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	var req AddRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	userID := auth.GetUserID(r)

	err := h.service.Add(userID, req.ProductID, req.Quantity)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserID(r)

	items, err := h.service.Get(userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(items)
}
