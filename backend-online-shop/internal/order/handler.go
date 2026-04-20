package order

import (
	"encoding/json"
	"net/http"

	"backend-online-shop/pkg/auth"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserID(r)

	orderID, err := h.service.Checkout(userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"order_id": orderID,
		"status":   "created",
	})
}
