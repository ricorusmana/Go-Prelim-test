package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Number of steps: ")
	scanner.Scan()
	input_number := scanner.Text()

	for index, value := range input_number {
		current_value, err := strconv.Atoi(string(value))
		if err != nil {
			fmt.Println("Please Input in Correct Format")
		}
		current_value = int(float64(current_value) * math.Pow(10, float64(len(input_number)-index-1)))
		fmt.Println(current_value)
	}
}
