package controller

import (
	"bigcalc/calculator"
	"bigcalc/parser"
	"fmt"
	"math/big"
)

type Controller struct {
	Calc   *calculator.Calculator
	Parser *parser.Parser
}

func (ctrl *Controller) Run() {
	var aStr, bStr, op string
	fmt.Print("Введите первое число (или a^b): ")
	_, err := fmt.Scan(&aStr)
	if err != nil {
		return
	}
	fmt.Print("Введите оператор (+, -, *, /, ^): ")
	_, err = fmt.Scan(&op)
	if err != nil {
		return
	}
	fmt.Print("Введите второе число (или a^b): ")
	_, err = fmt.Scan(&bStr)
	if err != nil {
		return
	}

	a := ctrl.Parser.ParseNumber(aStr)
	b := ctrl.Parser.ParseNumber(bStr)

	result := ctrl.compute(op, a, b)

	if result == nil {
		fmt.Println("Неверный оператор")
		return
	}
	fmt.Println("Результат:", result.String())
}

func (ctrl *Controller) compute(op string, a, b *big.Int) *big.Int {
	switch op {
	case "+":
		return ctrl.Calc.Add(a, b)
	case "-":
		return ctrl.Calc.Sub(a, b)
	case "*":
		return ctrl.Calc.Mul(a, b)
	case "/":
		return ctrl.Calc.Div(a, b)
	case "^":
		return ctrl.Calc.Pow(a, b)
	}

	return nil
}
