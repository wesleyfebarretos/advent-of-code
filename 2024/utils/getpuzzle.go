package utils

import (
	"fmt"
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
	if len(os.Args) > 1 && os.Args[1] == "test" {
		fmt.Println("NOTICE -> Test puzzle as input")
		return findPuzzle("puzzle-test.txt")
	}

	return findPuzzle("puzzle.txt")
}
