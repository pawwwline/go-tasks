package main

import (
	"fmt"
	"os"
)

//Интерфейсы под капотом - это структура с значеннием и конкретным типом объекта
// значение объекта может быть равно нулю, но это не делает интерфейс nil интерфейсом, для этого нужно чтобы и не было типа объекта, а в данном случае он есть

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
	fmt.Printf("%T\n", err)
}
