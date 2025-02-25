package value_objects

type PullRequestEvent struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
}

type PullRequest struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Head  Branch `json:"head"`
	Base  Branch `json:"base"`
	URL   string `json:"url"`
	User  User   `json:"user"`
}

type Branch struct {
	Ref string `json:"ref"`
	SHA string `json:"sha"`
}

type User struct {
	UserName string `json:"login"`
}

type Repository struct {
	Name string `json:"name"`
}
