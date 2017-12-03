package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	1	+8 +8*1
	9	+16 +8*2
	25	+24 +8*3 etc..
	49
	*/
	input := 312051

	//PART 1

	current := 1
	previous := 0
	nrsquares := 0

	for i := 1; current <= input; i++{

		current = int(math.Pow(float64(i), 2))
		previous = int(math.Pow(float64(i-2), 2))
		nrsquares = i/2
	//	fmt.Println(i/2) //figure how many squares
	//	fmt.Println(current)

		i++
	}
	fmt.Println(previous)
	fmt.Println(current)
	where := current-input
	fmt.Println(where)
	shit := (current-previous)/4 //find length of squares
	shit2 := (shit/2)-where		//find the halfway of the side and how far you are from it

	fmt.Println()
	fmt.Println(math.Abs(float64(shit2))+float64(nrsquares)) //add the length from middle and length from middle of side
	//surprising enough this was equal to the current(first square after input) minus the input

	//PART 2 //doesnt work.. found answer on OEIS.ORG/A141481
	curx := 5
	cury := 5
	sequence := make([][]int, 10)
	for r := range sequence{
		sequence[r] = make([]int, 10)
	}
	for i := 0; i < 10; i++{
		for j:= 0; j < 10; j++{
			sequence[i][j] = 0
		}
	}

	sequence[curx][cury] = 1
	curx++
	for i := 0; i < 10; i++{
		sum := 0
		sum += sequence[curx+1][cury]
		sum += sequence[curx-1][cury]
		sum += sequence[curx][cury-1]
		sum += sequence[curx][cury+1]
		sum += sequence[curx+1][cury+1]
		sum += sequence[curx-1][cury-1]
		sum += sequence[curx-1][cury+1]
		sum += sequence[curx+1][cury-1]

		sequence[curx][cury] = sum
		curx++
	}

	fmt.Println(sequence)
}
