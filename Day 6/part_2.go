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
	charactersToBeProcessed := 0
	numberOfDifferentCharactersNeeded := 14
	for scanner.Scan() {
		characters := strings.Split(scanner.Text(), "")
		for i:=0; i<len(characters); i++ {
			charactersToBeProcessed++
			if i < numberOfDifferentCharactersNeeded-1 {
				continue
			}
			if CheckIfMeetsCharactersNeeded(numberOfDifferentCharactersNeeded, i, characters) {
				break
			}
		}
	}
	fmt.Println("Answer:", charactersToBeProcessed)
}

func CheckIfMeetsCharactersNeeded(numberOfDifferentCharactersNeeded int, i int, characters []string ) bool {
	for j:=i; j>i-numberOfDifferentCharactersNeeded; j--{
		for k:=j-1; k>i-numberOfDifferentCharactersNeeded; k--{
			if characters[j] == characters[k] {
				return false
			}
		}
	}
	return true
}