package main

import "log"

func init() {
	log.SetPrefix("Trace: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}
func main() {
	//写到标准的日志记录器
	log.Println("Message:")
	//Fatalln 在调用Println之后会调用 os.Exit(1)
	log.Fatalln("Fatalln Message")
	//panicln 会在Println之后会直接调用 panic()
	log.Panicln("Panic Message")

}
