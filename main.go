package main

import (
	"fmt"

	"github.com/alex-salamatov/go_quick_sort/pkg/sorting"
)

func main() {
	fmt.Println("Quick sort")

	collection := []int{-250, 13, 4, 35, 104, 28, 6, 12, 200, 1, -5, 44, 55, 11, 99, -9, 21, -100, 67, -37, 12, -8, 1, 99, 99, 150, 6, -120, -100, 7}
	//collection := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Printf("Initial container = %v\n", collection)
	sorting.QuickSort(collection)
	fmt.Printf("Sorted container  = %v\n", collection)
}
