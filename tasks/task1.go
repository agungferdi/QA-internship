package tasks

import (
	"fmt"
	"math/rand"
	"time"
)

func random() int {
	return rand.Intn(2)
}

func Task1() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomLocal := func() int {
		return r.Intn(2)
	}

	results := make(map[int]int)
	iterations := 10000

	for i := 0; i < iterations; i++ {
		result := exactBiasedRandom(randomLocal)
		results[result]++
	}

	fmt.Println("Task 1: Create a function with biased probability")
	fmt.Printf("0 occurred: %d times (%.2f%%)\n", results[0], float64(results[0])*100/float64(iterations))
	fmt.Printf("1 occurred: %d times (%.2f%%)\n", results[1], float64(results[1])*100/float64(iterations))
	fmt.Println("Expected: 0 (75%), 1 (25%)")
	fmt.Println()
}

func exactBiasedRandom(randomFunc func() int) int {
	r := randomFunc() + 2*randomFunc()

	if r < 3 {
		return 0
	}
	return 1
}
