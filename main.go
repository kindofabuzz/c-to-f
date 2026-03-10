package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Move this outside the loop!

	for {
		fmt.Print("Enter Celsius: ")

		// 1. Read the input as a string
		input, _ := reader.ReadString('\n') // Note the '\n' with a backslash!
		input = strings.TrimSpace(input)

		// 2. Parse the string into a float64 (temp variable)
		val, err := strconv.ParseFloat(input, 32)
		if err != nil {
			fmt.Println("Error: Try with numbers, friend!")
			continue
		}

		// 3. Convert that float64 to float32 for your 'brains' function
		celsius := float32(val)
		f := brains(celsius)

		fmt.Printf("Fahrenheit: %.1f\n", f)
		break // Exit after one successful conversion
	}
}

func brains(celsius float32) (answer float32) {
	// var answer float32
	answer = (celsius * 9 / 5) + 32
	return
}
