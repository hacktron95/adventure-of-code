package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic("Couldn't open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	statements := getstatements(scanner.Text())

	for i := 0; i < len(statements); i += 4 {
		switch statements[i] {
		case 99:
			break
		case 1:
			makeAdd(i, statements)
		case 2:
			makeMultiply(i, statements)
		}
	}
	fmt.Println(statements[0])
}

func getstatements(s string) []int {
	a := strings.Split(s, ",")
	b := make([]int, len(a))
	for i, v := range a {
		b[i], _ = strconv.Atoi(v) //we already know the input style
	}
	return b
}

func makeAdd(i int, s []int) {
	index, v1, v2 := s[i+3], s[i+1], s[i+2]
	s[index] = s[v1] + s[v2]
}
func makeMultiply(i int, s []int) {
	index, v1, v2 := s[i+3], s[i+1], s[i+2]
	s[index] = s[v1] * s[v2]
}
