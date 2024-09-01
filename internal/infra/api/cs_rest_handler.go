// Package api implements the API layer.
package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/julioc98/starkbank/internal/domain"
)

// CreateAnalyst an new domain.Analyst.
func (h *RestHandler) CreateAnalyst(w http.ResponseWriter, r *http.Request) {
	var analyst domain.Analyst

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&analyst); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err := h.uc.Create(&analyst); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAllAnalysts gets all domain.Analyst.
func (h *RestHandler) GetAllAnalysts(w http.ResponseWriter, r *http.Request) {
	analysts, err := h.uc.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(analysts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetAnalystByID gets a domain.Analyst by ID.
func (h *RestHandler) GetAnalystByID(w http.ResponseWriter, r *http.Request) {
	analystID := chi.URLParam(r, "id")

	analystIDUint64, err := strconv.ParseUint(analystID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid analyst ID", http.StatusBadRequest)

		return
	}

	analyst, err := h.uc.GetByID(analystIDUint64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if analyst == nil {
		http.Error(w, "Analyst not found", http.StatusNotFound)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(analyst); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
