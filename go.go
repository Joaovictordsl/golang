package main

import "fmt"

//pega o resto da divisao por 10 e armazena numa variavel
//pega o inteiro que tem e divide por 10, e multi por 10
//pega o valor multiplicado e soma com o resto

// func isPalindrome(x int) bool{

// 	x

// 	return
// }

func main() {
	original := 323
	ultimo_digito := 0
	invertido := 0

	ultimo_digito = original % 10
	invertido = (invertido * 10) + ultimo_digito
	original = original / 10

	fmt.Println(ultimo_digito)
	fmt.Println(invertido)
	fmt.Println(original)

}
