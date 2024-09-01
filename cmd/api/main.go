package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/julioc98/starkbank/internal/app"
	"github.com/julioc98/starkbank/internal/infra/api"
	"github.com/julioc98/starkbank/internal/infra/db"
	"github.com/julioc98/starkbank/pkg/database"
	_ "github.com/lib/pq"

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
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

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
	uc := app.NewUseCase(repo, client)
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
