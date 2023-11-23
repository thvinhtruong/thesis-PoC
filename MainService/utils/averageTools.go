package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CalculateAverage(filename string) float64 {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum float64
	var count int

	// Read each line from the file
	for scanner.Scan() {
		// Convert the line to a number
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}

		// Accumulate the sum and count
		sum += num
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	if count > 0 {
		// Calculate the average
		average := sum / float64(count)
		// convert to megabytes
		average = average / 1024 / 1024
		return average
	} else {
		fmt.Println("No numbers found in the file.")
		return 0
	}
}
