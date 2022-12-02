//run with:
//go run part_1.go data.txt
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(readFile)
	mostCalories := 0
	caloriesOfCurrentElf := 0

	for scanner.Scan() {
		if(scanner.Text() == "") {
			if caloriesOfCurrentElf > mostCalories {
				mostCalories = caloriesOfCurrentElf
			}
			caloriesOfCurrentElf = 0
		} else {
			intVar, err := strconv.Atoi(scanner.Text())
			caloriesOfCurrentElf += intVar
			fmt.Println(intVar, err)		
		}
	}
	fmt.Println(mostCalories)
	readFile.Close()
}
