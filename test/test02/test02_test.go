package test02

import (
	"net/http"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func TestDownload(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.goinggo.net/feeds/posts/default?alt=rss",
			http.StatusOK,
		},
		{
			"http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			http.StatusNotFound,
		},
	}
	t.Log("需要下载不同url不同内容")
	{
		for _, u := range urls {
			t.Logf("检查地址\"%s\"和状态码\"%d\"", u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\t取得url内容失败",err, ballotX)
				}
				t.Log("取得url内容成功", checkMark)
				//关闭请求
				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\t应该是 \"%d\" 状态. %v",
						u.statusCode, checkMark)
				} else {
					t.Errorf("\t\t应该是 \"%d\" 状态. %v %v",
						u.statusCode, ballotX, resp.StatusCode)
				}
			}
		}
	}
}
