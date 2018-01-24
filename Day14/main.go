package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse2(numbers []int)[]int {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func knot_hash(data []byte)string{

	var lengths []byte
	var intLengths []int

	lengths = append(data,17, 31, 73, 47, 23)
	intLengths = make([]int,len(lengths))
	for i,element := range lengths{
			intLengths[i] = int(element)
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

	return hexString
}

func main() {

	input := "ugkiagan"
	//inpByt := []byte(input)
	var dataMap map[int]string
	dataMap = make(map[int]string)

	for i := 0; i < 128; i++{
		inputNr := input+"-"+strconv.Itoa(i)
		inputByte := []byte(inputNr)
		hashString := knot_hash(inputByte)

		var test []string
		test = strings.Split(hashString, "")
		var binString string
		for r := range test{
			nr, _ := strconv.ParseInt(test[r], 16, 64)

			binayay := strconv.FormatInt(nr, 2)
			binString += fmt.Sprintf("%04s", binayay)

		}
		dataMap[i] = binString
	}

	count := 0
	for r := range dataMap{
		count += countOne(dataMap[r])
	}

	fmt.Println("Count: ", count)
}

func countOne(s string)int{
	count := 0
	me := strings.Split(s, "")

	for r := range me{
		if me[r] == "1"{count++}
	}
	return count
}