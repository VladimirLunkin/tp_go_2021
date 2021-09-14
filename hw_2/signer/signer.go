package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	iterationsNum = 7
	goroutinesNum = 5
)

func startWorker(in int, waiter *sync.WaitGroup) {
	defer waiter.Done() // wait_2.go уменьшаем счетчик на 1
	localWg := &sync.WaitGroup{}
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		time.Sleep(time.Millisecond) // попробуйте убрать этот sleep
	}

	localWg.Wait()
}

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	for _, j := range jobs {
		wg.Add(1)
		go j(1, 2)
	}
	wg.Wait()
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}

func main() {
	ExecutePipeline()
}
