package randomizer

import (
	"fmt"
	"github.com/google/uuid"
	"path/filepath"
)

type Settings struct {
	// OutputPath is the path to the directory the Randomizer will output any resulting files to. If an empty string,
	// this will default to the process's working directory.
	OutputPath string `json:"output_dir,omitempty"`

	// Seed is the seed used to initialize the Randomizer's random number generator. If an empty string, the
	// Randomizer will generate a random seed value.
	Seed string `json:"seed,omitempty"`

	// Preset is a setting preset that will toggle a collection of pre-defined settings in the Randomizer. Please see
	// the Randomizer documentation for more information.
	Preset string `json:"preset,omitempty"`

	// RomPath is the path to a compressed or uncompressed Ocarina of Time ROM that will be used as the base to generate
	// the patch file form. If an empty string, the Randomizer will try to look for a ROM itself.
	RomPath string `json:"rom,omitempty"`

	// OutputName is the name that will be used or prefixed to all generated files. If left empty, the Randomizer will
	// choose a predefined name.
	OutputName string `json:"output_file,omitempty"`

	// CreatePatchFile toggles the generation of a ZPF patch file that can be used to patch a ROM.
	CreatePatchFile bool `json:"create_patch_file"`

	// CreateCompressedRom toggles the creation of a compressed ROM. This feature is only available on select platforms,
	// refer to the Randomizer documentation for more information.
	CreateCompressedRom bool `json:"create_compressed_rom"`

	// CreateUncompressedRom toggles the creation of an uncompressed ROM.
	CreateUncompressedRom bool `json:"create_uncompressed_rom"`

	// CreateCosmeticsLog toggles the creation of a cosmetics log.
	CreateCosmeticsLog bool `json:"create_cosmetics_log"`
}

func (s *Settings) SpoilerFilename() string {
	return fmt.Sprintf("%s_Spoiler.json", s.OutputName)
}

func (s *Settings) SpoilerPath() string {
	return filepath.Join(s.OutputPath, s.SpoilerFilename())
}

func (s *Settings) PatchFilename() string {
	return fmt.Sprintf("%s.zpf", s.OutputName)
}

func (s *Settings) PatchPath() string {
	return filepath.Join(s.OutputPath, s.PatchFilename())
}

func DefaultSettings() *Settings {
	return &Settings{
		Seed:                  uuid.New().String(),
		OutputName:            "TriforceBlitz",
		Preset:                "Triforce Blitz",
		CreatePatchFile:       true,
		CreateCompressedRom:   false,
		CreateUncompressedRom: false,
		CreateCosmeticsLog:    false,
	}
}
