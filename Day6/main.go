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

	var arr []int
	var arr2 []string

	for r.Scan() {
		path := r.Text()
		split := strings.Split(path, "\t")
		for i := range split{
			nr, _ := strconv.Atoi(split[i])
			arr = append(arr, nr)
		}
	}

	a1 := fmt.Sprint(arr)
	arr2 = append(arr2, a1)

	cycles := 0
	repeat := 0

	for{
		current := 0
		m := arr[0]
		for i, e := range arr {
			if e > m {
				m = e
				current = i
			}
		}

		iterate := current + 1
		kjor := arr[current]
		arr[current] = 0
		for i := 0; i < kjor; i++{
			if iterate == len(arr){
				iterate = iterate - len(arr) //or just 0
			}

			arr[iterate]++

			iterate++
		}

		cycles++

		data := fmt.Sprint(arr)

		exists := false
		for r := range arr2{
			if arr2[r] == data{
				exists = true
				repeat = r
			}
		}

		if exists{
			break
		}
		arr2 = append(arr2, data)
	}



	//PART 1
	fmt.Println("Part 1: ",cycles)
	//PART 2
	fmt.Println("Part 2: ",len(arr2)-repeat)
}
