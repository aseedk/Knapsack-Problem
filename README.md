# Knapsack Problem Solver

This Go program implements a dynamic programming solution for the classic 0/1 Knapsack Problem. The goal is to maximize the total value of items placed in a knapsack without exceeding its weight capacity.

## Problem Description

The 0/1 Knapsack Problem is defined as follows:

- Given a set of items, each with a weight and a value.
- You need to determine the maximum value that can be obtained by selecting a subset of these items, such that the total weight does not exceed a given capacity.

### Example Input

- **Items**:
  - Item-1: Value = 5, Weight = 2
  - Item-2: Value = 6, Weight = 3
  - Item-3: Value = 10, Weight = 4
- **Knapsack Capacity**: 6

### Output

The program prints detailed steps of how the solution is built and provides the final best value that can be obtained for the given capacity.

## How the Program Works

1. **Initialization**:
   - The program initializes a 2D array `records` where `records[i][j]` represents the maximum value obtainable with the first `i` items and a knapsack capacity `j`.
  
2. **Dynamic Programming Logic**:
   - If either `i` (number of items) or `j` (capacity) is zero, `records[i][j]` is set to 0.
   - If the current item's weight exceeds the current capacity `j`, the value from the previous row is used.
   - Otherwise, the program compares:
     - The best value from the previous row (`records[i-1][j]`).
     - The sum of the current item's value and the best value for the remaining capacity (`values[i-1] + records[i-1][j-weights[i-1]]`).
   - The maximum of these two values is stored in `records[i][j]`.

3. **Output**:
   - The program prints detailed information about each step and decision.
   - The final 2D `records` array is displayed in a tabular format.
   - The best possible value for the given capacity is shown at the end.

## Output Example

### Initial Data


### Detailed Steps
The program outputs each decision-making step, including when a value is set to 0, when an item's weight is too large for the current capacity, or when the value is chosen based on the best option.

### Final Output


The best value for a capacity of 6 is 15.

## How to Run

1. Ensure you have [Go installed](https://golang.org/doc/install) on your system.
2. Save the code in a file named `main.go`.
3. Open your terminal and navigate to the directory containing `main.go`.
4. Run the program using:

   ```bash
   go run main.go
