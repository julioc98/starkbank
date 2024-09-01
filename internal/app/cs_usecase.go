// Package app implements the application layer.
package app

import (
	"context"
	"log"

	language "cloud.google.com/go/language/apiv1"
	"cloud.google.com/go/language/apiv1/languagepb"
	"github.com/julioc98/starkbank/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Analyst, error)
	GetByID(id uint64) (*domain.Analyst, error)
	Create(a *domain.Analyst) error
	GetBySentimentAndSkill(sentiment, skill string) ([]domain.Analyst, error)
}

type UseCase struct {
	repo Repository
	l    *language.Client
}

func NewUseCase(repo Repository, l *language.Client) *UseCase {
	return &UseCase{
		repo: repo,
		l:    l,
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

func (uc *UseCase) GetBySentimentAndSkill(sentiment, skill string) ([]domain.Analyst, error) {
	return uc.repo.GetBySentimentAndSkill(sentiment, skill)
}

func (uc *UseCase) Analyze(msg *domain.Msg) (*domain.Msg, error) {
	log.Printf("Analyzing message: %s\n", msg.Content)

	ctx := context.Background()

	skill := msg.Skill

	if skill == "" {
		// Reverse the sentiment to find the right analyst.
		revSentiment := domain.Reverse(msg.Sentiment)

		analysts, err := uc.GetBySentimentAndSkill(revSentiment, skill)
		if err != nil {
			return nil, err
		}

		if len(analysts) == 0 {
			return nil, domain.ErrNoAnalystsFound
		}

		// Choose an analyst randomly.
		// For simplicity, we choose the first analyst.
		msg.Analyst = &analysts[0]
	}

	// Detects the sentiment of the text.
	sentiment, err := uc.l.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: msg.Content,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		return nil, err
	}

	emotion := "unknown"

	if sentiment.DocumentSentiment.Score >= 0.8 && sentiment.DocumentSentiment.Magnitude >= 0.8 {
		emotion = "positive"
	} else if sentiment.DocumentSentiment.Score >= 0.1 && sentiment.DocumentSentiment.Magnitude >= 0 {
		emotion = "neutral"
	} else if sentiment.DocumentSentiment.Score <= 0 && sentiment.DocumentSentiment.Magnitude <= 0.7 {
		emotion = "negative"
	} else if sentiment.DocumentSentiment.Score <= 0 && sentiment.DocumentSentiment.Magnitude >= 0.8 {
		emotion = "mixed"
	}

	msg.Sentiment = emotion

	return uc.Response(msg)
}

func (uc *UseCase) Response(msg *domain.Msg) (*domain.Msg, error) {
	// For simplicity, we return the same message.
	msg.Response = msg.Content

	return msg, nil
}
