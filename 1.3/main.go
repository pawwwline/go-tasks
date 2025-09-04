package main

import (
	"fmt"
	"sync"
)

func worker(jobs <-chan string) {
	for job := range jobs {
		fmt.Println(job)
	}
}

func workerPool(numWorkers int, jobs <-chan string, wg *sync.WaitGroup) {
	for i := 0; i < numWorkers; i++ {
		wg.Go(func() {
			worker(jobs)
		})
	}
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan string, 10)
	workerPool(10, jobs, &wg)

	for i := 0; i < 10; i++ {
		jobs <- fmt.Sprintf("task-%d", i)
	}
	close(jobs)

	wg.Wait()
}
