package main

import (
	"fmt"
	"os"
	"log"
)

// Create
func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln("Error:", err)
	}
}

func foo() {
	fmt.Println("Quando osExit() é chamada, a função defer não roda")
}

// 2022/07/23 13:00:21 Error: open no-file.txt: no such file or directory
// exit status 1