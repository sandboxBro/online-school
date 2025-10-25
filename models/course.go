package models

// Cours - курсы: название, описание, длительность, стоимость и спиок уроков.
// Также содержит иерархию: Parent и Children
type Course struct {
	Title         string  `json:"title"`  // название
	Description   string   // описание
	DurationWeeks int      `json:"duration_weeks"` // длителоьность в неделях 
	Price         float32  // стоимость обучения 
	Lessons       []Lesson `json:"lessons"`// список уроков (one-to-many)
	
	// Доп. поля
	Tags []string
	IsActive bool 
}

type ShortCourse struct {
	Title string
}