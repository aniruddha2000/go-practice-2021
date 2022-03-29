package counter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync/atomic"
)

func CountWord(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return -1, errors.New("file not found")
	}
	defer file.Close()

	var wordCount int64

	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		atomic.AddInt64(&wordCount, 1)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v", err)
	}

	return wordCount, nil
}
