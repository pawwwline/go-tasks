package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
//По завершению программы структура должна выводить итоговое значение счётчика.
//Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.

type ConcurrentCounter struct {
	value int
	sync.Mutex
}

func NewCounter() *ConcurrentCounter {
	return &ConcurrentCounter{}
}

func (cc *ConcurrentCounter) Value() int {
	cc.Lock()
	defer cc.Unlock()
	return cc.value
}

func (cc *ConcurrentCounter) Inc() {
	cc.Lock()
	cc.value++
	cc.Unlock()
}

type ConcurrentCounterAtomic struct {
	value atomic.Uint64
}

func NewCounterAtomic() *ConcurrentCounterAtomic {
	return &ConcurrentCounterAtomic{}
}

func (cc *ConcurrentCounterAtomic) Value() int {
	return int(cc.value.Load())
}

func (cc *ConcurrentCounterAtomic) Inc() {
	cc.value.Add(1)
}

func main() {
	counter := NewCounter()
	counterAtomic := NewCounterAtomic()
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Go(func() {
			counter.Inc()
		})
	}

	for i := 0; i < 1000; i++ {
		wg.Go(func() {
			counterAtomic.Inc()
		})
	}

	wg.Wait()

	fmt.Println(counter.Value())
	fmt.Println(counterAtomic.Value())

}
