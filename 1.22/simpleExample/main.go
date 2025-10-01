package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

//Разработать программу, которая перемножает, делит, складывает,
//вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).
//
//Комментарий: в Go тип int справится с такими числами,
//но обратите внимание на возможное переполнение для ещё больших значений. Для очень больших чисел можно использовать math/big.

// Простой пример программы для наглядной демонстрации переполнения int64

func Sum(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}

func SumBig(a, b *big.Int) *big.Int {
	z := big.NewInt(0)
	return z.Add(a, b)
}

func SubBig(a, b *big.Int) *big.Int {
	z := big.NewInt(0)
	return z.Sub(a, b)
}

func MulBig(a, b *big.Int) *big.Int {
	z := big.NewInt(0)
	return z.Mul(a, b)
}

func DivBig(a, b *big.Int) *big.Int {
	z := big.NewInt(0)
	return z.Div(a, b)
}

func main() {
	a := int(math.Pow(10, 30))
	b := int(math.Pow(10, 20))
	fmt.Print(">int64 " + strconv.Itoa(a) + " + " + strconv.Itoa(b) + " = ")
	fmt.Println(Sum(a, b))

	fmt.Print(">int64 " + strconv.Itoa(a) + " - " + strconv.Itoa(b) + " = ")
	fmt.Println(Sub(a, b))

	fmt.Print(">int64 " + strconv.Itoa(a) + " * " + strconv.Itoa(b) + " = ")
	fmt.Println(Mul(a, b))

	fmt.Print(">int64 " + strconv.Itoa(a) + " / " + strconv.Itoa(b) + " = ")
	fmt.Println(Div(a, b))

	aBase := big.NewInt(10)
	aExp := big.NewInt(30)
	aBig := big.NewInt(0).Exp(aBase, aExp, nil)

	bBase := big.NewInt(10)
	bExp := big.NewInt(20)
	bBig := big.NewInt(0).Exp(bBase, bExp, nil)

	fmt.Print(">big " + aBig.String() + " + " + bBig.String() + " = ")
	fmt.Println(SumBig(aBig, bBig))

	fmt.Print(">big " + aBig.String() + " - " + bBig.String() + " = ")
	fmt.Println(SubBig(aBig, bBig))

	fmt.Print(">big " + aBig.String() + " * " + bBig.String() + " = ")
	fmt.Println(MulBig(aBig, bBig))

	fmt.Print(">big " + aBig.String() + " / " + bBig.String() + " = ")
	fmt.Println(DivBig(aBig, bBig))

}
