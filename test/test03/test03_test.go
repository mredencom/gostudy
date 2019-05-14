package test03

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

//feed期望接受的文档
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an object oriented language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>`

//mockServer 返回用来处理请求的服务器指针
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		_, _ = fmt.Fprintf(w, feed)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownload(t *testing.T) {
	statusCode := http.StatusOK
	server := mockServer()
	t.Log("取的测试下载内容.")
	{
		t.Logf("\t检查地址 \"%s\" for 状态 code \"%d\"",
			server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\t获取失败.",
					ballotX, err)
			}
			t.Log("\t\t获取成功.",
				checkMark)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t\t应该接受一个 \"%d\" 状态. %v %v",
					statusCode, ballotX, resp.StatusCode)
			}
			t.Logf("\t\t应该接受一个 \"%d\" 状态. %v",
				statusCode, checkMark)

			var d Document
			if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
				t.Fatal("\t\t应该能够解组响应.",
					ballotX, err)
			}
			t.Log("\t\t应该能够解组响应.",
				checkMark)

			if len(d.Channel.Items) == 1 {
				t.Log("\t\t应该有 \"1\" 节点在 feed.",
					checkMark)
			} else {
				t.Error("\t\t应该有一个节点 \"1\" 节点在feed.",
					ballotX, len(d.Channel.Items))
			}
		}
	}
}

type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}
