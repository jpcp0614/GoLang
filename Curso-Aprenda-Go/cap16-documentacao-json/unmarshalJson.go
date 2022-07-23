/*
- JSON-to-Go
- Tags
- Marshal/unmarshal vs. encoder/decoder
    - Marshal vai pra uma variável
    - Encoder "vai direto"

	{"Nome":"James","Sobrenome":"Bond","Idade":32,"Profissao":"Agente Secreto"}
	{"Nome":"Anakin","Sobrenome":"Skywalker","Idade":44,"Profissao":"Sith"}
*/

package main

import (
	"fmt"
	"encoding/json"
)

type Informacoes struct {
	Nome				string `json:"Nome"`
	Sobrenome		string `json:"Sobrenome"`
	Idade				int    `json:"Idade"`
	Profissao		string `json:"Profissao"`
}

func main() {
	sb := []byte(`[
		{"Nome":"James","Sobrenome":"Bond","Idade":32,"Profissao":"Agente Secreto"},
		{"Nome":"Anakin","Sobrenome":"Skywalker","Idade":44,"Profissao":"Sith"}
	]`)

	var info []Informacoes
	err := json.Unmarshal(sb, &info)
	if err != nil { fmt.Println("Error SB:", err) }

	fmt.Printf("%v - %T", info[0].Nome, info)
}