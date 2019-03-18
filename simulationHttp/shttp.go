package main

import (
	"encoding/json"
	"fmt"
	"github.com/panjf2000/ants"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	//"sync/atomic"
	"time"
)

var sum int32

//定义一个常量域名
//const DOMAIN = "https://sdk.mokatyper.com/games"
//const DOMAIN = "http://172.18.70.52:9999/games"
const DOMAIN = "http://www.baidu.com"

/*****************************************start*************************************************/
const (
	VERSION = "1.0.1"
	AKEY    = "baiwenji123456789"
	WSR     = `{"scene":1001,"path":"","query":{},"referrerInfo":{},"__public":{"path":"","query":{},"scene":1001,"referrerInfo":{}}}`
)

//花费的时间
func Sm() int {
	return rand.Intn(200)
}

//系统打点参数
func Sys() string {
	rand.Seed(r())
	var sys []string
	sys = append(sys, `{"model":"iPhone 7 Plus","pixelRatio":3,"windowWidth":414,"windowHeight":736,"system":"iOS 10.0.1","language":"zh","version":"6.6.3","screenWidth":414,"screenHeight":736,"SDKVersion":"2.2.4","brand":"devtools","fontSizeSetting":16,"benchmarkLevel":1,"batteryLevel":100,"statusBarHeight":20,"platform":"devtools","devicePixelRatio":3}`)
	sys = append(sys, `{"model":"iPhone 8 Plus","pixelRatio":3,"windowWidth":414,"windowHeight":736,"system":"iOS 10.0.1","language":"zh","version":"6.6.3","screenWidth":414,"screenHeight":736,"SDKVersion":"2.2.4","brand":"devtools","fontSizeSetting":16,"benchmarkLevel":1,"batteryLevel":100,"statusBarHeight":20,"platform":"devtools","devicePixelRatio":3}`)
	sys = append(sys, `{"model":"iPhone X","pixelRatio":3,"windowWidth":414,"windowHeight":736,"system":"iOS 10.0.1","language":"zh","version":"6.6.3","screenWidth":414,"screenHeight":736,"SDKVersion":"2.2.4","brand":"devtools","fontSizeSetting":16,"benchmarkLevel":1,"batteryLevel":100,"statusBarHeight":20,"platform":"devtools","devicePixelRatio":3}`)
	sys = append(sys, `{"model":"iPhone XR","pixelRatio":3,"windowWidth":414,"windowHeight":736,"system":"iOS 10.0.1","language":"zh","version":"6.6.3","screenWidth":414,"screenHeight":736,"SDKVersion":"2.2.4","brand":"devtools","fontSizeSetting":16,"benchmarkLevel":1,"batteryLevel":100,"statusBarHeight":20,"platform":"devtools","devicePixelRatio":3}`)
	le := len(sys)
	random := rand.Intn(le)
	var left int
	var right int
	if random == le {
		right = random
		left = random - 1
	} else if random == 0 {
		left = random
		right = random + 1
	} else {
		left = random
		right = random + 1
	}
	var sy string
	for _, v := range sys[left:right] {
		sy = string(v)
	}
	return sy
}
func Env() string {
	rand.Seed(r())
	env := []string{"init", "show", "hide", "event", "ofo"}
	le := len(env)
	random := rand.Intn(le)
	var left int
	var right int
	if random == le {
		right = random
		left = random - 1
	} else if random == 0 {
		left = random
		right = random + 1
	} else {
		left = random
		right = random + 1
	}
	//fmt.Println(env)
	//fmt.Println(le, random, left, right, env[left:right])
	var str string
	for _, v := range env[left:right] {
		str = v
	}
	//fmt.Println(le, random, left, right, net[left:right])
	return str
}

func Net() string {
	rand.Seed(r())
	net := []string{"wifi", "2g", "3g", "4g", "unknown"}
	le := len(net)
	random := rand.Intn(le)
	var left int
	var right int
	if random == le {
		right = random
		left = random - 1
	} else if random == 0 {
		left = random
		right = random + 1
	} else {
		left = random
		right = random + 1
	}
	var str string
	for _, v := range net[left:right] {
		str = v
	}
	//fmt.Println(le, random, left, right, net[left:right])
	return str
}

func It() string {
	itt := strconv.Itoa(int(time.Now().UnixNano() / 1e6))
	rand7 := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	return itt + rand7
}

func Rt() string {
	rrt := strconv.Itoa(int(time.Now().UnixNano() / 1e6))
	rand7 := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	return rrt + rand7
}

func Uuid() string {
	return "bd87e2b6eb70b744d9f3fd3be8ba2685"
}

func St() int {
	return int(time.Now().UnixNano() / 1e6)
}

func Ver() string {
	return VERSION
}

func Akey() string {
	return AKEY
}

//启动参数
func Wsr() string {
	return WSR
}
func Ifo(e string) bool {
	if e == "init" {
		return true
	} else {
		return false
	}
}

//请求序号
func Rqn() int {
	return 1
}

//小游戏发生错误次数
func Ec() int {
	return 0
}

//自定义事件名称
func Tp() string {
	return "mt_banner_show"
}

//自定义事件参数
func Ct() string {
	return "banner广告展示事件"
}

//小游戏从进入前台到退入后台的时长，单位毫秒，只有在小游戏进入后台时会发送这个请求，即 env = "hide" 时。
func Dt() int {
	return rand.Intn(1000)
}

func Ufo() {
	//s := sufo{
	//	nickName:  nickName(),
	//	gender:    gender(),
	//	language:  language(),
	//	city:      city(),
	//	province:  province(),
	//	country:   country(),
	//	avatarUrl: avatarUrl(),
	//}
	s := sufo{
		NickName:  nickName(),
		Gender:    gender(),
		Language:  language(),
		City:      city(),
		Province:  province(),
		Country:   country(),
		AvatarUrl: avatarUrl(),
	}
	js, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(js))
}

type sufo struct {
	NickName  string
	Gender    int
	Language  string
	City      string
	Province  string
	Country   string
	AvatarUrl string
}

func nickName() string {
	return "sam"
}

func gender() int {
	return 1
}
func language() string {
	return "zh-CN"
}
func city() string {
	return "suzhou"
}
func province() string {
	return "jiangsu"
}
func country() string {
	return "China"
}
func avatarUrl() string {
	return "https://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83eof4h3DFf3d7WkgjlgLWaLpSVu1G44t6cxBibibkM8kKicXTM4S4ib9oO6QHoPct9ljMjKgP1ba7FXfhA/132"
}

func r() int64 {
	rand7 := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	a, _ := strconv.Atoi(rand7)
	return int64(a)
}

/*****************************************end*************************************************/
/**
拼接参数
*/
func paramsFilter() string {
	tp := Tp()
	sm := strconv.Itoa(Sm())
	sys := Sys()
	env := Env()
	it := It()
	rt := Rt()
	st := strconv.Itoa(St())
	uuid := Uuid()
	ver := Ver()
	akey := Akey()
	wsr := Wsr()
	ifo := "false"
	rqn := strconv.Itoa(Rqn())
	ec := strconv.Itoa(Ec())
	v := url.Values{}
	v.Add("tp", tp)
	v.Add("sm", sm)
	v.Add("ec", ec)
	v.Add("env", env)
	v.Add("it", it)
	v.Add("rt", rt)
	v.Add("st", st)
	v.Add("uuid", uuid)
	v.Add("ver", ver)
	v.Add("akey", akey)
	v.Add("wsr", wsr)
	v.Add("ifo", ifo)
	v.Add("rqn", rqn)
	v.Add("sys", sys)
	body := v.Encode()
	//m, _ := url.ParseQuery(body)
	//fmt.Println(m)
	//url := fmt.Sprintf("%s?tp=%s&sm=%s&ec=%s&env=%s&it=%s&rt=%s&st=%s&uuid=%s&ver=%s&akey=%s&wsr=%s&ifo=%s&rqn=%s&sys=%s", DOMAIN, tp, sm, ec, env, it, rt, st, uuid, ver, akey, wsr, ifo, rqn,sys)
	url := fmt.Sprintf("%s?%s", DOMAIN, body)
	//fmt.Println(url)
	return url
}
func sendHttp() {
	//统计请求
	//n := i.(int32)
	//atomic.AddInt32(&sum, n)

	//client := &http.Client{}
	url := paramsFilter()
	//url:=`https://sdk.mokatyper.com/games?sm=121&sys={"model":"iPhone 7 Plus","pixelRatio":3,"windowWidth":414,"windowHeight":736,"system":"iOS 10.0.1","language":"zh","version":"6.6.3","screenWidth":414,"screenHeight":736,"SDKVersion":"2.2.4","brand":"devtools","fontSizeSetting":16,"benchmarkLevel":1,"batteryLevel":100,"statusBarHeight":20,"platform":"devtools","devicePixelRatio":3}&env=init&it=15505685118482949502&rt=15505685118485693606&st=1550568511996&uuid=bd87e2b6eb70b744d9f3fd3be8ba2685&ver=1.0.1&akey=baiwenji123456789&wsr={"scene":1001,"path":"","query":{},"referrerInfo":{},"__public":{"path":"","query":{},"scene":1001,"referrerInfo":{}}}&ifo=false&rqn=1&ec=0`
	//fmt.Println(url)
	//req, err := http.NewRequest("GET", url, nil)
	req, err := http.Get(url)
	//设置响应头
	//req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	//..req.Header.Add("Accept", "text/html,application/xhtml+xm…plication/xml;q=0.9,*/*;q=0.8")
	if err != nil {
		//fmt.Println(i)
		panic("get 出错了")
	}

	//处理返回结果
	//response, err := client.Do(req)
	//关闭请求
	defer req.Body.Close()
	//_, err = httputil.DumpResponse(response, true)
	if err != nil {
		panic("错误了----------------------------------")
	}
	status := req.StatusCode
	fmt.Printf("请求状态：%s\n", status)
	//fmt.Printf(s, status, sum)
}
func main() {
	//fmt.Println(paramsFilter())
	//os.Exit(0)
	t1 := time.Now().Unix()
	runTimes := 1000

	// 使用普通池
	var wg sync.WaitGroup
	//使用协程池
	p, _ := ants.NewPoolWithFunc(8, func(i interface{}) {
		sendHttp()
		wg.Done()
	})
	//最后销毁协程池
	defer p.Release()
	//循环提交任务
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	t2 := time.Now().Unix()
	fmt.Printf("运行协程: %d\n", p.Running())
	fmt.Printf("结果 %d\n", sum)
	fmt.Printf("使用时间 %d\n", t2-t1)
}
