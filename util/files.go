package util

import (
	"bufio"
	"io"
	"log"
	"os"
)

type FileScanner struct {
	io.Closer
	*bufio.Scanner
}

func OpenFile(filePath string) *FileScanner {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Could not open input file")
	}

	scanner := bufio.NewScanner(file)

	return &FileScanner{file, scanner}
}
