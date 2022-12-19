package main

import (
	"book/common"
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
		two(href)
	})
}
//分类-列表
func two(url string)  {
	dom,_ := common.GetHttp(url)
	list := structs.BookList{}
	dom.Find(".ll .item").Each(func(i int, s *goquery.Selection) {
		title := utf(s.Find("dl dt a").Text())
		author :=  s.Find("dl dt span").Text()

		href , _ :=  s.Find("dl dt a").Attr("href")
		img , _ :=  s.Find(".image img").Attr("src")
		introduce :=  s.Find("dl dd").Text()
		list.BookList = append(list.BookList,structs.BookInfo{
			Title:title,
			Author:author,
			Href:href,
			Introduce: introduce,
			Img: img,
		})
		fmt.Println(title,author,href,img,introduce)
	})
}

func utf(text string) string  {
	utf8Data, _ :=  simplifiedchinese.GBK.NewDecoder().Bytes([]byte(text))
	return string(utf8Data)
}
