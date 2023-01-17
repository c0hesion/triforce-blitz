package randomizer

type Manager struct {
	installer     Installer
	InstallPath   string
	downloader    Downloader
	DownloadsPath string
}

// Downloader defines an interface for downloading Randomizer releases from a source.
type Downloader interface {
	// GetAvailableReleases will return a list of available Randomizer versions.
	GetAvailableReleases() ([]Version, error)

	// Download will download a Randomizer release to destination. If successful, it will return the full path of the
	// downloaded release.
	Download(version Version, destination string) (string, error)
}

// Installer defines an interface for dow
type Installer interface {
	// Supports checks if the file at name can be installed by this Installer.
	Supports(name string) bool

	// Install installs the source file to the destination.
	Install(destination string, source string) error
}
