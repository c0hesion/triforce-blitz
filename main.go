package main

import (
	"golang.org/x/exp/slog"
	"path/filepath"
	"triforce-blitz/python"
	"triforce-blitz/randomizer"
	"triforce-blitz/util"
)

var (
	interpreter = util.Must(python.FindInterpreter(&python.LocalFinder{}))
)

func main() {
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
