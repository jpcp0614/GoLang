/*
Callback: passe uma função como argumento a outra função.
*/

package main

import (
	"fmt"
)

func argumento(f func()) {
	f()
}

func ola() {
	fmt.Println("Olá")
}

func main() {
	argumento(ola)
}