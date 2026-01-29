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

	fmt.Printf("container = %v", container)
	fmt.Printf("  pivot index = %d, pivot value = %d", pivotIndex, pivot)

	var leftPivotAmount, pivotAmount int

	for _, v := range container {
		switch {
		case v < pivot:
			leftPivotAmount++
		case v == pivot:
			pivotAmount++
		}
	}

	firstRightPivotIndex := leftPivotAmount + pivotAmount

	if leftPivotAmount == 0 { // bad luck, pivot is a leftmost element in sorted sequence
		swap(&container[0], &container[pivotIndex])
		QuickSort(container[1:containerLen])
		return
	}

	if firstRightPivotIndex == containerLen { // bad luck, pivot is a rightmost element in sorted sequence
		swap(&container[containerLen-1], &container[pivotIndex])
		QuickSort(container[0 : containerLen-1])
		return
	}

	fmt.Printf("  leftPivotAmount = %d, pivot amount = %d, firstRightPivotIndex = %d\n", leftPivotAmount, pivotAmount, firstRightPivotIndex)

	left := container[0:firstRightPivotIndex]             //elements on the left + pivots
	right := container[firstRightPivotIndex:containerLen] //elements on the right of pivot

	for i, jRight := 0, 0; i < firstRightPivotIndex; i++ {
		if left[i] < pivot {
			continue //left[i] can stay in its position on the left of pivot
		}

		for k := jRight; k < len(right); k++ {
			if right[k] < pivot {
				swap(&left[i], &right[k])
				jRight = k + 1 //index of element of the right from pivot slice, starting from which there still might be elements to be moved to the left of pivot
				break
			}
		}
	}

	for i := firstRightPivotIndex; i < pivotAmount; i++ {
		container[i] = pivot //pivot elements are placed to their place. there might be more then one pivot
	}

	left = container[0:leftPivotAmount] //now left points only to elements on the left of pivot (pivots are excluded)
	QuickSort(left)                     //apply quick sort to the left  from pivot part of container
	QuickSort(right)                    //apply quick sort to the right from pivot part of container
}

func main() {
	fmt.Println("Quick sort")

	collection := []int{13, 4, 35, 104, 28, 6, 12, 200, 1, -5, 44, 55, 11, 99, -9, 21, -100, 67, -37, 12}
	fmt.Printf("Initial container = %v\n", collection)
	QuickSort(collection)
	fmt.Printf("Sorted container  = %v\n", collection)
}
