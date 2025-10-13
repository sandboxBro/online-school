package models

// Lesson - урок: заголовок, номер,описание и преподователь
type Lesson struct {
	Title         string     // заголовок
	Number        int        // порядковый номер
	Description   string     // текстовое описание
	Teacher       Teacher   // преподаватель (указатель чтобы разделять объекты)
	IsFree        bool       // пример булевого поля
}