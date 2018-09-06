package repository

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
)

const GITHUB_PROVIDER_API4 = "github4"

type Github struct {
	token  string
	apiUrl string
}

//data.organization.repositories.nodes
type GithubResponse struct {
	Data struct {
		Organization struct {
			Repositories struct {
				Nodes []Repository `json:"nodes"`
			} `json:"repositories"`
		} `json:"organization"`
	} `json:"data"`
}

func NewGithubV4(token string) Github {
	return Github{
		token:  token,
		apiUrl: "https://api.github.com/graphql",
	}
}

func (r Github) GetRepositories() []Repository {
	query := `
 query {
  organization(login:"IlluminateEducation") {
    	name,
      repositories(first:100) {
      totalCount
      nodes {
        name,
        description,
        url,
        sshUrl
      }
    }
  }
}
`
	t := map[string]string{
		"query": query,
	}

	body, err := json.Marshal(t)

	//Authorization: bearer
	req, err := http.NewRequest("POST", r.apiUrl, bytes.NewReader(body))
	if err != nil {
		return nil
	}

	req.Header.Set("Authorization", "bearer "+r.token)

	c := http.Client{}
	res, err := c.Do(req)

	if err != nil {
		panic(err)
	}

	ghr := GithubResponse{}
	decoder := json.NewDecoder(res.Body)
	err  = decoder.Decode(&ghr)

	if err != nil {
		fmt.Print(err)
	}

	return ghr.Data.Organization.Repositories.Nodes
}
