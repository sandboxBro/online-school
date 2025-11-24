package main

import (
	"fmt"
	"log"
)

// Creator (создатель строк)
type Creator struct {
	Chars []rune
}

// Т.е. это структура у которой есть поле Chars - которое является слайсом букв
// В поле Chars как раз будут перечислены буквы из которых Creator будет создавать строку
// Теперь нам надо описать этой структуре метод Create:

// Эта функция относится к структуре Creator - (c Creator)
// Называется Create
// Принимает на вход кол-во символов из скольких должна состоять строка (n int)
// И в результате возвращает строку string
func (c Creator) Create(n int) string {
	// если просим больше чем есть - берем сколько можем
	if n > len(c.Chars) {
		n = len(c.Chars)
	}
	result := c.Chars[:n] // берём первые n символов из слайса
	return string(result) // преобразуем слайс рун в строку и возвращаем
}

// Processor (обработчик строк)
type Processor struct {
	MaxLen int
}

// Т.е. это структура у которой есть поле MaxLen - максимально допустимое кол-во
// символов в строке
// Теперь нам надо описать этой структуре метод Process:

// Эта функция относится к структуре Processor - (p Processor)
// Называется Process
// Принимает на вход строку (text string)
// И в результате возвращает ошибку если строка длинее заданного MaxLen
func (p Processor) Process(text string) error {
	runes := []rune(text)
	if len(runes) > p.MaxLen {
		fmt.Errorf("Слишком длинная строка: %s", text)
	}
	return nil
}

// Formatter (тот кто преобразовывает нашу строку)
type Formatter struct{}

// У него нет никаких свойств (полей)
// Теперь нам надо описать этой структуре метод Format:

// Эта функция относится к структуре Formatter - (f Formatter)
// Называется Format
// Принимает на вход строку (text string)
// И в результате возвращает строку обернутую в квадратные скобки
func (f Formatter) Format(text string) string {
	return "[" + text + "]"
}

// Printer (тот кто печатает строку)
type Printer struct{}

// У него нет никаких свойств (полей)
// Теперь нам надо описать этой структуре метод Print:

// Эта функция относится к структуре Printer - (p Printer)
// Называется Print
// Принимает на вход строку (text string)
// И печатает ее на экран
func (p Printer) Print(text string) {
	fmt.Println(text)
}

// Все, мы описали все обработчики, теперь мы можем их инициализировать и использовать

func main() {
	// 1) Создаем создателя в переменную "creator"
	// И сразу задаем ему символы из которых он будет создавать строки
	creator := Creator{
		Chars: []rune{'a', 'b', 'c', 'd', 'e'},
	}

	// 2) Создаем обработчик который проверяет длину строк в переменную "processor"
	// И сразу задаем ему максимально разрешенное кол-во символов
	processor := Processor{
		MaxLen: 3,
	}

	// 3) Создаем преобразователь строк в переменную "formatter"
	// У него нет никаких полей по-этому прсто создаем его пустой
	formatter := Formatter{}

	// 4) Создаем принтер в переменную "printer"
	// У него нет никаких полей по-этому прсто создаем его пустой
	printer := Printer{}

	// Теперь у нас в переменных creator processor formatter printer
	// Существуют объекты типов Creator Processor Formatter Printer
	// Теперь мы можем пользоваться их методами
	// Например у переменнной creator есть метод Create()
	// Мы можем его вызывать через точку creator.Create()

	// 1. создаём строку, результат записываем в переменную myString
	myString := creator.Create(3) // получится "abc"

	// 2. проверяем длину нашей строки
	// Если процессор вернул ошибку завершаем нашу программу и печатаем сообщение(Fatalf)
	err := processor.Process(myString)
	if err != nil {
		log.Fatalf("Ошибка %v", err)
	}

	// 3. форматируем строку и результат записываем в переменную myFormattedString
	myFormattedString := formatter.Format(myString) // получится "[abc]"

	// 4. печатаем сначала строку myString
	// а потом еще печатаем форматированную строку myFormattedString
	printer.Print(myString)
	printer.Print(myFormattedString)

}
