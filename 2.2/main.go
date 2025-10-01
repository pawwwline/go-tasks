//Что выведет программа?
//
//Объяснить порядок выполнения defer функций и итоговый вывод.

package main

import "fmt"

// 2
// мы изначально записываем  переменную в return slot, поэтому мы не теряем связь с ней при выходе на return, и можем изменять через defer, который происходит после return
func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

// 1
// здесь мы сразу при return записываем значение не имея к нему в дальнейшем никакого доступа, то есть оно никак не связано с локальной переменной
func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())        //2
	fmt.Println(anotherTest()) //1
}
