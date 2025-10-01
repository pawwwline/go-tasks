package main

import "fmt"

//Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.
//
//Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента (copy(slice[i:], slice[i+1:]))
//и уменьшить длину слайса на 1.

func deleteSliceI(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(deleteSliceI(slice, 2))
}
