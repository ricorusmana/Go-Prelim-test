package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var current_step string
	var number_valley int

	total_steps, input_path := inputData()

	if !inputValidation(total_steps, input_path) {
		fmt.Printf("Incorrect input condition")
	} else {
		current_position := 0

		for index, value := range input_path {

			current_step = string(value)
			if current_position <= 0 && index+1 < len(input_path) {
				if (current_step == string(input_path[index+1])) && (current_step == "D") {
					number_valley += 1
				}
			}
			if current_step == "U" {
				current_position += 1
			} else {
				current_position -= 1
			}

		}
		fmt.Printf("number of valley: %d", number_valley)
	}
}

func inputValidation(total_steps int, input_path string) bool {
	if total_steps < len(input_path) {
		return false
	} else if total_steps < 2 || total_steps > 1000000 {
		return false
	} else {
		for _, value := range input_path {
			if string(value) != "U" && string(value) != "D" {
				return false
			}
		}
		return true
	}

}

func inputData() (int, string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Number of steps: ")
	scanner.Scan()
	input_total_steps := scanner.Text()

	fmt.Print("Path: ")
	scanner.Scan()
	input_path := scanner.Text()

	total_steps, err := strconv.Atoi(input_total_steps)
	if err != nil {
		fmt.Println("Please Input in Correct Format")
	}

	return total_steps, input_path
}
