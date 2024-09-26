package main

import (
	"fmt"
	"strings"
)

func main() {

	texto01 := "Generation Brasil"
	texto02 := "generation brasil"
	texto03 := "GENERATION BRASIL"

	// Len
	fmt.Printf("\nO tamanho da string texto01 é: %d", len(texto01))

	// Contains
	fmt.Printf("\nO texto generation, existe na string texto02? %t", strings.Contains(texto02, "generation"))
	fmt.Printf("\nO texto Go, existe na string texto03? %t", strings.Contains(texto03, "Go"))

	// ContainsAny
	fmt.Printf("\nPelo menos um caractere do texto 'aeiou', existe na string texto01? %t", strings.ContainsAny(texto01, "aeiou"))
	fmt.Printf("\nPelo menos um caractere do texto 'xyz', existe na string texto01? %t", strings.ContainsAny(texto01, "xyz"))

	// HasPrefix e HasSuffix
	fmt.Printf("\nO texto02 inicia com a string gen? %t", strings.HasPrefix(texto02, "gen"))
	fmt.Printf("\nO texto02 termina com a string sil? %t", strings.HasSuffix(texto02, "sil"))

	// ToUpper e ToLower
	fmt.Printf("\nString texto02 em letras maiúsculas: %s", strings.ToUpper(texto02))
	fmt.Printf("\nString texto03 em letras minúsculas: %s", strings.ToLower(texto03))

}
