package randomizer

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"triforce-blitz/python"
)

type Randomizer struct {
	path string

	// Stdout will receive the Randomizer's output.
	Stdout io.Writer

	// Stderr will receive the Randomizer's error output.
	Stderr io.Writer
}

func New(path string) *Randomizer {
	return &Randomizer{
		path: path,
	}
}

func (r *Randomizer) entrypoint() string {
	return filepath.Join(r.path, "OoTRandomizer.py")
}

func (r *Randomizer) writeSettingsFile(name string, settings *Settings) error {
	b, err := json.MarshalIndent(&settings, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(name, b, 0644)
}

func (r *Randomizer) Generate(interpreter python.Interpreter, settings *Settings) (err error) {
	if settings.OutputPath == "" {
		if settings.OutputPath, err = os.Getwd(); err != nil {
			return err
		}
	}
	// create the output directory if needed
	if err = os.MkdirAll(settings.OutputPath, 0644); err != nil {
		return fmt.Errorf("could not create output directory: %v", err)
	}
	// write out the settings struct
	settingsPath := filepath.Join(settings.OutputPath, "settings.json")
	if err = r.writeSettingsFile(settingsPath, settings); err != nil {
		return fmt.Errorf("could not write settings file: %v", err)
	}
	// Invoke the Interpreter
	return interpreter.RunScript(
		r.entrypoint(), r.Stdout, r.Stderr,
		"--settings", settingsPath,
		"--settings_preset", settings.Preset,
	)
}
