// Package app implements the application layer.
package app

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	language "cloud.google.com/go/language/apiv1"
	"cloud.google.com/go/language/apiv1/languagepb"
	"cloud.google.com/go/vertexai/genai"
	"github.com/julioc98/starkbank/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Analyst, error)
	GetByID(id uint64) (*domain.Analyst, error)
	Create(a *domain.Analyst) error
	GetBySentimentAndSkill(sentiment, skill string) ([]domain.Analyst, error)
}

type UseCase struct {
	ctx   context.Context
	repo  Repository
	l     *language.Client
	gen   *genai.Client
	model *genai.GenerativeModel
	chat  *genai.ChatSession
}

func NewUseCase(ctx context.Context, repo Repository, l *language.Client, g *genai.Client, m *genai.GenerativeModel) *UseCase {
	return &UseCase{
		ctx:   ctx,
		repo:  repo,
		l:     l,
		gen:   g,
		model: m,
		chat:  m.StartChat(),
	}
}

func (uc *UseCase) GetAll() ([]domain.Analyst, error) {
	return uc.repo.GetAll()
}

func (uc *UseCase) GetByID(id uint64) (*domain.Analyst, error) {
	return uc.repo.GetByID(id)
}

func (uc *UseCase) Create(a *domain.Analyst) error {
	log.Printf("Creating analyst: %s\n", a.Name)
	return uc.repo.Create(a)
}

func (uc *UseCase) GetBySentimentAndSkill(sentiment, skill string) ([]domain.Analyst, error) {
	return uc.repo.GetBySentimentAndSkill(sentiment, skill)
}

func (uc *UseCase) Analyze(msg *domain.Msg) (*domain.Msg, error) {
	log.Printf("Analyzing message: %s\n", msg.Content)

	// ctx := uc.ctx

	// Detects the sentiment of the text.
	sentiment, err := uc.l.AnalyzeSentiment(context.Background(), &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: msg.Content,
			},
			Type:     languagepb.Document_PLAIN_TEXT,
			Language: "pt-BR", // Brazilian Portuguese
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	if err != nil {
		return nil, err
	}

	emotion := "unknown"

	if sentiment.DocumentSentiment.Score >= 0.6 {
		emotion = "positive"
	} else if sentiment.DocumentSentiment.Score >= 0.1 {
		emotion = "neutral"
	} else if sentiment.DocumentSentiment.Score <= 0 {
		emotion = "negative"
	}

	msg.Sentiment = emotion

	log.Printf("Sentiment: %s\n", msg.Sentiment)

	skill := msg.Skill

	if skill == "" {
		// Find the first matching word in the message.
		skill = findingKey(findFirstMatchingWord(msg.Content))
		log.Printf("Skill: %s\n", skill)

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

	return uc.Response(msg)
}

func (uc *UseCase) Response(msg *domain.Msg) (*domain.Msg, error) {
	// ctx := uc.ctx

	// resp, err := uc.model.GenerateContent(ctx, genai.Text(msg.Content))
	// if err != nil {
	// 	return nil, err
	// }

	log.Printf("Sending message: %s\n", msg.Content)

	resp, err := uc.send(msg.Content)
	if err != nil {
		return nil, err
	}

	log.Printf("Response: %s\n", resp)

	msg.CreatedAt = time.Now()

	// msg.Response = genResponse(resp)

	msg.Response = resp

	return msg, nil
}

func genResponse(resp *genai.GenerateContentResponse) string {
	res := ""

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			res += fmt.Sprint(cand.Content.Parts)
		}
	}

	return res
}

func findFirstMatchingWord(text string) string {
	wordsInText := strings.Fields(strings.ToLower(text))

	words := []string{"api", "integração", "acesso", "antecipação", "pix", "boleto", "ted", "contrato", "negociação", "financeiro", "código", "preço", "negociação", "treinamento", "taxa"}

	for _, wordInText := range wordsInText {
		for _, word := range words {
			if strings.Contains(strings.ToLower(wordInText), strings.ToLower(word)) {
				return strings.ToLower(word)
			}
		}
	}

	return ""
}

func findingKey(word string) string {
	context := map[string][]string{
		"software":   {"api", "integração", "acesso", "código"},
		"venda":      {"contrato", "negociação", "preço", "treinamento"},
		"financeiro": {"antecipação", "pix", "boleto", "ted", "financeiro", "taxa"},
	}

	if word != "" {
		for key, value := range context {
			for _, palavra := range value {
				if (strings.ToLower(palavra)) == (strings.ToLower(word)) {
					return key
				}
			}
		}
	}
	return "general"
}

func (uc *UseCase) send(message string) (string, error) {
	// ctx := uc.ctx
	ctx := context.Background()
	r, err := uc.chat.SendMessage(ctx, genai.Text(message))
	if err != nil {
		log.Printf("SendMessage Error: %v\n", err)
		return "", err
	}

	log.Printf("Response: %v\n", r)

	// rb, err := json.MarshalIndent(r, "", "  ")
	// if err != nil {
	// 	return "", err
	// }

	return genResponse(r), nil
}
