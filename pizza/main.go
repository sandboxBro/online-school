package main

import (
	"fmt"
	"log"
)

func prepareDough(flour int, water int) (int, error) {
	// Функция "prepareDough" — готовит тесто
	// Принимает на вход 2 параметра:
	// flour int - количество муки (пускай будет в граммах)
	// water int - количество воды (пускай будет в милилитрах, не суть)
	// Эта функция смешивает два ингридиента и возвращает нам общий вес получившегося теста int
	fmt.Println("Готовим тесто...")
	dough := flour + water
	if water > flour {
		return dough, fmt.Errorf("Тесто получилось слишком жидкое.")
	}

	return dough, nil // возвращаем количество теста
}

// func addIngredients(dough int) (int, int) {
func addIngredients(dough int, cheese int, tomato int) int {
	// Функция "addIngredients" — добавляет начинку
	// Прнимает на вход количество теста и смешивает его с ингридиентами
	// Смешивает все ингридиенты
	// Возвращает вес готовой смеси и количество добавленного сыра
	fmt.Println("Добавляем начинку...")
	//cheese := 100 // грамм
	//tomato := 50  // грамм
	readyMass := dough + cheese + tomato

	// fmt.Println("Всего ингридиентов добавлено:", "\nТесто:", dough, "g", "\nСыр:", cheese, "g", "\nТоматы:", tomato, "g")
	// fmt.Print - печать при помощи форматированной строки с использованием %d
	fmt.Printf(
		"Всего ингридиентов добавлено:\nТесто: %dg\nСыр: %dg\nТоматы: %dg\n",
		dough,
		cheese,
		tomato,
	)

	return readyMass
	// return readyMass, cheese // возвращаем 2 значения: общую массу и количество сыра

}

func bakePizza(mass int, cheese int) string {
	// Функция "bakePizza" — печёт пиццу и проверяет результат
	// Эта функция принимает на вход вес общей массы и кол-во сыра добавленное на предыдущем шагу
	// Проверяет количество сыра и возвращает строук сформированную при помощи fmt.Sprintf
	// Т.е. мы сделали return и вернули результат вызова fmt.Sprintf - эта функция возвращает string значит и наша функция возвращает string
	fmt.Println("Выпекаем пиццу...")

	if cheese > 80 {
		// fmt.Sprintf("Пицца получилась отличная! Вес: %dg", mass)
		return fmt.Sprintf("Пицца получилась отличная! Вес: %d", mass)
	}

	// return fmt.Sprintf("Пицца так себе, мало сыра. Вес: %dg", mass)
	return fmt.Sprintf("Пицца так себе, мало сыра. Вес: %dg", mass)
}

func slicePizza(slice int) {
	fmt.Printf("Пицца порезана на %d кусков\n", slice)
}

// Главная функция "main"
func main() {
	fmt.Println("Начинаем готовить пиццу!")
	// 1. Печатаем на экран строку "Начинаем готовить пиццу!" с перреводом строки на следующую
	flour := 56
	water := 34
	dough, err := prepareDough(flour, water)
	if err != nil {
		log.Fatalf(err.Error())
	}

	cheese := 100 // грамм
	tomato := 50  // грамм
	// 2. Вызываем функцию prepareDough передав в нее какое-то количество муки и воды
	// Результат этой функции записываем в переменную dough(тесто)
	// mass, cheese := addIngredients(dough)
	mass := addIngredients(dough, cheese, tomato)
	// 3. Дальше нам это тесто надо смешать с начинкой
	// Вызываем функцию addIngredients передаем в нее тесто
	// Результат этой функции записываем в переменные mass и cheese
	message := bakePizza(mass, cheese)
	// 4. Запускаем готовку пиццы вызвав функцию bakePizza и передаем в нее нужные параметры
	// Эту функция нам возвращает строку, надо записать эту строку в переменную message и затем напечатать его на экран
	d:= 10
	slicePizza(d)


	fmt.Println(message)

}
