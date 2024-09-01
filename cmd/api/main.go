package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/google/generative-ai-go/genai"
	"github.com/julioc98/starkbank/internal/app"
	"github.com/julioc98/starkbank/internal/infra/api"
	"github.com/julioc98/starkbank/internal/infra/db"
	"github.com/julioc98/starkbank/pkg/database"
	_ "github.com/lib/pq"
	"google.golang.org/api/option"

	language "cloud.google.com/go/language/apiv1"
)

func main() {

	ctx := context.Background()

	// Connect to the database.
	conn, err := database.Conn()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// Creates a client.
	langClient, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer langClient.Close()

	genaiClient, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer genaiClient.Close()

	model := genaiClient.GenerativeModel("gemini-1.5-flash")
	model.SetMaxOutputTokens(50)
	model.SystemInstruction = genai.NewUserContent(genai.Text("Voce é um Customer Success Analyst do StarkBank com acesso ao knowledge base e FAQ, responda como estivesse pesquisado em uma base de dados e sempre seja solicito e com pedidos de desculpa quando necessário."))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/"))
	r.Use(middleware.AllowContentType("application/json", "text/xml"))
	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permite todos os domínios
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Max cache duration in seconds
	}))

	repo := db.NewAnalystPostgresRepository(conn)
	uc := app.NewUseCase(repo, langClient, genaiClient, model)
	h := api.NewRestHandler(r, uc)

	h.RegisterHandlers()

	http.Handle("/", r)

	// Start server.
	log.Println("Starting server on port 3000...")

	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
