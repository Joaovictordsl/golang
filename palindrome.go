package main

//pega o resto da divisao por 10 e armazena numa variavel
//pega o inteiro que tem e divide por 10, e multi por 10
//pega o valor multiplicado e soma com o resto

func isPalindrome(x int) bool {
	original := x
	last_digit := 0
	result := 0

	for x > 0 {
		last_digit = x % 10
		result = (result * 10) + last_digit
		x = x / 10
	}

	if original == result {
		return true

	} else {
		return false
	}
}

// func main() {

// 	fmt.Println(isPalindrome(121))

// }
