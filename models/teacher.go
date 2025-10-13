package models

// Teacher - преподаватель: поля разных типов + вложенная Contact
type Teacher struct {
	FirstName string       // имя
	LastName  string       // фамилия
	Age       int          // возраст
	ExperienceYears int    // стад работы (в годах)
	Contact   Contact      // вложенная структура 
	// Дополнительяные поля разных типов для примера
	IsOnLine   bool        // работает ли удаленно 
	Rating     float32     // рейтинг преподователя
	Skills     []string    // список навыхов 
}