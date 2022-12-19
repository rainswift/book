package structs

import "gorm.io/gorm"

type BookInfo struct {
	gorm.Model
	Title  string `json:"title"`
	Author  string `json:"author"`
	Href   string `json:"href"`
	Introduce string `json:"Introduce"`
	Img      string `json:"img"`
}

type BookList struct {
	BookList []BookInfo
}