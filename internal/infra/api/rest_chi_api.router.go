package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/julioc98/starkbank/internal/domain"
)

// UseCase represents a use case for CS analyst.
type UseCase interface {
	GetAll() ([]domain.Analyst, error)
	GetByID(id uint64) (*domain.Analyst, error)
	Create(a *domain.Analyst) error
	Analyze(msg *domain.Msg) (*domain.Msg, error)
}

// RestHandler represents a REST handler for  drivers.
type RestHandler struct {
	r  *chi.Mux
	uc UseCase
}

// NewRestHandler creates a new RestHandler.
func NewRestHandler(r *chi.Mux, uc UseCase) *RestHandler {
	return &RestHandler{
		r:  r,
		uc: uc,
	}
}

// RegisterHandlers registers the handlers of the REST API.
func (h *RestHandler) RegisterHandlers() {
	h.r.Post("/analysts", h.CreateAnalyst)
	h.r.Get("/analysts", h.GetAllAnalysts)
	h.r.Get("/analysts/{id}", h.GetAnalystByID)
	h.r.Post("/msg", h.ResponseMsg)
}
