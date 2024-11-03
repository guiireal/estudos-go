package main

import (
	"fmt"
	"os"
	"strconv"
)

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	n := make([]int, len(numbers))
	copy(n, numbers)

	pivotIndex := len(n) / 2
	pivot := n[pivotIndex]
	n = append(n[:pivotIndex], n[pivotIndex+1:]...)

	less := make([]int, 0)
	greater := make([]int, 0)

	for _, number := range n {
		if number <= pivot {
			less = append(less, number)
		} else {
			greater = append(greater, number)
		}
	}

	return append(
		append(quicksort(less), pivot),
		quicksort(greater)...,
	)
}

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))

	for i, value := range input {
		number, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("%s não é um número válido\n", value)
			os.Exit(1)
		}

		numbers[i] = number
	}

	fmt.Println(quicksort(numbers))
}
