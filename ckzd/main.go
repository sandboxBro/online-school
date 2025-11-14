package main

import "fmt"

func main() {
	// Выводим числа от 1 до 10
	fmt.Println("Задание1: Вывод чисел от 1 до 10")
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}

	// Вывести четные числа от 2 до 20
	fmt.Println("\nЗадание2: Вывод четных чисел от 2 до 20")
	for i := 2; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Printf(" %d", i)
		}
	}

	// Посчитать сумму чисел от 1 до 100
	fmt.Println("\nЗапдание3: Посчет суммы чисел от 1 до 100")
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		// sum = sum + i
	}
	fmt.Println("Сумма чисел от 1 до 100 =", sum)

	// Пропусти все числа меньше 5
	fmt.Println("\nЗадание4: Пропускаем все числа меньше 5")
	for i := 1; i <= 10; i++ {
		if i < 5 {
			continue // пропустить итерацию
		}
		fmt.Printf(" %d", i)
	}

	// Останови цикл на числе 7
	fmt.Println("\nЗадание5: Останови цикл на числе 7")
	for i := 1; i <= 10; i++ {
		if i == 7 {
			fmt.Println(" Стоп")
			break
		}
		fmt.Printf(" %d", i)
	}

	// Счетчик четных чисел
	fmt.Println("\nЗадание 6: Счетчик четных чисел")
	count := 0 // Переменная для подсчета четных чисел
	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			count++ // увеличиваем счетчик
		}
	}
	fmt.Println("Количество четных чисел:", count)

	// Вывести элементы списка имен
	fmt.Println("\nЗадание 7: Выводим элементы списка имен")
	names := []string{"Анна", "Борис", "Вика", "Гриша"}
	for _, name := range names {
		fmt.Printf(" %s", name)
	}

	// Найти самое длинное имя из списка
	fmt.Println("\nЗадание 8: Самое длинное имя из списка")
	names = []string{"Анна", "Борис", "Вика", "Гриша"}
	max := names[0]
	for _, name := range names {
		if len(name) > len(max) {
			max = name
		}
	}
	fmt.Println("Самое длинное имя:", max)

	// Подсчитай общее количество букв во всех именах
	fmt.Println("\nЗадание 9: Общее колицество букв во вусех именах")
	names = []string{"Анна", "Борис", "Вика", "Гриша"}
	count = 0
	for _, name := range names {
		// name := []rune(name)
		count += len([]rune(name)) // прибавляем длинну каждого имени
	}
	fmt.Println("Общая дина имен", count)

	// Таблица умножения
	fmt.Println("\nЗадание 10: Таблица умножения")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}

	// Нарисуй прямоугольник из звёздочек
	fmt.Println("\nЗадание 11: Рисуем прямоугольник из звездочек")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

	// Сумма всех чисел до первого, делящегося на 13
	fmt.Println("\nЗадание 12: Сумма всех чисел до первого, делящегося на 13")
	sum = 0
	for i := 1; i <= 100; i++ {
		if i%13 == 0 {
			break
		}
		sum += i
	}
	fmt.Println("Сумма всех чисел до первого, делящегося на 13:", sum)

	// Обратный отчет
	fmt.Println("\nЗадание 13: Обратный отчет")
	for i := 10; i >= 1; i-- {
		if i > 1 {
			fmt.Printf("%d", i)
			continue
		}
		fmt.Println("Cтарт!")
	}

	// Подсчет гласных в слове
	fmt.Println("\nЗадание 14: Подсчет гласных в слове")
	word := "Програмированние"
	vowels := []rune{'а', 'е', 'ё', 'и', 'о', 'у', 'ы', 'э', 'ю', 'я'}
	count = 0
	for _, letters := range word {
		for _, i := range vowels {
			if letters == i {
				count++
				break
			}
		}
	}
	fmt.Println("Колличество гласных букв", count)
}
