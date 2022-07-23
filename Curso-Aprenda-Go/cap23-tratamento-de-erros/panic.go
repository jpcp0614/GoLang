package main

import (
	"os"
)


func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
}

/*
panic: open no-file.txt: no such file or directory

goroutine 1 [running]:
main.main()
	/home/jp/Documents/GoLang/Curso-Aprenda-Go/cap23-tratamento-de-erros/panic.go:11 +0x49
exit status 2
*/