package main

import (
	"GO/algorithms"
	"fmt"
)

func main() {

	array := []int{22, 55, 67, 68, 89, 100, 123, 243, 4000, 8054}

	fmt.Println(algorithms.Binary(array, 4000))
}
