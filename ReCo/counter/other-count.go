package counter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync/atomic"
)

func CountOthers(filename string) ([]int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("file not found")
	}
	defer file.Close()

	var vowelCount int64
	var consonantCount int64
	var pMarkCount int64
	var numCount int64
	var spaceCount int64

	lines := make(chan string, 20)
	done := make(chan bool)

	go func() {
		scanner := bufio.NewScanner(bufio.NewReader(file))
		for scanner.Scan() {
			jobs <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
		close(jobs)
	}()

	go countQuery(jobs, done, &vowelCount, &pMarkCount, &numCount, &consonantCount, &spaceCount)
	<-done
	close(done)

	var result []int64
	result = append(result, vowelCount, consonantCount, spaceCount, digitCount, pMarkCount, vowelCount+consonantCount)
	return result, nil
}

func countQuery(jobs <-chan string, done chan<- bool, vowelCount, pMarkCount, numCount, consonantCount, spaceCount *int64) {
	for w := range jobs {
		for _, s := range w {
			switch s {
			case 97, 101, 105, 111, 117, 65, 69, 73, 79, 85:
				atomic.AddInt64(vowelCount, 1)
			case 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
				46, 47, 58, 59, 60, 61, 62, 63, 64, 92, 93, 95, 96,
				123, 124, 125, 126:
				atomic.AddInt64(pMarkCount, 1)
			case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
				atomic.AddInt64(numCount, 1)
			case 32, 10:
				atomic.AddInt64(spaceCount, 1)
			default:
				atomic.AddInt64(consonantCount, 1)
			}
		}
	}
}
