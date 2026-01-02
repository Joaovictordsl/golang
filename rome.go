package main

import "fmt"

func romanToInt(s string) {

	valores := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	fmt.Println(valores)

	// 	// for i := 0; i < len(s); i++ {
	// 	// 	valorAtual := valores[s[i]]

	// 	// }

	// 	// result := 0

	// 	// for _, algarismo := range s {

	// 	// 	fmt.Printf(" %c\n", algarismo)
	// 	// }

}
