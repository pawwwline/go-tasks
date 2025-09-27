package main

import "fmt"

//Разработать программу, которая переворачивает порядок слов в строке.
//Пример: входная строка:
//«snow dog sun», выход: «sun dog snow».
//Считайте, что слова разделяются одиночным пробелом. Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».

func reverseWords(s string) string {
	//преобразовываем в массив байтов, чтобы изменять строку
	bytes := []byte(s)
	//разворачиваем всю строку
	reverse(bytes)
	l := 0
	r := 0
	for l < len(bytes) {
		//проверяем наличие пробелов в начале
		for l < len(bytes) && bytes[l] == ' ' {
			l++
		}
		// ищем границу слова
		r = l
		for r < len(bytes) && bytes[r] != ' ' {
			r++
		}
		//разворачиваем слово
		reverse(bytes[l:r])
		l = r
	}

	return string(bytes)
}

func reverse(s []byte) {
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}

}

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str))
}
