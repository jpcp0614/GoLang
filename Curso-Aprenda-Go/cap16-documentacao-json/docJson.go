/*
Já entendemos ponteiros, já entendemos métodos. Já temos o conhecimento necessário para começar a utilizar a standard library.
- Nesse vídeo faremos uma orientação sobre como abordar a documentação.
- Essa aula não foi preparada. Vai ser tudo ao vivo no improviso pra vocês verem como funciona o processo.
    - golang.org → Documents → Package Documentation 
    - godoc.org → encoding/json
        - files
        - examples
        - funcs
        - types
        - methods
	
	All public fields, methods and functions starts with uppercase char.
	All private fields, methods and functions starts with lowercase char.
*/


package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Pessoa struct {
	Nome			string
	Sobrenome	string
	Idade			int
	Profissao	string
}

func main() {
	jamesBond := Pessoa {
		Nome: "James",
		Sobrenome: "Bond",
		Idade: 32,
		Profissao: "Agente Secreto",
	}

	darthVader := Pessoa {
		Nome: "Anakin",
		Sobrenome: "Skywalker",
		Idade: 44,
		Profissao: "Sith",
	}

	jb, err := json.Marshal(jamesBond)
	dv, err := json.Marshal(darthVader)
	
	if err != nil {
		fmt.Println("Error JB:", err)
	}

	if err != nil {
		fmt.Println("Error DV:", err)
	}

	// fmt.Println(string(jb))
	// fmt.Println(string(dv))

	os.Stdout.Write(jb)
	os.Stdout.Write(dv)
}