package git

import (
	"io/fs"
	"os"
	"os/exec"
	"path"
	"strings"
)

func GetRootDir() (*string, error) {
	result, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()

	if err != nil {
		return nil, err
	}

	rootDir := strings.Trim(string(result[:]), "\n")
	return &rootDir, nil
}

func ReadIgnoreFile() []string {
	rootPath, err := GetRootDir()
	if err != nil {
		return []string{}
	}

	gitIgnorePath := *rootPath + "/.gitignore"
	gitIgnoreFile, err := os.ReadFile(gitIgnorePath)

	if err != nil {
		return []string{}
	}

	lines := strings.Split(string(gitIgnoreFile[:]), "\n")
	return lines
}

func FilterFiles(nodes []fs.DirEntry, patterns []string) []fs.DirEntry {
	result := []fs.DirEntry{}

	for _, node := range nodes {
		matches := false
		for _, pattern := range patterns {
			match, _ := path.Match(pattern, node.Name())

			if match {
				matches = true
				break
			}
		}

		if !matches {
			result = append(result, node)
		}
	}
	return result
}
