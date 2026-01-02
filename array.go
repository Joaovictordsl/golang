package main

func Digitize(n int) []int {
	// your code here
	result := []int{}

	if n == 0 {
		result = append(result, 0)
		return result
	}

	for n > 0 {
		last_digit := n % 10
		result = append(result, last_digit)
		n = n / 10
	}

	return result
}

// func Digitize2(n int) []int {

// 	result := []int{}
// 	array := strconv.Itoa(n)

// 	for _, s := range array {

// 	}

// 	return result
// }
