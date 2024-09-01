// Package app implements the application layer.
package app

import "github.com/julioc98/starkbank/internal/domain"

type Repository interface {
	GetAll() ([]domain.Analyst, error)
	GetByID(id uint64) (*domain.Analyst, error)
	Create(a *domain.Analyst) error
}

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (uc *UseCase) GetAll() ([]domain.Analyst, error) {
	return uc.repo.GetAll()
}

func (uc *UseCase) GetByID(id uint64) (*domain.Analyst, error) {
	return uc.repo.GetByID(id)
}

func (uc *UseCase) Create(a *domain.Analyst) error {
	return uc.repo.Create(a)
}
