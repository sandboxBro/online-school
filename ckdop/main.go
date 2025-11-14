package main

import "fmt"

func main() {
	// Пример 1. Самый простой цикл с числом
	fmt.Println("Пример 1: Счетчик от 1 до 5")
	for i := 1; i <= 5; i++ {
		fmt.Println("Шаг", i)
	}

	// Пример 2. Цикл с сумированием
	fmt.Println("\nПример 2: Суммируем числа от 1 до 5")
	sum := 0 
	for i := 1; i <= 5; i++ {
		sum = sum + 1
	    fmt.Println("После шага", i, "сумма =", sum)
	}
	fmt.Println("Итоговая сумма:", sum)

	// Пример 3. Проход по слайсу
	fmt.Println("\nПример 3: Проходим по списку имён")
	names := []string{"Анна", "Борис", "Вика"}
	for i := 0; i < len(names); i++ {
		fmt.Println("Индекс:", i, "Значение:", names[i])
	}

	// Пример 4. Упрощенный цикл по слайсу с range
	fmt.Println("\nПример 4: То же самое, но с range")
	for i, name := range names {
		fmt.Println("Индекс:", i, "Имя:", name)
	}

	// Пример 5. Используем break и continue
	fmt.Println("\nПример 5. break и continue")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Println("Пропускаем шаг 3")
			continue // пропустить итерацию
		}
		if i == 5 {
			fmt.Println("Дошли до 5 — выходим из цикла")
			break // завершить цикл
		}
		fmt.Println("Шаг", i)
	}

}