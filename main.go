package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	if len(os.Args) > 1 {
		val, err := strconv.ParseFloat(os.Args[1], 32)
		if err != nil {
			fmt.Println("Try with numbers dumbass!")
			os.Exit(1)
		}
		result := brains(float32(val))
		fmt.Printf("Fahrenheit: %.1f\n", result)
		os.Exit(0)
	}

	for {
		fmt.Print("Enter Celsius: ")

		// Read the input as a string
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Parse the string into a float64 (temp variable)
		val, err := strconv.ParseFloat(input, 32)
		if err != nil {
			fmt.Println("Error: Try with numbers, dummy!")
			continue
		}

		// Convert that float64 to float32
		celsius := float32(val)
		f := brains(celsius)

		fmt.Printf("Fahrenheit: %.1f\n", f)
		os.Exit(0) // Exit after one successful conversion
	}
}

func brains(celsius float32) float32 {
	// var answer float32
	answer := (celsius * 9 / 5) + 32
	return answer
}
