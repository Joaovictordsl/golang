package main

import "fmt"

func binary(array []int, num int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := array[mid]

		if guess == num {
			return mid
		} else if num < guess {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}

	return -1
}

func main() {
	x := []int{1, 3, 10, 11, 34, 45, 78, 100, 220, 345, 555}

	fmt.Println(binary(x, 345))
}
