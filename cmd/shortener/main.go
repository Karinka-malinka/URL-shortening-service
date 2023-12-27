package main

import (
	"fmt"
	"net/http"

	"github.com/URL-shortening-service/internal/config"
	"github.com/URL-shortening-service/internal/handlers"
	"github.com/go-chi/chi/v5"
)

var cfg *config.ConfigData

func init() {
	cfg = config.NewConfig()
}

func main() {

	parseFlags(cfg)

	if err := run(); err != nil {
		panic(err)
	}

}

func run() error {

	r := chi.NewRouter()

	r.Get("/{id}", handlers.ResolveURL)
	r.Post("/", handlers.ShorteningURL(cfg.BaseAddr))

	fmt.Println("Running server on", cfg.RunAddr)
	return http.ListenAndServe(cfg.RunAddr, r)
}
