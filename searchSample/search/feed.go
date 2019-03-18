package search

import (
	"encoding/json"
	"fmt"
	"os"
)

const dataFile = "data.json"
//const dataFile = "/d/goworkspace/gostudy/searchSample/data/data.json"

// Feed contains information we need to process a feed.
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshals the feed data file.
func RetrieveFeeds() ([]*Feed, error) {
	// Open the file. 打开文件
	fmt.Println(dataFile)
	file, err := os.Open(dataFile)
	//fmt.Printf("%s", file)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once 当数据返回时
	// the function returns. 关闭文件
	defer file.Close()

	// Decode the file into a slice of pointers. 将文件解码到一个slice中
	// to Feed values. 这个切片的每一项向一个Feed类型值的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// We don't need to check for errors, the caller can do this. 这个函数不需要检查错误.用调用者会做这件事
	return feeds, err
}
