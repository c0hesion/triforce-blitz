package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
	"net/http"
	"triforce-blitz/python"
	"triforce-blitz/randomizer"
	"triforce-blitz/util"
)

var (
	router      = chi.NewRouter()
	rando       = randomizer.New(`C:\usr\share\triforce-blitz\generators\v7.1.3-blitz-0.40`)
	interpreter = util.Must(python.FindInterpreter(&python.LocalFinder{}))
)

func main() {
	version, err := interpreter.Version()
	if err != nil {
		slog.Error("failed to get Python version", err)
	}
	slog.Info("Python version: " + version)
	router.Use(middleware.Logger)
	registerHttpRoutes()
	if err := http.ListenAndServe(":8080", router); err != nil {
		slog.Error("http failed", err)
	}
}

func registerHttpRoutes() {
	router.Post("/seeds", util.ApiHandler(generateSeed))
}
