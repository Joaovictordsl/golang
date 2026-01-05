package main

func binary(array []int, num int) int {

	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := array[mid]

		if guess == num {
			return mid

		} else if guess > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
