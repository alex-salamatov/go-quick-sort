package main

import (
	"fmt"
	"math/rand/v2"
)

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func QuickSort(container []int) {
	if container == nil {
		return
	}

	containerLen := len(container)

	if containerLen <= 1 {
		return
	}

	pivotIndex := rand.IntN(containerLen) //random index of element which we take as pivot
	pivot := container[pivotIndex]

	//fmt.Printf("container = %v original pivot index = %d, pivot value = %d ", container, pivotIndex, pivot)

	var sortedPivotIndex int

	for _, v := range container {
		if v < pivot {
			sortedPivotIndex++
		}
	}

	if pivotIndex != sortedPivotIndex {
		swap(&container[pivotIndex], &container[sortedPivotIndex]) //pivot is placed to its place
	}

	//fmt.Printf("  sortedPivotIndex = %d, N LeftElements = %d N RightElements = %d \n", sortedPivotIndex, sortedPivotIndex, containerLen-sortedPivotIndex-1)

	var left, right []int

	if sortedPivotIndex > 0 {
		left = container[0:sortedPivotIndex] //elements on the left of pivot
	}

	if sortedPivotIndex < containerLen-1 {
		right = container[sortedPivotIndex+1 : containerLen] //elements on the right of pivot
	}

	if left != nil && right != nil {
		//start moving elements from left to right parts if needed
		for i, jRight := 0, 0; i < sortedPivotIndex; i++ {

			if left[i] < pivot {
				continue //left[i] can stay in its position on the left of pivot
			}

			for k := jRight; k < len(right); k++ {

				if right[k] < pivot { //swap current (bigger then pivot) element on the left with first element on the right which is smaller then pivot
					swap(&left[i], &right[k])
					jRight = k + 1 //index of element of the right from pivot slice, starting from which there still might be elements to be moved to the left of pivot
					break
				}
			}
		}
	}

	//	fmt.Printf("Output: %v - %d - %v\n", left, pivot, right)

	QuickSort(left)  //apply quick sort to the left  from pivot part of container
	QuickSort(right) //apply quick sort to the right from pivot part of container
}

func main() {
	fmt.Println("Quick sort")

	collection := []int{-250, 13, 4, 35, 104, 28, 6, 12, 200, 1, -5, 44, 55, 11, 99, -9, 21, -100, 67, -37, 12, -8, 1, 99, 99, 150, 6, -120, -100, 7}
	fmt.Printf("Initial container = %v\n", collection)
	QuickSort(collection)
	fmt.Printf("Sorted container  = %v\n", collection)
}
