package main

import "fmt"

var carousel []int

func main() {

	jumps := 363
	index := 0


	//init
	carousel = append(carousel, 0)

	for i := 1; i < 2018; i++{
		index = (index + jumps) % len(carousel)
		index++
		insertMapShift(index, i)
	}

	for r := range carousel{
		if carousel[r] == 2017{
			fmt.Println("Part 1: ", carousel[r+1])
		}
	}

	//PART 2
	part2 := 0
	for i := 1; i < 50000000+1; i++{
		index = (index + jumps) % i+1
		if index == 1{
			part2 = i
		}
	}

	fmt.Println("Part 2: ", part2)

}


func insertMapShift(index int, value int){
	var buffer []int

    buffer = append(buffer, value) //putting value first in buffer

	for i := index; i < len(carousel); i++{
		buffer = append(buffer, carousel[i])
	}

	carousel = append(carousel, 0) //expand slice with temp value

	for i := 0; i < len(buffer); i++{
		carousel[index+i] = buffer[i]
	}
}