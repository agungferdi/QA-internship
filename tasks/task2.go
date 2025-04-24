package tasks

import (
	"fmt"
)

func Task2() {
	testStr := "19827439111498912111"
	result := mostFrequentDigit(testStr)

	fmt.Println("Task 2: Find the most frequent digit")
	fmt.Printf("Input string: %s\n", testStr)
	fmt.Printf("Most frequent digit: %d\n", result)
	fmt.Println()
}

func mostFrequentDigit(str string) int {
	frequency := make(map[int]int)

	for _, char := range str {
		digit := int(char - '0')
		frequency[digit]++
	}

	maxFreq := 0
	mostFreqDigit := 0

	for digit, freq := range frequency {
		if freq > maxFreq {
			maxFreq = freq
			mostFreqDigit = digit
		}
	}

	return mostFreqDigit
}
