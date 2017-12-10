package main

import (
	"fmt"
	"os"
	"bufio"
)

func reverse2(numbers []int)[]int {
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

	var lengths []byte
	var intLengths []int
	for r.Scan() {
		path := r.Bytes()
		lengths = append(path,17, 31, 73, 47, 23)
		intLengths = make([]int,len(lengths))
		for i,element := range lengths{
			intLengths[i] = int(element)
		}
	}

	var list []int
	for i := 0; i < 256; i++{
		list = append(list, i)
	}

	current := 0
	skip := 0

	input := intLengths

	for rounds := 0; rounds < 64; rounds++ {
		for r := range input {
			var curr []int
			length := input[r]

			for i := 0; i < length; i++ {
				index := (current+ i) % len(list)
				curr = append(curr, list[index])
			}

			curr = reverse2(curr)
			index2 := 0
			for i := 0; i < length; i++ {
				index := (current+ i) % len(list)
				list[index] = curr[index2]
				index2++
			}
			current += length + skip
			skip++
		}
	}

	denseHash := make([]int,len(list)/16)
	for i:=0; i < len(list);i+=16{
		for j:=0; j <16; j++{
			denseHash[i/16] ^= list[i+j]
		}
	}

	hexString :=""
	for _,element := range denseHash{
		hexString += fmt.Sprintf("%.02x", element)
	}

	fmt.Println("Part 2: ", hexString)
}
