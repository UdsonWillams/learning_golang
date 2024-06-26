// - Partindo do código abaixo, ordene a []int e a []string.

package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi) // N retorna nda, apenas ordena o valor
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs) // N retorna nda, apenas ordena o valor
	fmt.Println(xs)
}
