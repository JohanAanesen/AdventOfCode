package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func recurs(nr int){
	visited = append(visited, nr)

	for _, val := range values[nr]{
		if !isVisited(val){
			recurs(val)
			magic++
		}
	}
}

func isVisited(nr int)bool{
	for r := range visited{
		if nr == visited[r]{
			return true
		}
	}
	return false
}

var values = make(map[int][]int)
var visited []int
var magic int

func main() {
	b, err := os.Open("input.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)



	for r.Scan() {
		path := r.Text()

		temp := strings.Split(path, " <-> ")
		key, _ := strconv.Atoi(temp[0])


		tunnels := strings.Split(temp[1], ", ")
		for r := range tunnels{
			nr, _ := strconv.Atoi(tunnels[r])
			values[key] = append(values[key], nr)
		}
	}

	magic++
	recurs(0)
	part1 := magic

	groups := 1
	for r := range values{
		if !isVisited(r){ //find unvisited node and run the recursive thing on it, increment groups
			groups++
			recurs(r)
		}
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", groups)
}
