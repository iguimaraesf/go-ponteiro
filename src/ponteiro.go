package main

import "fmt"

type Carro struct {
	Nome string
}

// Como não passa um ponteiro, a alteração na variável é somente local.
func (c Carro) andou() {
	c.Nome = "Sonic"
	fmt.Printf("%s andou\n", c.Nome)
}

func main() {
	a := 10

	// endereço de memória onde a variável "a" está armazenada
	println(&a)

	var ponteiro *int = &a

	// o valor armazenado em um ponteiro
	println(*ponteiro)

	// muda o valor da variável "a"
	*ponteiro = 50

	println(a, *ponteiro)

	// não retorna nada, mas passa o endereço de memória da variável e altera o conteúdo dela.
	aaa := 1
	teste(&aaa)

	fmt.Println(aaa)

	carro := Carro{
		Nome: "Ká",
	}
	carro.andou()
	println(carro.Nome)
}

func teste(a *int) int {
	*a = 999
	return *a
}
