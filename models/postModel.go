package models

type Post struct {
	gorm.model
	Title string
	Body  string
}
