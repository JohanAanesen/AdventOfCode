package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main(){
	b, err := os.Open("input.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	var chars []string

	for r.Scan() {
		path := r.Text()

		chars = strings.Split(path, "")
	}

	total := 0
	garbage := false
	layer := 0
	cancelled := 0

	for r := 0; r < len(chars); r++{

		if chars[r] == "!"{
			r++ //skip next char
			continue
		}
		if garbage && chars[r] == ">"{
			garbage = false
			continue
		}
		if !garbage {
			if chars[r] == "<" {
				garbage = true
				continue
			}
			if chars[r] == "{" {
				layer++
				total += layer
				continue
			}
			if chars[r] == "}" {

				layer--
				continue
			}
		}
		if garbage{
			cancelled++
		}
	}

	//part 1
	fmt.Println("Part 1: ", total)
	//part 2
	fmt.Println("Part 2: ", cancelled)
}
