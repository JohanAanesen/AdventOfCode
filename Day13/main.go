package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

/*
Solved part1 using the map, part 2 using 2 separate slices
*/


func main() {
	b, err := os.Open("input.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	var data = make(map[int]int)
	var keys []int
	var vals []int

	for r.Scan() {
		path := r.Text()

		input := strings.Split(path, ": ")

		nr, _ := strconv.Atoi(input[0])
		nr2, _ := strconv.Atoi(input[1])

		keys = append(keys, nr) //keep a sorted keys slice to iterate the map
		vals = append(vals, nr2)
		data[nr] = nr2
	}

	severity := 0

	for i := 0; i < keys[len(keys)-1]; i++{
		if data[i] == 0{
			continue
		}
		if i % ((data[i] - 1) * 2) == 0{
			severity += data[i] * i
		}
	}

	fmt.Println("Part 1: ", severity)

	delay := 0
	searching := true

	for searching == true{
		searching = false
		for i := 0; i < len(keys); i++{
			if (keys[i] + delay) % ((vals[i] * 2) - 2) == 0{
				searching = true
				delay++
				break
			}
		}
	}


	fmt.Println("Part 2: ", delay)
}
