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
	charactersToBeProcessed := 0
	numberOfDifferentCharactersNeeded := 4
	for scanner.Scan() {
		characters := strings.Split(scanner.Text(), "")
		for i:=0; i<len(characters); i++ {
			charactersToBeProcessed++
			if i < numberOfDifferentCharactersNeeded-1 {
				continue
			} else if characters[i] == characters[i-1] {
				continue
			} else if characters[i] == characters[i-2] {
				continue
			} else if characters[i] == characters[i-3] {
				continue
			} else if characters[i-1] == characters[i-2] {
				continue
			} else if characters[i-1] == characters[i-3] {
				continue
			} else if characters[i-2] == characters[i-3] {
				continue
			} else {
				break
			}
		}
	}
	fmt.Println("Answer:", charactersToBeProcessed)
}
