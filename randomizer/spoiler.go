package randomizer

import (
	"encoding/json"
	"os"
)

type SpoilerLog struct {
	Version  string   `json:":version"`
	Seed     string   `json:":seed"`
	Settings string   `json:":settings_string"`
	Hash     []string `json:"file_hash"`
}

func ReadSpoilerLog(name string) (*SpoilerLog, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	spoiler := &SpoilerLog{}
	if err := json.Unmarshal(b, spoiler); err != nil {
		return nil, err
	}
	return spoiler, nil
}
