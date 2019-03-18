package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

type Sufo struct {
	NickName  string
	Gender    int
	Language  string
	City      string
	Province  string
	Country   string
	AvatarUrl string
}

//func (s * Student)ShowStu() {
//	fmt.Println("show Student :")
//	fmt.Println("\tName\t:", s.Name)
//	fmt.Println("\tAge\t:", s.Age)
//	fmt.Println("\tGuake\t:", s.Guake)
//	fmt.Println("\tPrice\t:", s.Price)
//	fmt.Printf("\tClasses\t: ")
//	for _, a := range s.Classes {
//		fmt.Printf("%s ", a)
//	}
//	fmt.Println("")
//}

func main() {
	//st := Student{
	//	"Xiao Ming",
	//	16,
	//	true,
	//	[]string{"Math", "English", "Chinese"},
	//	9.99,
	//}
	st := Sufo{
		"sam",
		1,
		"zh-CN",
		"苏州",
		"江苏",
		"China",
		"222",
	}

	fmt.Println("before JSON encoding :")
	//st.ShowStu()

	b, err := json.Marshal(st)
	if err != nil {
		fmt.Println("encoding faild")
	} else {
		fmt.Println("encoded data : ")
		fmt.Println(b)
		fmt.Println(string(b))
	}
	ch := make(chan string, 1)
	go func(c chan string, str string) {
		c <- str
	}(ch, string(b))
	strData := <-ch
	fmt.Println("--------------------------------")
	stb := &Student{}
	//stb.ShowStu()
	err = json.Unmarshal([]byte(strData), &stb)
	if err != nil {
		fmt.Println("Unmarshal faild")
	} else {
		//fmt.Println("Unmarshal success")encoded data :
		//stb.ShowStu()
	}
}
