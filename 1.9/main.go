package main

import (
	"fmt"
)

//Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2.
//После этого данные из второго канала должны выводиться в stdout.
//То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
//Убедитесь, что чтение из второго канала корректно завершается.

func writeChan(arr []int, ints chan<- int) {
	for _, v := range arr { //проходися по массиву
		ints <- v
	}
	close(ints) //у нас нет больше данных для записи закрываем канал

}

func processChan(ints <-chan int, doubled chan<- int) {
	for v := range ints { //читаем все данные из канала
		doubled <- v * 2 //записываем в канал удвоенных
	}
	close(doubled) //у нас закончились данные закрываем канал
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	ints := make(chan int, len(arr)) //создаем буферизированные каналы, чтобы уменьшить количество задержек
	doubled := make(chan int, len(arr))
	go writeChan(arr, ints)
	go processChan(ints, doubled)
	for v := range doubled { //пока не закроется канал doubled читаем в основной горутине
		fmt.Println(v)
	}

}
