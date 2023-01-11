package main

import (
	"golang.org/x/exp/slog"
	"triforce-blitz/python"
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
}
