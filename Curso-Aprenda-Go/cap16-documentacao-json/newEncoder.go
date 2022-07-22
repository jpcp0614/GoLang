package main

import (
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

	// no encoder vai direto, não precisa salvar numa variável para depois imprimir na tela
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(jamesBond)
}