package main

import (
	"fmt"
	"net/http"
)

// handler — это функция, которая обрабатывает запросы
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Добро пожаловать в онлайн школу!")
}
func infoHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Онлайн-школа: Go-разработка для начинающих")
}
func main() {
  http.HandleFunc("/", handler) // все запросы на "/" обрабатывает handler
  http.HandleFunc("/info", infoHandler)
  fmt.Println("Сервер запущен на http://127.0.0.1:8080 ...")
  // Запуск сервера на порту 8080
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    fmt.Println("Ошибка запуска сервера:", err)
  }
}
