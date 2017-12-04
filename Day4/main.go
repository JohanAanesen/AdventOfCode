package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"sort"
)

func main() {
	b, err := os.Open("input.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	var arr []string

	for r.Scan() {
		path := r.Text() // 0x0A separator = newline

		arr = append(arr, path)
	}

	//PART 1

	invalid := 0
	for i := 0; i < len(arr); i++{
		equal := false
		check := strings.Split(arr[i], " ")
		sort.Strings(check)

		for j := 1; j < len(check); j++{
			if !equal && check[j] == check[j-1]{
				equal = true
				invalid++
				break
			}
		}
	}

	fmt.Println(len(arr)-invalid)

	//PART 2

	invalid = 0
	for i := 0; i < len(arr); i++{
		equal := false
		check := strings.Split(arr[i], " ")

		//splitting each word, sorting by character and joins them back together
		for s := 0; s < len(check); s++{
			s1 := strings.Split(check[s], "")
			sort.Strings(s1)
			check[s] = strings.Join(s1, "")
		}


		sort.Strings(check)

		for j := 1; j < len(check); j++{
			if !equal && check[j] == check[j-1]{
				equal = true
				invalid++
				break
			}
		}
	}

	fmt.Println(len(arr)-invalid)



}
