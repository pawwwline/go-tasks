package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала
//– читать эти значения. По истечении N секунд программа должна завершаться.
//
//Подсказка: используйте time.After или таймер для ограничения времени работы.

// Продюсер генерирует бесконечное количество чисел и записывает их в канал (с перерывом в секунду)

func producer(ctx context.Context, ch chan<- int) {
	defer close(ch) // закрываем канал при выходе из функции продюсера
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			return
		case ch <- i:
		}
		time.Sleep(time.Second)
	}
}

func consumer(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-ch:
			if !ok { //проверяем канал на то был ли он закрыт
				return
			}
			fmt.Println(v)

		}
	}
}

func main() {
	ch := make(chan int) //создаем не буферизированный канал, при котором будет блокироваться консьюмер
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) //создаем контекст с тайм-аутом в 10 секунд
	defer cancel()
	wg.Go(func() {
		producer(ctx, ch)
	})
	wg.Go(func() {
		consumer(ctx, ch)
	})
	wg.Wait() //исключаем рассинхронизацию горутин
}
