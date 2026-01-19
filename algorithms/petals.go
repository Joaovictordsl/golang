package main

import "fmt"

func HowMuchILoveYou(i int) string {
	counter := 0

	arr := []string{"I love you", "a little", "a lot", "passionately", "madly", "not at all"}

	if i == 1 {
		return arr[0]
	}

	for x := 0; x < i; x++ {

		if counter == len(arr) {
			counter = 0
		}
		counter++

	}
	counter--

	fmt.Println(counter)
	return arr[counter]
}

func main() {
	fmt.Println(HowMuchILoveYou(7))
}
