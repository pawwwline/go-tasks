package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep, которая приостанавливает выполнение текущей горутины.
//
//Важно: в отличие от настоящей time.Sleep,
//ваша функция должна именно блокировать выполнение (например, через таймер или цикл), а не просто вызывать time.Sleep :) — это упражнение.
//
//Можно использовать канал + горутину, или цикл на проверку времени (не лучший способ, но для обучения).

func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C

	return
}

func TestGoroutine() {
	fmt.Println("test")
}

func main() {
	go TestGoroutine()
	sleep(time.Second * 5)
}
