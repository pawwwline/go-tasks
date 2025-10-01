package main

import "fmt"

// Что выведет программа?
//
// Объясните вывод программы.
// Выведет [77 78 79]
func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}
