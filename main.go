package main

import (
	"golang.org/x/exp/slog"
	"path/filepath"
	"triforce-blitz/python"
	"triforce-blitz/randomizer"
)

func main() {
	interpreter, err := python.FindInterpreter(&python.LocalFinder{})
	if err != nil {
		slog.Error("failed to find Python", err)
	}
	version, err := interpreter.Version()
	if err != nil {
		slog.Error("failed to get Python version", err)
	}
	slog.Info("Python version: " + version)
	// attempt generating a TFB seed
	r := randomizer.New(`C:\usr\share\triforce-blitz\generators\v7.1.3-blitz-0.40`)
	settings := randomizer.DefaultSettings()
	settings.OutputPath = filepath.Join(`C:\usr\share\triforce-blitz\seeds`, settings.Seed)
	if err := r.Generate(interpreter, settings); err != nil {
		slog.Error("could not generate seed", err)
	}
}
