package main

import "fmt"

//Разработать программу, которая в runtime способна определить тип переменной, переданной в неё (на вход подаётся interface{}).
//Типы, которые нужно распознавать: int, string, bool, chan (канал).
//Подсказка: оператор типа switch v.(type) поможет в решении.

func getVarType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	case chan bool:
		fmt.Println("chan bool")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	getVarType(1)
	getVarType("hello")
	getVarType(true)
	getVarType(make(chan int))
	getVarType(make(chan string))
	getVarType(make(chan bool))
}
