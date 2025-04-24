# Internship Programming Tasks

This repository contains solutions to five programming tasks implemented in Go as part of an internship application. Each task is implemented in a separate file with a clean, modular structure.

## Project Structure

```
internship/
├── go.mod             # Go module definition
├── main.go            # Main entry point that runs all tasks
├── README.md          # This file
└── tasks/             # Directory containing individual task implementations
    ├── task1.go       # Task 1: Biased Random Function
    ├── task2.go       # Task 2: Most Frequent Digit
    ├── task3.go       # Task 3: Second Highest Salary
    ├── task4.go       # Task 4: Element Exists in Sorted Array 
    ├── task5.go       # Task 5: Highest Altitude
    └── ss results/    # Directory containing screenshot of results
        └── Screenshot 2025-04-24 at 16.47.27.png
```

## Tasks Description

### Task 1: Biased Random Function
Implements a function that returns 0 with 75% probability and 1 with 25% probability using a provided random function that returns 0 and 1 with 50% probability each.

### Task 2: Most Frequent Digit
Finds the most frequent digit in a string containing only digits.

### Task 3: Second Highest Salary
Finds the second highest salary from an array of different salaries with O(n) complexity.

### Task 4: Element Exists in Sorted Array
Checks if a given string element exists in a sorted array of strings with O(log n) complexity using binary search.

### Task 5: Highest Altitude
Calculates the highest altitude reached during a road trip based on an array of altitude gains.

## How to Run

1. Make sure you have Go installed on your machine
2. Navigate to the project directory
3. Run the program with:
   ```
   go run main.go
   ```

## Results

All tasks have been successfully implemented and tested. The screenshot below shows the output of running all five tasks:

![Task Results](/Users/ferdiansyahmuhammadagung/Projects/amartha/internship/tasks/ss%20results/Screenshot%202025-04-24%20at%2016.47.27.png)

The results demonstrate that:
- Task 1: The biased random function correctly returns 0 with ~75% probability and 1 with ~25% probability
- Task 2: The most frequent digit in "19827439111498912111" is correctly identified as 1
- Task 3: The second highest salary in [5000, 3000, 6000, 1000] is correctly identified as 5000
- Task 4: The element existence checker correctly implements binary search for O(log n) complexity
- Task 5: The highest altitude calculator correctly identifies 1 as the highest altitude

## Implementation Details

All tasks are implemented in Go with attention to:
- Clean code practices
- Efficient algorithms
- Proper error handling
- Modular design with separate files for each task

Each task function is designed to be reusable and follows the specifications provided in the problem statements.