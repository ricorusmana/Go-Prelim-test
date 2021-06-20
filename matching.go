package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var total int

	totalmatching := make(map[int]int)

	total_socks, coll_color_socks := inputData()

	if !inputValidation(total_socks, coll_color_socks) {
		fmt.Printf("Incorrect input condition")
	} else {
		for _, value := range coll_color_socks {
			if checkMatchingColor(value, coll_color_socks) > 0 {
				if totalmatching[value] == checkMatchingColor(value, coll_color_socks) {
					continue
				} else {
					totalmatching[value] = checkMatchingColor(value, coll_color_socks)
					total += checkMatchingColor(value, coll_color_socks)
				}
			}
		}

		fmt.Printf("Total Matching Pairs of Socks: %d", +total)
	}
}

func inputValidation(total_socks int, coll_color_socks []int) bool {
	if total_socks < len(coll_color_socks) {
		return false
	} else if ((total_socks < 0) || (total_socks > 100)) && ((len(coll_color_socks) < 0) || (len(coll_color_socks) > 100)) {
		return false
	} else {
		for _, value := range coll_color_socks {
			if (value < 0) || (value > 100) {
				return false
			}
		}
		return true
	}
}

func checkMatchingColor(value int, collcolor []int) int {

	countColor := 0

	for index := range collcolor {
		if collcolor[index] == value {
			countColor += 1
		}
	}
	return countColor / 2
}

func inputData() (int, []int) {

	scanner := bufio.NewScanner(os.Stdin)
	var coll_color_socks []int

	fmt.Print("Enter total socks: ")
	scanner.Scan()
	input_total_socks := scanner.Text()

	fmt.Print("Enter color socks: ")
	scanner.Scan()
	input_color_socks := scanner.Text() //1 2 3 1 2 3 1

	total_socks, err := strconv.Atoi(input_total_socks)
	if err != nil {
		fmt.Println("Please Input in Correct Format")
	}

	temp_coll_color_socks := strings.Split(input_color_socks, " ")

	for _, value := range temp_coll_color_socks {
		j, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Please Input in Correct Format")
		}
		coll_color_socks = append(coll_color_socks, j)
	}
	return total_socks, coll_color_socks
}
