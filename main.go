package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	baseDir := os.Args[1]
	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		dirPath := filepath.Join(baseDir, entry.Name())
		gitDir := filepath.Join(dirPath, ".git")
		info, err := os.Stat(gitDir)
		if err != nil || !info.IsDir() {
			continue
		}
		fmt.Println(dirPath)
	}
}
