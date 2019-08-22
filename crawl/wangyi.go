package main

import (
	"gostudy/crawl/bs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TopMusic struct {
	Status int `json:"status"`
	CopyrightId int  `josn:"copyrightId"`
	Mvid int  `json:"mvid"`
	CommentThreadId string `json:"commentThreadId"`
	TransNames []string `json:"transNames"`
	Score  float64  `json:"score"`
	Fee  int `json:"fee"`
	Ftype  int  `json:"ftype"`
	No int  `json:"no"`
	PublishTime int64 `json:"publishTime"`
	Type int `json:"type"`
	Duration int  `json:"duration"`
	Alias  []string `json:"alias"`
	Privilege  Privilege `json:"privilege"`
	Djid int `json:"djid"`
	Album  Album `json:"album"`
	Name string `json:"name"`
	Id int `json:"id"`
	astRank int  `json:"lastRank"`
}

type Privilege struct {
	St      int  `json:"st"`
	Flag    int  `json:"flag"`
	Subp    int  `json:"subp"`
	Fl      int  `json:"fl"`
	Fee     int  `json:"fee"`
	Dl      int  `json:"dl"`
	Cp      int  `json:"cp"`
	PreSell bool `json:"presell"`
	Cs      bool `json:"cs"`
	Toast   bool `json:"toast"`
	Maxbr   int  `json:"maxbr"`
	Id      int  `json:"id"`
	Pl      int  `json:"pl"`
	Sp      int  `Json:"sp"`
	Payed   int  `json:"payed"`
}

type Album struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	PicUrl  string   `json:"picUrl"`
	Tns     []string `json:"tns"`
	Pic_str string   `json:"pic_str"`
	Pic     int64    `json:"pic"`
}

//https://github.com/yitimo/api-163-go
//网易云热歌榜
func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://music.163.com/discover/toplist?id=3778678", nil)
	if err != nil {
		fmt.Println("获取热歌榜失败！")
		return
	}
	//设置请求头，防止504
	req.Header.Set("Referer", "https://music.163.com/")
	req.Header.Set("Host", "music.163.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取热歌榜失败！")
		return
	}
	fmt.Println("获取热歌榜成功！")
	soup := bs.Init(string(body))
	// 获取属性 class="f-hide" 的 ul 标签
	for _, j := range soup.Sel("ul", &map[string]string{"class": "f-hide"}) {
		// 找出子标签为 a 的标签
		for k, v := range j.SelByTag("a") {
			for _, url := range *v.Attrs {
				fmt.Println(int(k)+1, url, v.Value)
			}
		}
	}
	// 获取属性 class="n-cmt" 的 div 标签  
	for _, j := range soup.Sel("div", &map[string]string{"class": "n-cmt"}) {
		// 获取评论id 总评论数
		for _, tid := range *j.Attrs {
			fmt.Println(tid)
		}
	}
	// 获取 textarea 标签 id 为 song-list-pre-data 
	// 得到热歌榜详情
	for _, j := range soup.SelById("song-list-pre-data") {
		fmt.Println(j.Value)
		var data []TopMusic
		if err := json.Unmarshal([]byte(j.Value), &data); err != nil {
			panic(err)
		}
		for _, item := range data {
			fmt.Println(item.Name)
		}
	}
}
