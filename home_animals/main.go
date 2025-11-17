package main

import (
	"fmt"
)

// Pet - структура со списком действий животных,
// которые они делают в течение дня
type Pet struct {
	Name string
	Type string
	Actions []string // список действий, которые сделал питомец в течение дня 
	Score int // уровень активности
} 

// p - это параметр функции
// тип параметра Pet
// через p мы обращаемся к структуре только в рамках этой функции
func CalculateScore(p Pet) Pet {
	// Создаем переменную score, счетчик активности
	score := 0
	for _, action := range p.Actions {
		// Выбор действия
		switch action {
		case "run":
			score += 3
		case "eat":
			score += 1
		case "sleep":
			score += 0  
		}
	}
	// После завершения цикла мы записываем итоговое score внутрь структуры Pet
	p.Score = score
	// Возвращаем обновленную структуру 
	return p 
}

// Функция которая красиво выводит поля
func PrintPet(p Pet) {
	fmt.Println("Имя:",p.Name)
	fmt.Println("Вид:",p.Type)
	fmt.Println("Действия:",p.Actions)
	fmt.Println("Уровень активности:",p.Score)
} 

func main() {
	// Создаем двух животных pt1 и pt2
	pt1 := Pet{
		Name: "Шарик",
		Type: "Собака",
		Actions: []string{"run","eat","sleep"},
	}
	pt2 := Pet{
		Name: "Миу",
		Type: "Кошка",
		Actions: []string{"run","eat","sleep"},
	}
	// Передаём структуру в функцию, которая вернёт новую
	pt1 = CalculateScore(pt1)  
	pt2 = CalculateScore(pt2)
	
	// Печатаем итог 
	PrintPet(pt1)
	PrintPet(pt2) 
}
