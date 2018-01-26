package main

import "fmt"

func main() {

	var a, b, factorA, factorB, div, count uint64
	a = 873
	b = 583

	factorA = 16807
	factorB = 48271

	div = 2147483647

	count = 0

	for i := 0; i < 40000000; i++{
		a = a * factorA % div
		b = b * factorB % div

		if a&0xffff == b&0xffff {
			count++
		}

	}


	fmt.Println("Part1: ", count)
	fmt.Println("Part2: ", partTwo())

}

func partTwo()uint64{
	var a, b, factorA, factorB, count uint64
	a = 873
	b = 583

	factorA = 16807
	factorB = 48271

	count = 0

	for i := 0; i < 5000000; i++{
		a = generate(a, factorA, 4)
		b = generate(b, factorB, 8)
		if a&0xffff == b&0xffff {
			count++
		}
	}

	return count
}

func generate(i uint64, factor uint64, mod uint64)(uint64){
	var div uint64
	div = 2147483647

	for{
		i = i * factor % div
		if i%mod == 0{
			return i
		}
	}
}