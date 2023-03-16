package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string
	Email string
	Cpf   int
}

func (c Cliente) Exibe() {
	fmt.Println("Exibindo dentro do método:", c.Nome)
}

type ClienteInternacional struct {
	Cliente
	Pais string `json:"pais"`
}

func main() {
	cliente := Cliente{
		Nome:  "Fulano",
		Email: "ze@meio.com",
		Cpf:   12312343332,
	}
	fmt.Println(cliente)
	cliente2 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "John",
			Email: "john@mail.com",
			Cpf:   12312312300,
		},
		Pais: "EUA",
	}
	// Não precisa colocar cliente2.Cliente.Nome... mas até pode.
	fmt.Printf("Nome: %s E-mail: %s Cpf: %d País: %s\n", cliente2.Cliente.Nome, cliente2.Email, cliente2.Cpf, cliente2.Pais)
	fmt.Println(cliente2)

	// O método do struts se comporta como se fosse uma herança
	cliente.Exibe()
	cliente2.Exibe()

	jsonCliente, err := json.Marshal(cliente2)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(jsonCliente)
	fmt.Println(string(jsonCliente))

	jsonClienteNovo := `{"Nome":"John","Email":"john@mail.com","Cpf":12312312300,"pais":"EUA"}`
	cliente4 := ClienteInternacional{}
	json.Unmarshal([]byte(jsonClienteNovo), &cliente4)

	fmt.Println(cliente4)

}
