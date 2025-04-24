package tasks

import (
	"fmt"
)

func Task3() {
	testSalaries := []int{5000, 3000, 6000, 1000}

	result := secondHighestPaidSalary(testSalaries)

	fmt.Println("Task 3: Find the second highest salary")
	fmt.Printf("Input salaries: %v\n", testSalaries)
	fmt.Printf("Second highest salary: %d\n", result)
	fmt.Println()
}

func secondHighestPaidSalary(salaries []int) int {
	if len(salaries) < 2 {
		return -1
	}

	highest := -1
	secondHighest := -1

	for _, salary := range salaries {
		if salary > highest {
			secondHighest = highest
			highest = salary
		} else if salary > secondHighest && salary != highest {
			secondHighest = salary
		}
	}

	return secondHighest
}
