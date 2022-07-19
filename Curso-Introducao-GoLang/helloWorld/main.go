package main

import "fmt"

func main() {
	// uma função pode retornar mais de um valor
	// a construção := cria a variável casa ela não exista, já definindo seu tipo
	ok, err := say("Hello World") // sintaxe sugar
	if err != nil {
		fmt.Println(err.Error())
	}
	switch ok {
	case true:
		fmt.Println("Deu certo!!!")
	default:
		fmt.Println("Deu errado!!!")
	}
}

/*
var ok bool
var err error
ok, err = say("Hello World")
*/


// as funções devem declarar o tipo de cada variável que recebe ou que retorna
func say(what string) (bool, error) {
	if what == "" {
		return false, fmt.Errorf("Empty string")
	}
	fmt.Println(what)
	return true, nil
}