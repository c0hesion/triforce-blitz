package main

import (
	"github.com/c0hesion/httpx"
	"github.com/google/uuid"
	"path/filepath"
	"time"
	"triforce-blitz/randomizer"
)

type SeedInformation struct {
	Id                string    `json:"id"`
	Hash              []string  `json:"hash"`
	RequestedAt       time.Time `json:"requestedAt"`
	GeneratedAt       time.Time `json:"generatedAt"`
	RandomizerVersion string    `json:"randomizerVersion"`
}

func generateSeed(r *httpx.Request[any]) *httpx.Response {
	id := uuid.NewString()
	requestedAt := time.Now()
	settings := randomizer.DefaultSettings()
	settings.OutputPath = filepath.Join(`C:\usr\share\triforce-blitz\seeds`, id)
	if err := rando.Generate(interpreter, settings); err != nil {
		return httpx.InternalServerError(err)
	}
	generatedAt := time.Now()
	spoiler, err := randomizer.ReadSpoilerLog(settings.SpoilerPath())
	if err != nil {
		return httpx.InternalServerError(err)
	}
	return httpx.OkWithBody(&SeedInformation{
		Id:                id,
		Hash:              spoiler.Hash,
		RequestedAt:       requestedAt,
		GeneratedAt:       generatedAt,
		RandomizerVersion: spoiler.Version,
	})
}
