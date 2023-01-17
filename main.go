package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
	"net/http"
	"triforce-blitz/randomizer"
	"triforce-blitz/routes"
	"triforce-blitz/util"
)

var (
	r = chi.NewRouter()
)

func main() {
	LoadConfiguration()
	dl := randomizer.NewGithubDownloader("Elagatua", "OoT-Randomizer")
	versions, err := dl.GetAvailableReleases()
	if err == nil {
		slog.Info("The following versions are available:")
		for _, v := range util.Filter(versions, func(v randomizer.Version) bool { return v.Branch == "blitz" }) {
			slog.Info(fmt.Sprintf("\t%v", v.String()))
		}
	}
	// create router, register routes and listen
	r.Use(middleware.Logger)
	routes.Register(r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error("http failed", err)
	}
}
