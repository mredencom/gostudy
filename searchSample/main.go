package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"

	//_ "matchers"

	"github.com/goinaction/code/chapter2/sample/search"
	//"search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program. main是整个搜索程序入口
func main() {
	// Perform the search for the specified term. 使用特定的项做搜索
	search.Run("president")
}
