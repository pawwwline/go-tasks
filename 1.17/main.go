package main

import "fmt"

// Сложность O(log(n)), память O(1)
// работает только в случае отсортированных данных
func binarySearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			l = mid + 1 //если средний элемент меньше искомого значит ищем в правой половине
		} else {
			r = mid - 1 //ищем в левой половине
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 10
	index := binarySearch(arr, target)
	fmt.Println(index)
}
