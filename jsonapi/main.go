package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-school/models"
)

func coursHandler(w http.ResponseWriter, r *http.Request) {
	t1 := models.Teacher{
		FirstName:       "Виктор",
		LastName:        "Дулин",
		Age:             35,
		ExperienceYears: 10,
		Contact:         models.Contact{Email: "viktor/dl@mail.ru", Phone: "8-433-907-85-65"},
		IsOnLine:        true,
		Rating:          4.5,
		Skills:          []string{"Go", "Databases", "Algorithms"},
	}

	t2 := models.Teacher{
		FirstName:       "Варвара",
		LastName:        "Владимировна",
		Age:             46,
		ExperienceYears: 6,
		Contact:         models.Contact{Email: "vr.vara@mail.ru", Phone: "8-906-778-52-69"},
		IsOnLine:        true,
		Rating:          4.7,
		Skills:          []string{"HTML", "CSS", "JS"},
	}

	// Создаем уроки, привязываем преподователей
	l1 := models.Lesson{Title: "Введение в Go", Number: 1, Description: "Обзор языка", Teacher: t1}
	l2 := models.Lesson{Title: "Типы и переменные", Number: 2, Description: "Базовые типы", Teacher: t1}
	l3 := models.Lesson{Title: "Горутины", Number: 3, Description: "Параллельность в Go", Teacher: t1}
	l4 := models.Lesson{Title: "Основы HMT", Number: 1, Description: "Структура документа", Teacher: t2}
	l5 := models.Lesson{Title: "CSS: стили", Number: 2, Description: "Оформление страниц", Teacher: t2}

	// Создаем курсы
	Coursego := models.Course{
		Title:         "Go: с нуля до PRO",
		Description:   "Полный курс по Go",
		DurationWeeks: 9,
		Price:         25000,
		Lessons:       []models.Lesson{l1, l2, l3},
		Tags:          []string{"backend", "golang"},
		IsActive:      true,
	}

	frontendCourse := models.Course{
		Title:         "Frontend базовый",
		Description:   "HTML и CSS для начинающих",
		DurationWeeks: 7,
		Price:         13500,
		Lessons:       []models.Lesson{l4, l5},
		Tags:          []string{"frontend", "web"},
		IsActive:      true,
	}

	// Создаем школу
	_ = models.School{
		Name:     "WEB City",
		Website:  "https://webcityworld.com",
		Courses:  []models.Course{frontendCourse, Coursego},
		Teachers: []models.Teacher{t1, t2},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Coursego)
}

func main() {
	http.HandleFunc("/course", coursHandler)
	fmt.Println("Сервер запущен на http://127.0.0.1:8080 ...")
	// Запуск сервера на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
