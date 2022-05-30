package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const MAX_ATTEMPTS = 1024

func selectRandomDir(currentDir string) string {
	items, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not read directory '%s': %s\n", currentDir, err)
		os.Exit(1)
	}

	var dirs []string
	for _, item := range items {
		if item.IsDir() {
			dirs = append(dirs, item.Name())
		}
	}

	if len(dirs) == 0 || rand.Intn(2) == 0 {
		return currentDir
	}

	index := rand.Intn(len(dirs))
	return selectRandomDir(filepath.Join(currentDir, dirs[index]))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not get home directory of user: %s\n", err)
		os.Exit(1)
	}

	currentPath, err := filepath.Abs(os.Args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not get absolute path of %s: %s\n", os.Args[0], err)
		os.Exit(1)
	}

	currentName := filepath.Base(os.Args[0])
	for i := 0; i < MAX_ATTEMPTS; i++ {
		finalPath := filepath.Join(selectRandomDir(home), currentName)

		if finalPath != currentPath {
			if os.Rename(currentPath, finalPath) == nil {
				return
			}
		}
	}
}
