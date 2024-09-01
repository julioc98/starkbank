// Package domain represents the domain layer of the application.
package domain

import "time"

// Analyst is a CS analyst.
type Analyst struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Skill     string    `json:"skill,omitempty"`
	Sentiment string    `json:"sentiment,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
