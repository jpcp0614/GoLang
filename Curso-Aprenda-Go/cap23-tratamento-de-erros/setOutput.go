package main

import (
	"fmt"
	"os"
	"log"
)

// Create
func main() {
	f1, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f1.Close()
	log.SetOutput(f1)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("Error:", err)
	}
	defer f2.Close()

	fmt.Println("Verificar os arquivos log.txt")
}

// criou o arquivo log.txt com o erro dentro