package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {

	b, err := os.Open("input.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	var input = make(map[int][]int)
	index := 1
	for r.Scan() {
		path := r.Text()
		split := strings.Split(path, "\t")
		for i := range split{
			nr, _ := strconv.Atoi(split[i])
			input[index] = append(input[index], nr)
		}
		index++
	}



	//PART 1
	total := 0
	for i := 1; i <= 16; i++{
		min := 999999
		max := 0
		for j := 0; j < len(input[i]); j++{
			if input[i][j] < min{ min = input[i][j] }
			if input[i][j] > max{ max = input[i][j] }

		}
		total += max-min
	}

	fmt.Println(total)

	//PART 2
	total = 0

	for i := 1; i <= 16; i++{
		sum := 0

		for j := 0; j < len(input[i]); j++{
			for k := 0; k < len(input[i]); k++{
				nr := input[i][j]
				nr2 := input[i][k]
				if nr > nr2 && j != k{
					if nr%nr2 == 0{
						sum = nr/nr2
					}
				}
			}
		}
		total += sum
	}
	fmt.Println(total)
}