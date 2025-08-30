package main

import (
	"fmt"
	"time"
)

func pow(nums []int) {
	for _, num := range nums {
		go func(x int) {
			x = x * x
			fmt.Println(x)
		}(num)
	}
}

// просто блокируем мейн горутину ожиданием (очень костыльный способ, просто для демонстрации)
func main() {
	nums := []int{2, 4, 6, 8, 10}
	go pow(nums)
	time.Sleep(3 * time.Second)

}
