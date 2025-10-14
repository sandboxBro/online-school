package models

// Cours - курсы: название, описание, длительность, стоимость и спиок уроков.
// Также содержит иерархию: Parent и Children
type Course struct {
	Title         string   // название
	Description   string   // описание
	DurationWeeks int      // длителоьность в неделях 
	Price         float32  // стоимость обучения 
	Lessons       []Lesson // список уроков (one-to-many)
	
	// Доп. поля
	Tags []string
	IsActive bool 
}