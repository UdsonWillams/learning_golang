package main

import (
	"cachorro"
	"fmt"
)

// Função main chama a função HumanEarsToDog para saber quantos anos um cachorro possui.
func main() {
	fmt.Println(cachorro.HumanEarsToDog(10))
}

// - Crie um package "cachorro".
//     - Este package deverá exportar uma função Idade,
// que toma como parâmetro um número de anos e retorna a idade equivalente em anos caninos. (1 ano humano → 7 anos caninos)
//     - Documente seu código com comentários, e utilize a função Idade na sua função main.
// - Rode seu programa para verificar se ele funciona.
// - Rode um local server com godoc e leia sua documentação.