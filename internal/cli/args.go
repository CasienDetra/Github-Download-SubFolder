package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Args struct {
	OutputDir string
	RepoLink  string
}

func Parse() (*Args, error) {
	args := os.Args[1:]

	if len(args) < 1 || len(args) > 3 {
		return nil, errors.New("invalid number of arguments")
	}

	var outputDir string
	var repoLink string

	i := 0
	for i < len(args) {
		arg := args[i]

		switch {
		case arg == "-d" || arg == "--dir":
			if i+1 >= len(args) {
				return nil, fmt.Errorf("%s flag requires a directory path", arg)
			}
			outputDir = args[i+1]
			i += 2

		case strings.HasPrefix(arg, "-d="):
			outputDir = arg[3:]
			i++

		case strings.HasPrefix(arg, "--dir="):
			outputDir = arg[6:]
			i++

		case strings.HasPrefix(arg, "https://github.com/"):
			repoLink = arg
			i++

		default:
			return nil, fmt.Errorf("unknown argument: %s", arg)
		}
	}

	if repoLink == "" {
		return nil, errors.New("GitHub repository link is required")
	}

	if outputDir == "" {
		outputDir = "."
	}

	return &Args{
		OutputDir: outputDir,
		RepoLink:  repoLink,
	}, nil
}
