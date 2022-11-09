package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
	Rating int    `json:"rating"`
}
