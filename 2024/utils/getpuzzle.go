package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func findPuzzle(name string) string {
	_, currentFilePath, _, ok := runtime.Caller(3)
	if !ok {
		log.Fatal("puzzle file not found")
	}

	currentDir := filepath.Dir(currentFilePath)

	puzzle, err := os.ReadFile(filepath.Join(currentDir, name))
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(puzzle))
}

func GetPuzzle() string {
	return findPuzzle("puzzle.txt")
}

func GetTestPuzzle() string {
	return findPuzzle("puzzle-test.txt")
}
