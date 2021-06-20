package main

import "fmt"

func main() {
	counter_lap := 1
	counter_switch := 1
	total_switch := 100
	switch_on := make(map[int]int)
	for counter_lap = 1; counter_lap <= total_switch; counter_lap++ {
		for counter_switch = 1; counter_switch <= total_switch; counter_switch++ {
			if counter_switch%counter_lap == 0 {
				if switch_on[counter_switch] == 1 {
					delete(switch_on, counter_switch)
				} else {
					switch_on[counter_switch] = 1
				}
			}
		}
	}

	fmt.Printf("\nTotal switches on after 100 lap: %d\n", len(switch_on))
}
