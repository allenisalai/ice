package repository

import "github.com/allenisalai/ice/internal"

type Interface interface {
	GetRepositories() ([]Repository, error)
}

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	SshUrl      string `json:"sshUrl"`
}

func Factory(config ice.Configuration) Interface {
	switch config.RepositoryProvider {
	case GITHUB_PROVIDER_API4:
		return NewGithubV4(config.ApiToken)
	}

	panic("Invalid Repository Provider")
}
