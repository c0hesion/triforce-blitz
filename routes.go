package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
	"time"
	"triforce-blitz/randomizer"
)

type GenerateSeedResponse struct {
	Id                string    `json:"id"`
	Hash              []string  `json:"hash"`
	RequestedAt       time.Time `json:"requestedAt"`
	GeneratedAt       time.Time `json:"generatedAt"`
	RandomizerVersion string    `json:"randomizerVersion"`
}

func generateSeed(w http.ResponseWriter, r *http.Request) (*GenerateSeedResponse, int, error) {
	body := &GenerateSeedResponse{
		Id:          uuid.NewString(),
		RequestedAt: time.Now(),
	}
	settings := randomizer.DefaultSettings()
	settings.OutputPath = filepath.Join(`C:\usr\share\triforce-blitz\seeds`, body.Id)
	if err := rando.Generate(interpreter, settings); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	body.GeneratedAt = time.Now()
	// read the spoiler log for our hash and such
	spoilerFilename := fmt.Sprintf("%s_Spoiler.json", settings.OutputName)
	spoiler, err := randomizer.ReadSpoilerLog(filepath.Join(settings.OutputPath, spoilerFilename))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	body.Hash = spoiler.Hash
	body.RandomizerVersion = spoiler.Version
	return body, http.StatusOK, nil
}
