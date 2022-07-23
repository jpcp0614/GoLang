package main

import (
	"os"
	"log"
)


func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Panicln("Error:", err)
	}
}

/*
2022/07/23 13:04:13 Error: open no-file.txt: no such file or directory
panic: Error: open no-file.txt: no such file or directory


goroutine 1 [running]:
log.Panicln({0xc00009af50?, 0xb?, 0x0?})
	/usr/lib/go-1.18/src/log/log.go:399 +0x65
main.main()
	/home/jp/Documents/GoLang/Curso-Aprenda-Go/cap23-tratamento-de-erros/panicln.go:12 +0x74
exit status 2
*/