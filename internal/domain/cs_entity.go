// Package domain represents the domain layer of the application.
package domain

import "time"

const Positive = "positive"
const Negative = "negative"
const Neutral = "neutral"

func Reverse(sentiment string) string {
	reverse := map[string]string{
		Positive: Negative,
		Negative: Positive,
		Neutral:  Neutral,
	}

	return reverse[sentiment]
}

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

// Msg is a chat message.
type Msg struct {
	Content   string    `json:"content,omitempty"`
	Sentiment string    `json:"sentiment,omitempty"`
	Skill     string    `json:"skill,omitempty"`
	Response  string    `json:"response,omitempty"`
	Analyst   *Analyst  `json:"analyst,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
