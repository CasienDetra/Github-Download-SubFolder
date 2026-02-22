package main

import (
	"fmt"
	"os"

	"github.com/casien/ghsub-download/internal/cli"
	"github.com/casien/ghsub-download/internal/downloader"
	"github.com/casien/ghsub-download/internal/github"
)

func main() {
	args, err := cli.Parse()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Usage: app [-d <dir>] <github-repo-link>")
		os.Exit(1)
	}

	repoInfo, err := github.ParseLink(args.RepoLink)
	if err != nil {
		fmt.Println("Invalid GitHub link:", err)
		os.Exit(1)
	}

	fmt.Printf("Downloading %s/%s to %s...\n", repoInfo.Owner, repoInfo.Repo, args.OutputDir)

	if err := downloader.DownloadSubDir(*repoInfo, args.OutputDir); err != nil {

		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("✓ Download completed successfully!")
}
