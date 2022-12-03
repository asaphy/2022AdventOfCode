/*
to run:
	go run part_1.go data.txt

time complexity:

space complexity:

notes:

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(readFile)

	scoreMap := map[string]int{}
	i := 1
	for ch := 'a'; ch <= 'z'; ch++ {
		scoreMap[string(ch)] = i
		i++
	}
	i = 27
	for ch:= 'A'; ch <= 'Z'; ch++ {
		scoreMap[string(ch)] = i
		i++
	}

	prioritySum := 0
	for scanner.Scan() {
		currentRucksack := map[string]bool{}
		items := strings.Split(scanner.Text(), "")
		for i := 0; i < len(items)/2 ; i++ {
			currentRucksack[items[i]] = true
		}
		for i := len(items)/2; i < len(items); i++ {
			if currentRucksack[items[i]] == true {
				prioritySum += scoreMap[items[i]]
				break
			}
		}
	}
	fmt.Println("Answer:", prioritySum)
}
