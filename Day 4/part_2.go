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
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(readFile)

	assignmentsFullyContainedSum := 0
	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")

		assignmentOne := strings.Split(assignments[0], "-") 
		assignmentTwo := strings.Split(assignments[1], "-")
		assignmentOneStart, _ := strconv.Atoi(assignmentOne[0])
		assignmentOneEnd, _ := strconv.Atoi(assignmentOne[1])
		assignmentTwoStart, _ := strconv.Atoi(assignmentTwo[0])
		assignmentTwoEnd, _ := strconv.Atoi(assignmentTwo[1])
		
		if (assignmentOneStart >= assignmentTwoStart && assignmentOneEnd >= assignmentTwoEnd) || (assignmentTwoStart <= assignmentOneStart && assignmentTwoEnd >= assignmentOneEnd) {
			assignmentsFullyContainedSum += 1
		}
	}
	fmt.Println("Answer:", assignmentsFullyContainedSum)
}
