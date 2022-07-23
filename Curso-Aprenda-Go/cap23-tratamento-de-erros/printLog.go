/*
Opções:
    - fmt.Println() → stdout
    - log.Println() → timestamp + pode-se determinar onde o erro ficará logado
    - log.Fatalln() → os.Exit(1) sem defer
    - log.Panicln() → println + panic → funções em defer rodam; dá pra usar recover
    - panic()
- Recomendação: use log.
- Código: 
    - 1. fmt.Println
    - 2. log.Println
    - 3. log.SetOutput
    - 4. log.Fatalln
    - 5. log.Panicln
    - 6. panic
*/

package main

import (
	// "fmt"
	"os"
	"log"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("Error:", err) // Error: open no-file.txt: no such file or directory
		log.Println("Error:", err) // 2022/07/23 12:48:47 Error: open no-file.txt: no such file or directory
	}
}