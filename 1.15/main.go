package main


//Рассмотреть следующий код и ответить на вопросы: к каким негативным последствиям он может привести и как это исправить?
//
//Приведите корректный пример реализации.


var justString string

func someFunc() {
	v := createHugeString(1 &lt; &lt; 10)
	justString = v[:100] //когда мы берем подстроку большой строки,
	// то мы все еще работаем с изначальной большой и GC не может очистить большую строку даже в том случае если она в дальнейшем не будет никак использоваться
}


func correctedFunc(){
	v := createHugeString(1 &lt;&lt; 10)
	justString = string(v[:100])

}


func main() {
	someFunc()
	correctedFunc()
}