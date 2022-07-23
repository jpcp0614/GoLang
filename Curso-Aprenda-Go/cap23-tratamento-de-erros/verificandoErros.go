/*
Na minha religião, underscore é pecado.
    Verifique seus erros!
    (Exceção: fmt.Println)
    Na prática:
        Exemplo 0: fmt.Println
        Exemplo 1: fmt.Scan(&var)
        Exemplo 2: os.Create → strings.NewReader → io.Copy
        Exemplo 3: os.Open → io.ReadAll
*/

package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

// Println
// func main() {
// 	n, err := fmt.Println("Hello World!")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(n)
// }

// Scan
// func main() {
// 	var answer1, answer2, answer3 string

// 	fmt.Print("Name: ")
// 	_, err := fmt.Scan(&answer1)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Print("Comida Favorita: ")
// 	_, err = fmt.Scan(&answer2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Print("Esporte favorito: ")
// 	_, err = fmt.Scan(&answer3)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(answer1, answer2, answer3)
// }

// Create
// func main() {
// 	f, err := os.Create("names.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer f.Close()
// 	defer fmt.Println("Rodou a função que estava em defer")

// 	r := strings.NewReader("James Bond")

// 	io.Copy(f, r)
// }

// Open
func main() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	defer fmt.Println("Rodou a função que estava em defer")

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}