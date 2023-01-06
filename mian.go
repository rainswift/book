package main

import (
	"book/common"
	"book/manager"
	"book/structs"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"time"
)

func weibo() {

}

func main() {
	//common.WeiboH("https://weibo.com/ajax/getNavConfig")
	one()
}

func one() {
	dom, _ := common.GetHttp("http://www.b5200.org/")
	dom.Find(".nav a").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(i)
		href, _ := s.Attr("href")
		text := s.Text()
		utf8Data, _ := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(text)) //gbk 转 utf-8
		fmt.Println(string(utf8Data), href)
		fmt.Println(href)
		two(href)
	})
}

//分类-列表
func two(url string) {
	dom, err := common.GetHttp(url)
	if err != nil {
		return
	}

	dom.Find(".ll .item").Each(func(i int, s *goquery.Selection) {
		title := utf(s.Find("dl dt a").Text())
		author := utf(s.Find("dl dt span").Text())

		href, _ := s.Find("dl dt a").Attr("href")
		img, _ := s.Find(".image img").Attr("src")
		introduce := utf(s.Find("dl dd").Text())
		obj := structs.BookInfo{
			Title:     title,
			Author:    author,
			Href:      href,
			Introduce: introduce,
			Img:       img,
		}
		manager.Mgr.Save(&obj)
		manager.Mgr.CreateZZ(title)
		//fmt.Println(title,author,href,img,introduce)
		zhangj(href, title)
	})
}

//章节信息
func zhangj(url string, name string) {
	dom, err := common.GetHttp(url)
	if err != nil {
		return
	}
	len := dom.Find("#list dd").Length()
	if len == 0 {
		return
	}
	//list := structs.BookList{}
	dom.Find("#list dd").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Find("a").Attr("href")
		t := time.NewTimer(1 * time.Second)
		<-t.C
		go zhangCentent(href, name)

		//title := utf(s.Find("dd").Text())
		//obj := structs.Chapter{Title: title}
		//manager.Mgr.SaveZZ(name,&obj)
	})
}

//章节内容
func zhangCentent(url string, name string) {
	//t := time.NewTimer(1*time.Second)
	//<-t.C
	dom, err := common.GetHttp(url)
	if err != nil {
		return
	}
	title := dom.Find(".bookname h1").Text()
	content := dom.Find("#content").Text()
	fmt.Println(utf(title), utf(content))
	obj := structs.Chapter{Title: utf(title), Chapter: utf(content)}

	manager.Mgr.SaveZZ(name, &obj)
}

func utf(text string) string {
	utf8Data, _ := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(text))
	return string(utf8Data)
}
