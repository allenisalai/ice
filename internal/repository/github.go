package repository

const GITHUB_PROVIDER_API4 = "github4"

type Github struct {
	token  string
	apiUrl string
}

func NewGithubV4(token string) Github {
	return Github{
		token:  token,
		apiUrl: "https://api.github.com/graphql",
	}
}

func (r Github) GetRepositories() []Repository {
	var repos []Repository
/*
	query := `
query FindIssueID {
	repository(owner:"octocat", name:"Hello-World") {
    issue(number:349) {
      id
    }
  }
}
`

	//Authorization: bearer
	req, err := http.NewRequest("POST", r.apiUrl, query)
	if err != nil {
		return nil
	}

	req.Header.Set("Authorization", "bearer "+r.token)
*/
	return repos
}
