package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func reverse(numbers []int)[]int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func main() {
	b, err := os.Open("input.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	var input []int

	for r.Scan() {
		path := r.Text()
		chars := strings.Split(path, ",")
		for r := range chars{
			nr, _ := strconv.Atoi(chars[r])
			input = append(input, nr)
		}
	}

	var list []int
	for i := 0; i < 256; i++{
		list = append(list, i)
	}

	current := 0
	skip := 0


	for r := range input{
		var curr []int
		length := input[r]

		for i := 0; i < length; i++{
			index := current+i
			if index >= len(list){
				index -= len(list)
			}
			curr = append(curr, list[index])
		}

		curr = reverse(curr)
		index2 := 0
		for i := 0; i < length; i++{
			index := current+i
			if index >= len(list){
				index -= len(list)
			}
			list[index] = curr[index2]
			index2++
		}
		current += length + skip
		if current >= len(list){
			current -= len(list)
		}
		skip++
	}


	fmt.Println("Part 1: ", list[0]*list[1])

}

