package main

import (
	"os"
	"fmt"
	"bufio"
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

	var arr []int
	var arr2 []int

	for r.Scan() {
		path := r.Text()
		nr, _ := strconv.Atoi(path)
		arr = append(arr, nr)
		arr2 = append(arr2, nr)
	}

	//PART 1

	now := 0
	steps := 0
	out := len(arr)

	for now < out{
		arr[now] += 1
		now += arr[now] - 1

		steps++
	}

	fmt.Println(steps)

	//PART 2

	now = 0
	steps = 0
	out = len(arr2)

	for now < out{
		was := arr2[now]
		if was < 3{
			arr2[now] += 1
			now += arr2[now] - 1
		}else if was >= 3{
			arr2[now] -= 1
			now += arr2[now] + 1
		}
		steps++
	}

	fmt.Println(steps)

}
