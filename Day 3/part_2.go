/*
to run:
	go run part_2.go data.txt

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
	groupRucksack := map[string]int{}
	groupCount := 0

	for scanner.Scan() {
		items := strings.Split(scanner.Text(), "")
		groupCount += 1
		if groupCount == 3 {
			for i := 0; i < len(items) ; i++ {
				if groupRucksack[items[i]] == 2 {
				  prioritySum += scoreMap[items[i]]
					groupRucksack = make(map[string]int) 
					groupCount = 0	
					break
				}

			}
			groupRucksack = make(map[string]int) 
			groupCount = 0
		} else{
			currentRucksack := make(map[string]bool)
			for i := 0; i < len(items) ; i++ {
				value, ok := currentRucksack[items[i]]
        if ok {
          fmt.Println("value: ", value)
        } else {
					groupRucksack[items[i]]	+= 1
					currentRucksack[items[i]] = true 
        }
			}	
		}
	}
	fmt.Println("Answer:", prioritySum)
}
