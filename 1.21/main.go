package main

import (
	"fmt"
)

//Реализовать паттерн проектирования «Адаптер» на любом примере.
//
//Описание: паттерн Adapter позволяет сконвертировать интерфейс одного класса в интерфейс другого, который ожидает клиент.
//
//Продемонстрируйте на простом примере в Go: у вас есть существующий интерфейс (или структура)
//и другой, несовместимый по интерфейсу потребитель — напишите адаптер,
//который реализует нужный интерфейс и делегирует вызовы к встроенному объекту.
//Поясните применимость паттерна, его плюсы и минусы, а также приведите реальные примеры использования.

type OldLogger struct {
}

func (l *OldLogger) WriteLog(message string) {
	fmt.Println(message)
}

// у нас есть пример интерфейса логгера, который используется в остальной части программы, но при этом методы старого логгера ему не соответсвуют
type Logger interface {
	LogInfo(message string)
}

type HTTPServer struct {
	Logger Logger
}

func NewHTTPServer(logger Logger) *HTTPServer {
	return &HTTPServer{
		Logger: logger,
	}
}

func (h HTTPServer) SomeFunc() {
	method := "someMethod was called"
	h.Logger.LogInfo(method)
}

//Создаем структуру, которая будет реализовывать методы интерфейса и при этом вызывать метод старого логгера
type AdapterLogger struct {
	oldLogger *OldLogger
}

func NewAdapterLogger(oldLogger *OldLogger) *AdapterLogger {
	return &AdapterLogger{
		oldLogger: oldLogger,
	}
}

func (l *AdapterLogger) LogInfo(message string) {
	l.oldLogger.WriteLog(message)
}

func main() {
	oldLogger := &OldLogger{}
	adapterLogger := NewAdapterLogger(oldLogger)
	httpServer := NewHTTPServer(adapterLogger)
	httpServer.SomeFunc()

}
