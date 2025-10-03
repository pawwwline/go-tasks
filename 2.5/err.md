**Что выведет программа?**

Объяснить вывод программы.


```go 
package main

type customError struct {
	msg string
}   

func (e *customError) Error() string {
    return e.msg
}

func test() *customError {
// ... do something
    return nil
}

func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
    }

```

Программа выведет error, несмотря на то само значение ошибки у нас == nil.
Это связано с тем, что error представляет собой интерфейс, который реализуется всеми типами с методом Error.
и так как в функции test() мы возвращаем нашу кастомную ошибку, которая имеет свою реализацию, поэтому у интерфейса error появляется тип, значит он не является nil интерфейсом (то есть, таким где и тип и значение равны nil) 