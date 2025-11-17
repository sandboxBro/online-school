package main

import (
	"fmt"
)

// Student — структура со списком оценок
type Student struct {
	Name    string
	Grades  []int
	Average float64
}

// CalculateAverage принимает студента, считает средний балл и возвращает обновлённую структуру
func CalculateAverage(s Student) Student {
	sum := 0
	// цикл по оценкам
	for _, grade := range s.Grades {
		sum += grade
	}
	// считаем среднее
	if len(s.Grades) > 0 {
		s.Average = float64(sum) / float64(len(s.Grades))
	}
	return s
}

// PrintStudent выводит данные студента
func PrintStudent(s Student) {
	fmt.Println("Имя:", s.Name)
	fmt.Println("Оценки:", s.Grades)
	fmt.Println("Средний балл:", s.Average)
}

func main() {
	// создаём студента
	st := Student{
		Name:   "Иван",
		Grades: []int{5, 4, 3, 5, 5, 4},
	}
	// Передаём структуру в функцию, которая вернёт новую
	st = CalculateAverage(st)
	// Печатаем итог
	PrintStudent(st)
}
