package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/julioc98/starkbank/internal/app"
	"github.com/julioc98/starkbank/internal/infra/api"
	"github.com/julioc98/starkbank/internal/infra/db"
	"github.com/julioc98/starkbank/pkg/database"
	_ "github.com/lib/pq"
)

func main() {

	conn, err := database.Conn()
	if err != nil {
		log.Fatal(err)
	}

	defer func() { _ = conn.Close() }()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/"))
	r.Use(middleware.AllowContentType("application/json", "text/xml"))

	repo := db.NewAnalystPostgresRepository(conn)
	uc := app.NewUseCase(repo)
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
