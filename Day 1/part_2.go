//run with:
//go run part_2.go data.txt
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(readFile)
	topThreeCalories := make([]int, 3)
	caloriesOfCurrentElf := 0

	for scanner.Scan() {
		if(scanner.Text() == "") {
			if isCaloriesMoreThanCurrentTopThree(topThreeCalories, caloriesOfCurrentElf) {
				sort.Ints(topThreeCalories)
				topThreeCalories[0] = caloriesOfCurrentElf
			}
			caloriesOfCurrentElf = 0
		} else {
			intVar, err := strconv.Atoi(scanner.Text())
			caloriesOfCurrentElf += intVar
			fmt.Println(intVar, err)		
		}
	}
	fmt.Println(topThreeCalories)
	readFile.Close()
}

func isCaloriesMoreThanCurrentTopThree(topThreeSlice []int, caloriesOfCurrentElf int) bool {
	for i := 0; i < len(topThreeSlice) ; i++ {
		if caloriesOfCurrentElf > topThreeSlice[i] {
			return true
		}
	}
	return false
}
