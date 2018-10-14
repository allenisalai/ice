package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const GITHUB_PROVIDER_API4 = "github4"

type Github struct {
	token  string
	apiUrl string
}

type GithubV4Response struct {
	Data struct {
		Organization struct {
			Repositories struct {
				PageInfo struct {
					EndCursor   string `json:"endCursor"`
					HasNextPage bool   `json:"hasNextPage"`
				} `json:"pageInfo"`
				Edges []struct {
					Node Repository `json:"node"`
				} `json:"edges"`
				TotalCount int32 `json:"totalCount"`
			} `json:"repositories"`
		} `json:"organization"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

func NewGithubV4(token string) Github {
	return Github{
		token:  token,
		apiUrl: "https://api.github.com/graphql",
	}
}

func (r Github) GetRepositories() ([]Repository, error) {
	pageSize := 100
	lastCursor := ""
	query := `
 query {
  organization(login:"IlluminateEducation") {
      name,
      repositories(first:%d, %s orderBy:{field:NAME, direction:ASC}) {
      pageInfo {
        endCursor
        hasNextPage
      }
      totalCount
      edges {
        node{
          id,
          name,
          description,
          url,
          sshUrl
        }
        cursor
      }
    }
  }
}
`
	var repositories []Repository
	ghr, err := r.queryGithubV4Repositories(query, pageSize, lastCursor)
	if err != nil {
		return []Repository{}, err
	}
	repositories = append(repositories, ghr.getRepositories()...)

	for ghr.Data.Organization.Repositories.PageInfo.HasNextPage {
		ghr, err = r.queryGithubV4Repositories(query, pageSize, ghr.Data.Organization.Repositories.PageInfo.EndCursor)

		if err != nil {
			return []Repository{}, err
		}
		repositories = append(repositories, ghr.getRepositories()...)
	}
	return repositories, nil
}

func (r Github) formatGithubV4Query(query string, pageSize int, lastCursor string) string {
	afterClause := ""
	if len(lastCursor) > 0 {
		afterClause = fmt.Sprintf("after:\"%s\",", lastCursor)
	}

	return fmt.Sprintf(query, pageSize, afterClause)
}

func (r Github) queryGithubV4Repositories(query string, pageSize int, lastCursor string) (GithubV4Response, error) {
	t := map[string]string{
		"query": r.formatGithubV4Query(query, pageSize, lastCursor),
	}

	body, err := json.Marshal(t)
	req, err := http.NewRequest("POST", r.apiUrl, bytes.NewReader(body))
	if err != nil {
		return GithubV4Response{}, err
	}

	req.Header.Set("Authorization", "bearer "+r.token)

	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return GithubV4Response{}, err
	}
	defer res.Body.Close()

	ghr := GithubV4Response{}
	err = json.NewDecoder(res.Body).Decode(&ghr)

	return ghr, err
}

func (ghr GithubV4Response) getRepositories() []Repository {
	var repos []Repository
	for _, e := range ghr.Data.Organization.Repositories.Edges {
		repos = append(repos, e.Node)
	}

	return repos
}
