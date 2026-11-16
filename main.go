package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements returns size positive random integers.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(math.MaxInt) + 1
	}
	return data
}

// maximum returns the largest value, or zero for empty input.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	largest := data[0]
	for _, value := range data[1:] {
		if value > largest {
			largest = value
		}
	}
	return largest
}

// maxChunks searches nonempty chunks concurrently.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	count := min(CHUNKS, len(data))
	maxima := make([]int, count)
	chunkSize, remainder := len(data)/count, len(data)%count
	var workers sync.WaitGroup
	workers.Add(count)
	start := 0
	for i := range count {
		end := start + chunkSize
		if i < remainder {
			end++
		}
		go func(index int, chunk []int) {
			defer workers.Done()
			maxima[index] = maximum(chunk)
		}(i, data[start:end])
		start = end
	}
	workers.Wait()
	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Последовательный поиск")
	start := time.Now()
	largest := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение: %d\nВремя поиска: %d мкс\n", largest, elapsed)

	fmt.Printf("Параллельный поиск в %d горутинах\n", CHUNKS)
	start = time.Now()
	largest = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение: %d\nВремя поиска: %d мкс\n", largest, elapsed)
}
