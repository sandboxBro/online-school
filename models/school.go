package models

// School - школа: название, сайт, список курсов и список преподавателей
type School struct {
	Name       string      // название
	Website    string      // сайт
	Courses    []Course    // список курсов
	Teachers   []Teacher  // список преподавателей
}