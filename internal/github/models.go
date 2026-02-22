package github

type Content struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Type        string `json:"type"`
	DownloadURL string `json:"download_url"`
	URL         string `json:"url"`
}

type RepoInfo struct {
	Owner  string
	Repo   string
	Branch string
	Path   string
}
