package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetDirContent(repo RepoInfo, path string) ([]Content, error) {
	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/contents/%s?ref=%s",
		repo.Owner,
		repo.Repo,
		path,
		repo.Branch,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "myGithubCli")
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Error github api please try again %d: %s", resp.StatusCode, string(body))
	}
	var contents []Content
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		return nil, err
	}
	return contents, nil
}
