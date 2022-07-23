/*
 uma maneira de encriptar senhas utilizando hashes.
- x/crypto/bcrypt
    - GenerateFromPassword
    - CompareHashAndPassword
- Sem Go Playground!
    - go get golang.org/x/crypto/bcrypt
*/

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt" // não funciona de jeito nenhum
)


const (
	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

func main() {
	senha := "minhaSenhaMuitoSegura2022"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), DefaultCost)
	if err != nil { fmt.Println("Error SB:", err)}

	fmt.Println(sb)
}

