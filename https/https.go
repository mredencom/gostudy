package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "hello world %s", request.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", handler)
	_ = http.ListenAndServe(":8089", nil)
}
