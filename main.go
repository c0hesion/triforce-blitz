package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
	"net/http"
	"triforce-blitz/routes"
)

var (
	r = chi.NewRouter()
)

func main() {
	LoadConfiguration()
	// create router, register routes and listen
	r.Use(middleware.Logger)
	routes.Register(r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error("http failed", err)
	}
}
