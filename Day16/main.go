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

	var commands []string

	for r.Scan() {
		path := r.Text()

		input := strings.Split(path, ",")
		for r := range input{
			commands = append(commands, input[r])
		}
	}

	chars := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p"}
	var ans = make(map[int]string)
	turtall := 0

	for j := 0; j < 1000000000; j++ { //bruteforce, although we wont be using this many operations
		for i := 0; i < len(commands); i++ {
			//temp := strings.Split(commands[i], "")
			if strings.HasPrefix(commands[i], "x") {
				temp := strings.TrimPrefix(commands[i], "x")
				char := strings.Split(temp, "/")

				var tempx, tempy int
				var buffer string

				tempx, _ = strconv.Atoi(char[0])
				tempy, _ = strconv.Atoi(char[1])
				buffer = chars[tempx]
				chars[tempx] = chars[tempy]
				chars[tempy] = buffer
			}
			if strings.HasPrefix(commands[i], "s") {
				temp := strings.TrimPrefix(commands[i], "s")
				nr, _ := strconv.Atoi(temp)

				for count := 1; count <= nr; count++ {
					tmp := chars[len(chars)-1]
					for n := len(chars) - 2; n >= 0; n-- {
						chars[n+1] = chars[n]
					}
					chars[0] = tmp
				}
			}
			if strings.HasPrefix(commands[i], "p") {
				temp := strings.TrimPrefix(commands[i], "p")
				char := strings.Split(temp, "/")

				var a, b int
				var buffer string

				for r := range chars {
					if char[0] == chars[r] {
						a = r
					}
					if char[1] == chars[r] {
						b = r
					}
				}

				buffer = chars[a]
				chars[a] = chars[b]
				chars[b] = buffer

			}

		}

		if checkAns(strings.Join(chars, ""), ans){
			break
		}

		ans[turtall] =  strings.Join(chars, "")
		turtall++
	}

	fmt.Println(turtall)
	nr := 1000000000 % turtall
	fmt.Println(nr)
	fmt.Println("Part 1: ", ans[0])
	fmt.Println("Part 2: ", ans[nr-1])
}

func checkAns(s1 string, ans map[int]string)bool{
	for i := 0; i < len(ans)-1; i++{
		if ans[i] == s1{
			return true
		}
	}
	return false
}