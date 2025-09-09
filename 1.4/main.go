package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(ctx context.Context, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: контекст отменён")
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Println(job)
		}
	}
}

func workerPool(ctx context.Context, numWorkers int, jobs <-chan string, wg *sync.WaitGroup) {
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, jobs, wg)
	}
}

func chanWriter(ctx context.Context, numJobs int, jobs chan<- string) {
	for i := 0; i < numJobs; i++ {
		select {
		case <-ctx.Done():
			close(jobs)
			return
		case jobs <- fmt.Sprintf("task-%d", i):
		}
	}
	close(jobs)
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	jobs := make(chan string)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		cancel()
	}()
	workerPool(ctx, 10, jobs, &wg)
	go chanWriter(ctx, 100000, jobs)
	wg.Wait()
}
