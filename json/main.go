package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero string `json:"numero"`
	Saldo  int    `json:"saldo"`
}

func main() {
	conta := Conta{Numero: "1", Saldo: 100}
	res, err := json.Marshal(conta) // quando eu uso Marshal eu guardo o valor pra mim
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta) // quando eu uso NewEncoder eu pego o valor, faco o processo de serializacao e ja mando pra algum lugar
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"numero":"1","saldo":200}`)
	var contaX Conta
	json.Unmarshal(jsonPuro, &contaX) // quando eu uso Unmarshal eu pego o valor serializado e transformo em um valor que eu posso usar
	println(contaX.Numero, contaX.Saldo)
}
