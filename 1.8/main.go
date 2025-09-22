package main

import "fmt"

//Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
//Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
//Подсказка: используйте битовые операции (|, &^).

func changeBit(x int, i int, v int) int {
	idx1 := i - 1
	if idx1 < 0 && (v != 0 && v != 1) { //используем индексацию с 1, как в условии
		return -1
	}
	if v == 1 {
		return x | (1 << idx1) // устанавливаем бит на 1 с помощью оператора OR (мы указываем на нужный индекс с помощью маски)
	}
	return x &^ (1 << idx1) //устанавливаем бит на 0 с помощью оператора AND NOT
}

func main() {
	var x int
	var index int
	var value int
	_, err := fmt.Scan(&x, &index, &value)
	if err != nil {
		return
	}
	newInt := changeBit(x, index, value)
	if newInt == -1 {
		fmt.Println("error getting new int, notice that index starts from 1")
	}
	fmt.Println(newInt)

}
