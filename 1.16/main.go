package main

import (
	"fmt"
	"math/rand"
)

//Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.
//
//Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел. Для выбора опорного элемента можно взять середину или первый элемент.

// используем рекурсию, поэтому для начала определим базовый случай (массив отсортирован)
// разделим на две части левую, где будут все элементы меньше опорного
// правая где все элементы больше опорного
// рекурсивно соединяем

// O(n log n) в среднем, в худшем случае O(n^2); память O(n)
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr //базовый случай нам уже нечего сортировать
	}

	pivot := arr[0]
	left := []int{}
	right := []int{}
	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v) //разделяем на левую часть меньше опорного
		} else {
			right = append(right, v) // правая часть больше опорного
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...) //рекурсивно соединяем левую часть + пивот + правая часть
}

// уменьшаем вероятность худшего случая используя псевдорандомность,
// исключая вероятность того что при отсортированных данных одна часть у нас будет переполнена, а другая пустой
func quickSortRand(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivotIndex := rand.Intn(len(arr))
	pivot := arr[pivotIndex]

	left := []int{}
	right := []int{}

	for i, v := range arr {
		if i == pivotIndex {
			continue
		}
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(quickSortRand(left), append([]int{pivot}, quickSortRand(right)...)...)

}

func main() {
	arr := []int{10, 2, 7, 1, 1, 3, 20, 2, 7, 8}

	fmt.Println(quickSort(arr))
	fmt.Println(quickSortRand(arr))
}
