package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Function to calculate the median
func calculateMedian(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	sort.Float64s(data)
	n := len(data)
	if n%2 == 1 {
		return data[n/2] // Odd number of elements
	}
	return (data[n/2-1] + data[n/2]) / 2 // Even number of elements
}

// Function to calculate the absolute median deviation (MAD)
func calculateMAD(data []float64, median float64) float64 {
	var deviations []float64
	for _, value := range data {
		deviations = append(deviations, abs(value-median))
	}
	return calculateMedian(deviations) // Median of absolute deviations
}

// Helper function to calculate absolute value
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// Dynamic factor calculation: larger datasets might need smaller factors for MAD

func main() {
	var data []float64
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error reading number:", err)
			continue
		}

		// Add current number to data first
		data = append(data, num)

		// Calculate median and MAD only if we have more than 1 data point
		if len(data) > 1 {
			median := calculateMedian(data)
			mad := calculateMAD(data, median)

			// Use dynamic factor based on data size
	

			// Calculate optimized range
			lowerBound := median - 1.5*mad
			upperBound := median + 1.5*mad
			fmt.Println(lowerBound, upperBound)
		} else {
			fmt.Println("Not enough data to base range on")
		}
	}
}
