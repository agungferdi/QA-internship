package tasks

import (
	"fmt"
)

func Task5() {
	testGain := []int{-5, 1, 5, 0, -7}

	result := findHighestAltitude(testGain)

	fmt.Println("Task 5: Find the highest altitude")
	fmt.Printf("Altitude gain: %v\n", testGain)
	fmt.Printf("Highest altitude: %d\n", result)
	fmt.Println()
}

func findHighestAltitude(gain []int) int {
	currentAltitude := 0
	highestAltitude := 0

	for _, altitudeChange := range gain {
		currentAltitude += altitudeChange

		if currentAltitude > highestAltitude {
			highestAltitude = currentAltitude
		}
	}

	return highestAltitude
}
