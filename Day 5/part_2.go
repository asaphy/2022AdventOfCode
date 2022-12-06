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
	"regexp"
	"strings"
	"strconv"
)

var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(t T) {
	s.items = append(s.items, t)
}

func (s *Stack[T]) Pop() T {
	n := len(s.items)
	t := s.items[n-1]
	var zero T
	s.items[n-1] = zero
	s.items = s.items[:n-1]
	return t
}

func (s *Stack[T]) Reverse() {
	for i, j := 0, len(s.items)-1; i < j; i, j = i+1, j-1 {
    s.items[i], s.items[j] = s.items[j], s.items[i]
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(readFile)
	stackSlice := map[int]*Stack[string]{}
	for i := 1; i < 10; i++ {
		stack := new(Stack[string])
	  stackSlice[i] = stack	
	}
	lineCounter := 0
	for scanner.Scan() {
		if lineCounter < 9 {
			characters := strings.Split(scanner.Text(), "")
			for i:=0; i<len(characters); i+=4 {
				if IsLetter(characters[i+1]) {
				  stackSlice[(i+1)/4 + 1].Push(characters[i+1])
				}
			}
		} else if lineCounter == 9 {
			for i := 1; i < 10; i++ {
				stackSlice[i].Reverse()	
			} 
		} else if lineCounter >= 10 {
			moveWords := strings.Split(scanner.Text(), " ")
			numItemsToMove, _ := strconv.Atoi(moveWords[1])
			fromStack, _ := strconv.Atoi(moveWords[3])
			toStack, _ := strconv.Atoi(moveWords[5])
			Move(stackSlice, numItemsToMove, fromStack, toStack)
		}
	  lineCounter ++
	}
	answer := ""
	for i := 1; i < 10; i++ {
		answer += stackSlice[i].Pop() 
	}
	fmt.Println("Answer:", answer)
}

func Move(stackSlice map[int]*Stack[string], numItemsToMove int, fromStack int, toStack int) {
	itemsToMove := new(Stack[string])
	for i:=0; i<numItemsToMove; i++{
	  itemToMove := stackSlice[fromStack].Pop()
		itemsToMove.Push(itemToMove)
	}
	for !itemsToMove.IsEmpty() {
	  itemToMove := itemsToMove.Pop()
		stackSlice[toStack].Push(itemToMove)
	}
}
