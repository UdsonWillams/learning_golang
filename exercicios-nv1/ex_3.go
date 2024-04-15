// - Utilizando a solução do exercício anterior:
//  1. Em package-level scope, atribua os seguintes valores às variáveis:
//  1. para "x" atribua 42
//  2. para "y" atribua "James Bond"
//  3. para "z" atribua true
//  2. Na função main:
//  1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável.
//     Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o
//     operador curto de declaração.
//  2. Demonstre a variável "s".
package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// Usando o operador curto.
	s := fmt.Sprintf("valores: %v - %v - %v\n", x, y, z)

	// Usando var para setar o valor da variavel.
	var super_s string = fmt.Sprintf("valores: %v - %v - %v\n", x, y, z)

	fmt.Printf("valor de cada variavel era: %v", s)
	fmt.Printf("valor de cada variavel era: %v", super_s)
}

