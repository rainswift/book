package structs

import "gorm.io/gorm"

//书籍信息
type BookInfo struct {
	gorm.Model
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Href        string   `json:"href"`
	Introduce   string   `json:"Introduce"`
	Img         string   `json:"img"`
	ChapterList []string `json:"chapterList"`
}

type BookList struct {
	BookList []BookInfo
}
