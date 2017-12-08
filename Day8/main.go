package main

import (
	"os"
	"fmt"
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

	var vars = make(map[string]int)
	var lines []string
	m := 0

	for r.Scan() {
		path := r.Text()
		lines = append(lines, path)
		split := strings.Split(path, " ")
		vars[split[0]] = 0
	}
	for r := range lines{
		path := lines[r]
		line := strings.Split(path, " ")

		if line[5] == ">"{
			nr, _ := strconv.Atoi(line[6])
			if vars[line[4]] <= nr{
				continue
			}
		}
		if line[5] == "<"{
			nr, _ := strconv.Atoi(line[6])
			if vars[line[4]] >= nr{
				continue
			}
		}
		if line[5] == "<="{
			nr, _ := strconv.Atoi(line[6])
			if vars[line[4]] > nr{
				continue
			}
		}
		if line[5] == ">="{
			nr, _ := strconv.Atoi(line[6])
			if vars[line[4]] < nr{
				continue
			}
		}
		if line[5] == "=="{
			nr, _ := strconv.Atoi(line[6])
			if vars[line[4]] != nr{
				continue
			}
		}
		if line[5] == "!="{
			nr, _ := strconv.Atoi(line[6])
			if vars[line[4]] == nr{
				continue
			}
		}

		if line[1] == "inc" {
			nr, _ := strconv.Atoi(line[2])
			vars[line[0]] += nr
		}
		if line[1] == "dec" {
			nr, _ := strconv.Atoi(line[2])
			vars[line[0]] -= nr
		}
		if vars[line[0]] > m {
			m = vars[line[0]]
		}
	}
	n := 0
	for r := range vars{
		if vars[r] > n{
			n = vars[r]
		}
	}
	fmt.Println("Part 1: ", n)
	fmt.Println("Part 2: ", m)
}
