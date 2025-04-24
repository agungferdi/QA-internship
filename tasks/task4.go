package tasks

import (
	"fmt"
)

func Task4() {
	testArray := []string{"apple", "boy", "cat", "dog", "elephant"}
	testElement1 := "car"
	testElement2 := "cat"

	result1 := elementExists(testElement1, testArray)
	result2 := elementExists(testElement2, testArray)

	fmt.Println("Task 4: Check if element exists in sorted array")
	fmt.Printf("Array: %v\n", testArray)
	fmt.Printf("\"%s\" exists in array: %t\n", testElement1, result1)
	fmt.Printf("\"%s\" exists in array: %t\n", testElement2, result2)
	fmt.Println()
}

func elementExists(element string, array []string) bool {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := left + (right-left)/2

		if array[mid] == element {
			return true
		}

		if array[mid] < element {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
