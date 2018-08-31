package repository

import "github.com/allenisalai/ice/internal"

type RepositoryInterface interface {
	GetRepositories() []Repository
}

type Repository struct {
	Name string
	Url  string
}

func Factory(config ice.Configuration) RepositoryInterface {
	switch config.RepositoryProvider {
	case GITHUB_PROVIDER_API4:
		return NewGithubV4(config.ApiToken)
	}

	panic("Invalid Repository Provider")
}

