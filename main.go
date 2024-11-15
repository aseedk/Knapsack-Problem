package main

import (
	"fmt"
	"strings"
)

/*

Item Name | Value | Weight
Item-1    |   5   |   2
Item-2    |   6   |   3
Item-3    |   10  |   4

*/

func main() {
	values := []int{5, 6, 10}
	weights := []int{2, 3, 4}
	count := len(values)
	capacity := 6
	records := make([][]int, count+1)
	for i := range records {
		records[i] = make([]int, capacity+1)
	}
	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println("Weights: ", weights)
	fmt.Println("Values: ", values)
	fmt.Println("Count: ", count)
	fmt.Println("Capacity: ", capacity)
	fmt.Println("-----------------------------------------------------------------------------------------")

	for i := 0; i <= count; i++ {
		fmt.Println("-----------------------------------------------------------------------------------------")
		for j := 0; j <= capacity; j++ {
			fmt.Printf("Index (%d, %d)\n", i, j)
			if i == 0 || j == 0 {
				fmt.Printf("Either I or J is 0.\n")
				fmt.Printf("Set records[%d][%d] to 0\n", i, j)
				records[i][j] = 0
			} else if j < weights[i-1] {
				fmt.Printf("Capacity %d < weight of Item-%d : %d \n", j, i-1, weights[i-1])
				fmt.Printf("Using Best Value of Capacity %d at Index(%d,%d) \n", j, i-1, j)
				fmt.Printf("Storing Value: %d \n", records[i-1][j])
				records[i][j] = records[i-1][j]
			} else {
				fmt.Printf("Capacity %d >= weight of Item-%d : %d \n", j, i-1, weights[i-1])
				fmt.Printf("Either using best value at this capacity at index(%d,%d) = %d\n", i-1, j, records[i-1][j])
				fmt.Printf("Or Adding Value of Item-%d: %d with by filling remaining capacity at Index(%d,%d)\n", i-1, values[i-1], i-1, j-weights[i-1])
				records[i][j] = max(records[i-1][j], values[i-1]+records[i-1][j-weights[i-1]])
			}
			fmt.Println("-----------------------------------------------------------------------------------------")
		}
		print_records(records)
	}

	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println("Final records")
	print_records(records)

	fmt.Printf("Best value for capacity %d is %d", capacity, records[count][capacity])
}

func print_records(records [][]int) {
	// Determine the number of columns
	capacity := len(records[0])

	// Construct and print the header row
	var header string
	for i := 0; i < capacity; i++ {
		header += fmt.Sprintf(" C%-4d |", i)
	}
	fmt.Println(strings.TrimRight(header, "|"))

	// Print a separator line
	separator := strings.Repeat("------", capacity)
	fmt.Println(separator)

	// Print each record row
	for _, record := range records {
		var row string
		for _, value := range record {
			row += fmt.Sprintf(" %-6d|", value)
		}
		fmt.Println(strings.TrimRight(row, "|"))
	}
}
