package main

import (
	"bigcalc/calculator"
	"bigcalc/controller"
	"bigcalc/parser"
)

// 1.22 Калькулятор для работы с большими числами
func main() {
	calc := calculator.Calculator{}
	pars := parser.Parser{}
	ctrl := controller.Controller{
		Calc:   &calc,
		Parser: &pars,
	}
	ctrl.Run()
}
