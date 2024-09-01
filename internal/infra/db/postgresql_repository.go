// Package db implements the database layer.
package db

import (
	"database/sql"

	"github.com/julioc98/starkbank/internal/domain"
)

// AnalystPostgresRepository represents a repository for CS analysts.
type AnalystPostgresRepository struct {
	db *sql.DB
}

// NewAnalystPostgresRepository creates a new AnalystPostgresRepository.
func NewAnalystPostgresRepository(db *sql.DB) *AnalystPostgresRepository {
	return &AnalystPostgresRepository{db: db}
}

// GetAll gets all analysts.
func (repo *AnalystPostgresRepository) GetAll() ([]domain.Analyst, error) {
	query := "SELECT id, name, email, skill, sentiment, created_at, updated_at FROM analysts"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()
	var analysts []domain.Analyst

	for rows.Next() {
		var a domain.Analyst
		if err := rows.Scan(&a.ID, &a.Name, &a.Email, &a.Skill, &a.Sentiment, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}

		analysts = append(analysts, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return analysts, nil
}

// GetByID gets an analyst by ID.
func (repo *AnalystPostgresRepository) GetByID(id uint64) (*domain.Analyst, error) {
	query := "SELECT id, name, email, skill, sentiment, created_at, updated_at FROM analysts WHERE id = $1"
	row := repo.db.QueryRow(query, id)

	var a domain.Analyst
	if err := row.Scan(&a.ID, &a.Name, &a.Email, &a.Skill, &a.Sentiment, &a.CreatedAt, &a.UpdatedAt); err != nil {
		return nil, err
	}

	return &a, nil
}

// Create creates a new analyst.
func (repo *AnalystPostgresRepository) Create(a *domain.Analyst) error {
	query := "INSERT INTO analysts (name, email, skill, sentiment) VALUES ($1, $2, $3, $4)"
	_, err := repo.db.Exec(query, a.Name, a.Email, a.Skill, a.Sentiment)
	if err != nil {
		return err
	}

	return nil
}

// GetBySentimentAndSkill gets analysts by sentiment and skill.
func (repo *AnalystPostgresRepository) GetBySentimentAndSkill(sentiment, skill string) ([]domain.Analyst, error) {
	query := "SELECT id, name, email, skill, sentiment, created_at, updated_at FROM analysts WHERE sentiment = $1 AND skill = $2"
	rows, err := repo.db.Query(query, sentiment, skill)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()
	var analysts []domain.Analyst

	for rows.Next() {
		var a domain.Analyst
		if err := rows.Scan(&a.ID, &a.Name, &a.Email, &a.Skill, &a.Sentiment, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}

		analysts = append(analysts, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return analysts, nil
}
