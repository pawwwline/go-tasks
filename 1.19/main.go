package main

import "fmt"

//Разработать программу, которая переворачивает подаваемую на вход строку.
//Например: при вводе строки «главрыба» вывод должен быть «абырвалг».
//Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
//то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).

// Сложность O(n); память O(n)
func reverse(s string) string {
	runes := []rune(s)
	l := 0
	r := len(runes) - 1

	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
	return string(runes)

}

func main() {
	str := "приветик"
	fmt.Println(reverse(str))
}
