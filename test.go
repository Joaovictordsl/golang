package main

import "fmt"

func main() {

	str := []string{"Ola", "teste", "eu"}
	result := []string{}

	encontrados := make(map[string]bool)

	// fmt.Printf("%c\n", str[0][1])

	// result = append(result, string(str[0][0]))

	for i := 0; i < len(str); i++ {
		primeiraLetra := string(str[i][0])

		if !encontrados[primeiraLetra] {
			result = append(result, primeiraLetra)
			encontrados[primeiraLetra] = true
		}

	}

	fmt.Println(result)
}
