package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"math"
)

func main() {
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

		chars = strings.Split(path, ",")
	}

	x := 0
	y := 0
	z := 0

	var dist []float64

	for r := range chars{
		if chars[r] == "n"{
			y++
			z--
		}
		if chars[r] == "ne"{
			x++
			z--
		}
		if chars[r] == "nw"{
			x--
			y++
		}
		if chars[r] == "se"{
			y--
			x++
		}
		if chars[r] == "s"{
			y--
			z++
		}
		if chars[r] == "sw"{
			x--
			z++
		}
		dist = append(dist, (math.Abs(float64(x))+ math.Abs(float64(y))+ math.Abs(float64(z)))/2)
	}

	part1 := (math.Abs(float64(x))+ math.Abs(float64(y))+ math.Abs(float64(z)))/2


	part2 := 0.0
	for r := range dist{
		if dist[r] > part2{
			part2 = dist[r]
		}
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
