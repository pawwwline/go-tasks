package main

import (
	"fmt"
)

// используем канал для синхронизации
func main() {
	nums := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	for _, num := range nums {
		go func(x int) {
			ch <- x * x
		}(num)
	}

	for range nums {
		fmt.Println(<-ch)
	}
}
