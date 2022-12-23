package main

import (
	"book/common"
	"book/manager"
	"book/structs"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func main() {
	one()
}

func one()  {
	dom,_ := common.GetHttp("http://www.b5200.org/")
	dom.Find(".nav a").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(i)
		href, _ := s.Attr("href")
		text := s.Text()
		utf8Data, _ := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(text)) //gbk 转 utf-8
		fmt.Println(string(utf8Data),href)
		fmt.Println(href)
		two(href)
	})
}
//分类-列表
func two(url string)  {
	dom,err := common.GetHttp(url)
	if err != nil {
		return
	}

	dom.Find(".ll .item").Each(func(i int, s *goquery.Selection) {
		title := utf(s.Find("dl dt a").Text())
		author := utf(s.Find("dl dt span").Text())

		href , _ := s.Find("dl dt a").Attr("href")
		img , _ :=  s.Find(".image img").Attr("src")
		introduce := utf(s.Find("dl dd").Text())
		obj := structs.BookInfo{
			Title:title,
			Author:author,
			Href:href,
			Introduce: introduce,
			Img: img,
		}

		fmt.Println(title,author,href,img,introduce)
		zhangj(href,obj)
	})
}
//章节信息
func zhangj(url string,obj structs.BookInfo)  {
	dom,err := common.GetHttp(url)
	if err != nil {
		return
	}
	len := dom.Find("#list dd").Length()
	if len ==0 {
		return
	}
	//list := structs.BookList{}
	list := []string{}
	dom.Find("#list dd").Each(func(i int, s *goquery.Selection) {
		title := utf(s.Text())
		fmt.Println(title)
		list = append(s.Text())
		//list.BookList = append(list.BookList,structs.BookInfo{
		//	Title:obj.Title,
		//	Author:obj.Author,
		//	Href:obj.Href,
		//	Introduce: obj.Introduce,
		//	Img: obj.Img,
		//})
		info := structs.BookInfo{
			Title:     obj.Title,
			Author:    obj.Author,
			Href:      obj.Href,
			Introduce: obj.Introduce,
			Img:       obj.Img,
			ChapterList: lst
		}
		manager.Mgr.Save(&info)
	})
}


func utf(text string) string  {
	utf8Data, _ :=  simplifiedchinese.GBK.NewDecoder().Bytes([]byte(text))
	return string(utf8Data)
}
