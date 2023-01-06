package common

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
)

func WeiboH(url string) {
	//obj := {
	//	"accept": "application/json, text/plain, */*",
	//	"accept-encoding": "gzip, deflate, br",
	//	"accept-language": "zh-CN,zh;q=0.9",
	//	"client-version": "v2.37.21",
	//	"cookie": "UOR=www.google.com.hk,weibo.com,www.google.com.hk; SINAGLOBAL=8775497468164.082.1660742210939; PC_TOKEN=73ce6d06f2; login_sid_t=772b16dcd9271a19eda4edf4992bd693; cross_origin_proto=SSL; WBStorage=4d96c54e|undefined; wb_view_log=2048*11521.25; _s_tentry=passport.weibo.com; Apache=3781580880872.408.1672654301359; ULV=1672654301363:2:1:1:3781580880872.408.1672654301359:1660742211029; SUB=_2AkMU7iLFf8NxqwJRmP4RzW7lao9wyQ_EieKistMeJRMxHRl-yj92qmlStRB6P24MKqCdhAaNBdUvgLEd_mscZt5k0jvO; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9WFaZYQ8VZxWL4NyXnWniOeV; XSRF-TOKEN=aQVZfAV2OUIwOPwNZe9s5APC; WBPSESS=kErNolfXeoisUDB3d9TFH32pG8kDygsiSTbUK3zgLH0iAvo9M-i-0X61y__Dznj7j5WFNvikyEpuGwwg4x10ObycPGdo_ZZ8fzt_4HUUzssQl21so2efXHV9PiO_pgUCqwZONsj8TOjw6iKXq194IV1AqjpmqhBb14gaXJQGHR4=",
	//	"referer": "https://weibo.com/newlogin?tabtype=weibo&gid=102803&openLoginLayer=0&url=https%3A%2F%2Fweibo.com%2F%3Ftopnav%3D1%26mod%3Dlogo",
	//	"sec-ch-ua": "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"",
	//	"sec-ch-ua-mobile": "?0",
	//	"sec-ch-ua-platform": "\"Windows\"",
	//	"sec-fetch-dest": "empty",
	//	"sec-fetch-mode": "cors",
	//	"sec-fetch-site": "same-origin",
	//	"server-version": "v2022.12.30.4",
	//	"traceparent": "00-6034e6ff608aa9c862183a0eccd1a16a-fc7ecd37aa3ad524-00",
	//	"user-agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
	//	"x-requested-with": "XMLHttpRequest",
	//	"x-xsrf-token": "aQVZfAV2OUIwOPwNZe9s5APC"
	//}
	client := &http.Client{}
	//newUrl := strings.Replace(url, "http://", "https://", 1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	var cookie = "UOR=www.google.com.hk,weibo.com,www.google.com.hk; SINAGLOBAL=8775497468164.082.1660742210939; PC_TOKEN=73ce6d06f2; login_sid_t=772b16dcd9271a19eda4edf4992bd693; cross_origin_proto=SSL; WBStorage=4d96c54e|undefined; wb_view_log=2048*11521.25; _s_tentry=passport.weibo.com; Apache=3781580880872.408.1672654301359; ULV=1672654301363:2:1:1:3781580880872.408.1672654301359:1660742211029; SUB=_2AkMU7iLFf8NxqwJRmP4RzW7lao9wyQ_EieKistMeJRMxHRl-yj92qmlStRB6P24MKqCdhAaNBdUvgLEd_mscZt5k0jvO; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9WFaZYQ8VZxWL4NyXnWniOeV; XSRF-TOKEN=aQVZfAV2OUIwOPwNZe9s5APC; WBPSESS=kErNolfXeoisUDB3d9TFH32pG8kDygsiSTbUK3zgLH0iAvo9M-i-0X61y__Dznj7j5WFNvikyEpuGwwg4x10ObycPGdo_ZZ8fzt_4HUUzssQl21so2efXHV9PiO_pgUCqwZONsj8TOjw6iKXq194IV1AqjpmqhBb14gaXJQGHR4="
	req.Header.Add("cookie", cookie)
	//req.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("user-agent", "user-agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	Body, err := ioutil.ReadAll(resp.Body)

	type resultInfo struct {
		Background string `json:"--w-badge-background"`
	}

	type resultArr resultInfo

	type Response struct {
		Skin     resultArr `json:"skin"`
		LogoType string    `json:"logoType"`
	}

	var response Response

	json.Unmarshal([]byte(Body), &response)
	fmt.Println(response.Skin.Background)

	fmt.Printf("%s", Body)
	//fmt.Println(data)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("错误")
	}

	defer resp.Body.Close()
}

func GetHttp(url string) (*goquery.Document, error) {

	client := &http.Client{}
	//newUrl := strings.Replace(url, "http://", "https://", 1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	//var cookie = "sid=d19ccebd-93c3-40ae-b017-d65f41cba91b; ec=zNcFTxkc-1658817194050-b7d228d5a8a51758227824; FSSBBIl1UgzbN7NO=564GYo5tJiXt55eztYvOGlSoFDzpqoOP5D_oYsGSWmC6wYyoh0dYCIK2GNQlrLixOdluCaL_7sHdzrmh.9MpdLA; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1664351595,1665192272; hcbm_tk=MjZvkSuzEcEn+h+FqtBa3HV6hH8Eg3NlNBy1wbKLYaXhR48uEfwHgTui9pLYwh9B1J4pLW+vA62NCscxJEubtJ6coFvhiM/qTxB6GihTBu03bZ0vO4dPYhkMnCJGDnSx7990b3vZu7S6Ik9/4G0pK2jr0wpCIetEEBa5MzvSW21fxyuQEeUn5vLJmGuRYWarBL4NUuWShM5ZnpkkIaMLAw5JdSABXZMoJCFWlAoapShNe9iwhRjmzlK+hLqDdDC8Jwok7ZiOe6soZ3D1xAL/XuGM4j2b7Q2V8UmSpEFQONZ5CwVeXb6tn+f1UKqM+n3DV9Lp4QF6hLjhV2RROpbI8DWIqHpaSiMnRVVVSBWqxmLc86SwwK5lgNbMy7iunrvwGD5zisWI75W8PDxs7PxcCcTVABQRcxqWZU5tQCfV_RERXV4AmIa4VjY7e; _exid=jj2s4bUUvSutjZKye41pBIrFHIcLV9QkotT5MtBq%2Bbcs7hBSbSQJ6Hq6rQPEFHDgLPFXY8c%2BKAKaJGE5OslQ8A%3D%3D; _efmdata=b6SRTGbgMTft%2BK8GBf8Vsm%2B0LIYOqUdrbyQFOKRTj6PxqKOb67LGM6r%2Fk3Hsz%2BHJttSWxnnXY0ywZZnTxLzTdFcx07Tuba8EFcKTBVB%2FH%2BI%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1665459167; FSSBBIl1UgzbN7NP=530s24DJly8AqqqDkS1EbEAbEMZomZKukcAxAcXvjCFLNu7mlMrPaN698eShYzZyKWroxtV0QPfRvLiWRRKILJNJo3ZKBxujW2qLfYvkgwFhRQDeMEOTcRodo358s_NaEKzI9idt3e9yNg7S2LrvSZmobRBa_edlt.GsEkDa34SUfik2bRCmpE_gaV8Z_yIgG754JsYhRPSi9mji3c6ipT5DnnArwDhrhO0O2lxrd0XpZpzufLJx3CotbdHtWV2QRQ"
	//req.Header.Add("cookie", cookie)
	//req.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("user-agent", "user-agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("错误")
		return nil, fmt.Errorf("Error: StatusCode is %d\n", resp.StatusCode)
	}

	defer resp.Body.Close()

	//res, err := http.Get(url)
	//if err != nil {
	//	log.Fatal(err)
	//}
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	dom, err := goquery.NewDocumentFromReader(resp.Body)
	return dom, err
}
