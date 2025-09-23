package main

import "fmt"

//Поменять местами два числа без использования временной переменной.
//Подсказка: примените сложение/вычитание или SwapXOR-обмен.

// Swap используем множественное присваивание, самый оптимальный для использования в разработке (компилятор все равно использует временные переменные под капотом)
func Swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

// SwapXOR Используем битовое XOR (работает только для целочисленных переменных)
func SwapXOR(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func SwapMath(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func main() {
	a := 10
	b := 20
	fmt.Printf("a = %d b = %d\n", a, b)
	a, b = Swap(a, b)
	fmt.Printf("a = %d b = %d\n", a, b)

	a = 10
	b = 20
	fmt.Printf("a = %d b = %d\n", a, b)
	a, b = SwapXOR(a, b)
	fmt.Printf("a = %d b = %d\n", a, b)

	a = 10
	b = 20
	fmt.Printf("a = %d b = %d\n", a, b)
	a, b = SwapMath(a, b)
	fmt.Printf("a = %d b = %d\n", a, b)

}
