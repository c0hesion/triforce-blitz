package randomizer

import (
	"context"
	"fmt"
	"github.com/google/go-github/v49/github"
	"golang.org/x/exp/slog"
)

type GithubDownloader struct {
	client         *github.Client
	ownerName      string
	repositoryName string
}

func NewGithubDownloader(owner string, repository string) *GithubDownloader {
	return &GithubDownloader{
		client:         github.NewClient(nil),
		ownerName:      owner,
		repositoryName: repository,
	}
}

func (d *GithubDownloader) GetAvailableReleases() ([]Version, error) {
	releases, _, err := d.client.Repositories.ListReleases(context.Background(), d.ownerName, d.repositoryName, nil)
	if err != nil {
		return []Version{}, err
	}
	versions := make([]Version, len(releases))
	for i, r := range releases {
		versions[i], err = ParseVersionFromString(r.GetTagName())
	}
	return versions, nil
}

func (d *GithubDownloader) Download(version Version, destination string) (string, error) {
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(context.Background(), d.ownerName, d.repositoryName, nil)
	if err != nil {
		slog.Error("failed to retrieve releases", err)
		return "", err
	}
	for i, r := range releases {
		v, _ := ParseVersionFromString(r.GetTagName())
		if v.Branch == "blitz" {
			slog.Info(fmt.Sprintf("%d. %v", i+1, v))
		}
	}
	return "", nil
}
