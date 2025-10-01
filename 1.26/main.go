package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке встречаются один раз (т.е. строка состоит из уникальных символов).
//
//Вывод: true, если все символы уникальны, false, если есть повторения. Проверка должна быть регистронезависимой,
//т.е. символы в разных регистрах считать одинаковыми.
//
//Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.
//
//Подумайте, какой структурой данных удобно воспользоваться для проверки условия.

// Использование множества позволит легко проверить условие
func checkUnique(str string) bool {
	set := make(map[rune]struct{})
	str = strings.ToLower(str)
	for _, r := range str {
		_, ok := set[r]
		if ok {
			return false
		}
		set[r] = struct{}{}
	}
	return true
}

func main() {
	str := "abCdefAaf"
	fmt.Println(checkUnique(str))
}
