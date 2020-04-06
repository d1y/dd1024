package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func request(av *string) string {
	plate := "DAVK-042"
	if av != nil {
		plate = *av
	}
	site := fmt.Sprintf("https://javzw.com/%s", plate)
	fmt.Println("开始播放车牌号: ", plate)
	// return ""
	resp, err := http.Get(site)
	if err != nil {
		log.Fatal("请求失败: ", av, err)
	}
	jquery, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("解析失败: ", err)
	}
	var m3u8 string = ""
	jquery.Find("#my-player").Each(func(i int, wrapper *goquery.Selection) {
		m3u8, _ = wrapper.Find("source").Attr("src")
	})
	if len(m3u8) == 0 {
		log.Fatal("未找到播放地址, 车牌号: ", plate)
	}
	defer resp.Body.Close()
	return m3u8
}
