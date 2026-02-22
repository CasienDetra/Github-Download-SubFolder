package downloader

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/casien/ghsub-download/internal/github"
)

func DownloadSubDir(repo github.RepoInfo, outputDir string) error {
	folderName := filepath.Base(repo.Path)
	outputPath := filepath.Join(outputDir, folderName)

	if err := os.MkdirAll(outputPath, 0o755); err != nil {
		return err
	}

	return downloadDir(repo, repo.Path, outputPath)
}

func downloadDir(repo github.RepoInfo, remotePath, localPath string) error {
	contents, err := github.GetDirContent(repo, remotePath)
	if err != nil {
		return err
	}

	for _, item := range contents {

		itemLocalPath := filepath.Join(localPath, item.Name)

		if item.Type == "dir" {

			if err := os.MkdirAll(itemLocalPath, 0o755); err != nil {
				return err
			}

			if err := downloadDir(repo, item.Path, itemLocalPath); err != nil {
				return err
			}

		} else if item.Type == "file" {
			if err := downloadFile(item.DownloadURL, itemLocalPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func downloadFile(downloadURL, localPath string) error {
	req, err := http.NewRequest("GET", downloadURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "myGithubCli")

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
