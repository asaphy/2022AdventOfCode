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
	shapeScores := map[string]map[string]int{}
	aMapping := map[string]int{}
	bMapping := map[string]int{}
	cMapping := map[string]int{}

	aMapping["X"] = 4 // draw(3) + rock(1)
	aMapping["Y"] = 8 // win(6) + paper(2)
	aMapping["Z"] = 3 // loss(0) + scissors(3)

	bMapping["X"] = 1 // loss(0) + rock(1)
	bMapping["Y"] = 5 // draw(3) + paper(2)
	bMapping["Z"] = 9 // win(6) + scissors(3)

	cMapping["X"] = 7 // win(6) + rock(1)
	cMapping["Y"] = 2 // loss(0) + paper(2)
	cMapping["Z"] = 6 // draw(3) + scissors(3)

	shapeScores["A"] = aMapping 
  shapeScores["B"] = bMapping
	shapeScores["C"] = cMapping

	scores := 0

	for scanner.Scan() {
		letters := strings.Fields(scanner.Text())
		scores += shapeScores[letters[0]][letters[1]] 
	}
	fmt.Println("Answer:", scores)
}
