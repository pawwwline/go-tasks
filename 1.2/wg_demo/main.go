package main

import (
	"fmt"
	"sync"
)

// используем новый метод Go, который заменяет Add(1) и Done()
func powSync(nums []int, wg *sync.WaitGroup) {
	for _, num := range nums {
		num := num
		wg.Go(func() {
			num = num * num
			fmt.Println(num)
		})
	}
}

// синхронизация через WaitGroup
func main() {
	nums := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	powSync(nums, &wg)
	wg.Wait()
}
