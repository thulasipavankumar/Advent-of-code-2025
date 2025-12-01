package util

import (
	"bufio"
	"log"
	"os"
)

func GetFileContent(file_path string) (lines []string) {
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}
