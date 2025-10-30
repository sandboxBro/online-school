package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Данные для подключения
	connStr := "host=127.0.0.1 port=5432 user=postgres password=postgres-14 dbname=school sslmode=disable"

	// Открываем соединение
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка при подключении", err)
	}
	defer db.Close()

	// Проверяем соединение
	err = db.Ping()
	if err != nil {
		log.Fatal("БД недоступна", err)
	}

	fmt.Println("Подключение к базе данных успешно!")

	// Пробуем выполнить простой sql-запрос
	rows, err := db.Query("SELECT id, first_name, last_name, grade FROM school.students;")
	if err != nil {
		log.Fatal("Ошибка при выполнении запроса", err)
	}
	defer rows.Close()

	fmt.Println("Список студентов:")
	for rows.Next() {
		var id int 
		var firstName, lastName, grade string 

		err = rows.Scan(&id, &firstName, &lastName, &grade)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d | %s %s %s\n", id, firstName, lastName, grade) 
		fmt.Println("-----")
	}
    
	// sql.Open(...)    - Настраивает соединение с базой данных (но не подключается сразу).
    // db.Ping()        - Проверяет, что соединение реально работает.
    // db.Query(...)    - Выполняет SQL-запрос и возвращает результат.
    // rows.Scan(...)   - Считывает каждую строку результата в переменные Go.
    // defer db.Close() - Закрывает соединение, когда программа завершится.
}