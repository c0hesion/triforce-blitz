package routes

import (
	"github.com/c0hesion/httpx"
	"github.com/go-chi/chi/v5"
	"triforce-blitz/python"
	"triforce-blitz/randomizer"
	"triforce-blitz/util"
)

var (
	rando       = randomizer.New(`C:\usr\share\triforce-blitz\generators\v7.1.3-blitz-0.40`)
	interpreter = util.Must(python.FindInterpreter(&python.LocalFinder{}))
)

var (
	SeedsDir string
)

func Register(router *chi.Mux) {
	router.Post("/seeds", httpx.Handler[SeedInformation](generateSeed))
}
