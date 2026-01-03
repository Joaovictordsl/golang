package main

import "fmt"

func romanToInt(s string) int {

	result := 0

	valores := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		// Acede ao valor usando o byte do caractere atual

		valorAtual := valores[s[i]]

		if i < len(s)-1 && valorAtual < valores[s[i+1]] {
			result = result - valorAtual

		} else {

			result = result + valorAtual
		}

		fmt.Println("Convertendo", s, "em", result)
	}
	return result

	// 	// for i := 0; i < len(s); i++ {
	// 	// 	valorAtual := valores[s[i]]

	// 	// }

	// 	// result := 0

	// 	// for _, algarismo := range s {

	// 	// 	fmt.Printf(" %c\n", algarismo)
	// 	// }

}
