package gggitea

type environment struct {
	baseUrl  string
	apiToken string
}

func getEnvironment() (*environment, error) {
	e := environment{}
	// TODO: Read CLI, ENV, or use default
	e.baseUrl = "https://gitea.com"
	// TODO: Read CLI, ENV, or fail
	e.apiToken = ""
	return &e, nil
}
