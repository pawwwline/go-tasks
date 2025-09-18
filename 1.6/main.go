package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

//основные способы остановки горутины

// 1. завершение горутины через канал (закрываем канал, который горутина слушает || записываем значение в канал, который она слушает)
// 2. использование пакета контекст (отправляем в горутину и вызываем context.Cancel снаружи)
// 3. использование runtime.Goexit() // но это не является идиоматическим подходом
// 4. return выводит горутину из стека
// 5. panic (отлавливаем выше чтобы она не прошла дальше)
// 6. Закрытие ресурса

func main() {
	ch := make(chan int)
	stopGoChan(ch)

	ch <- 0
	close(ch)

	ctx, cancel := context.WithCancel(context.Background())
	stopGoCtx(ctx)
	cancel()

	stopGoExit()
	stopGoReturn()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic caught:", r)
		}
	}()
	stopGoPanic()

	data := make(chan []int)
	stopGoClose(data)
	data <- []int{1, 2, 3}
	close(data)

	time.Sleep(time.Second * 5)

}

func stopGoChan(ch chan int) {
	go func() {
		for {
			select {
			case <-ch:
				return
			default:
				fmt.Println("working")
			}
		}
	}()
}

func stopGoCtx(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("working")
			}
		}
	}()
}

func stopGoExit() {
	go func() {
		for {
			fmt.Println("working")
			runtime.Goexit()
		}
	}()
}

func stopGoReturn() {
	go func() {
		for {
			fmt.Println("working")
			return
		}
	}()
}

func stopGoPanic() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("panic caught:", r)
			}
		}()

		for {
			fmt.Println("working")
			panic("panic")
		}
	}()
}

func stopGoClose(data chan []int) {
	go func() {
		for _ = range data {
			fmt.Println("working")
		}
	}()
}
