package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

var nodemap = make(map[string][]string)
var values = make(map[string]int)

//recursively build the tree until a divergent value appears.
func child_values(child string)[]int{
	childs := nodemap[child]
	var supports []int
	for r := range childs{
		value := 0
		if exists(childs[r]){
			child_support := child_values(childs[r])
			value = sumSlice(child_support) + values[childs[r]]
		}else{
			value = values[childs[r]]
		}
		supports = append(supports, value)
	}
	if isDivergent(supports){
		fmt.Println("Divergent child found on node: ", child, " with children: ", nodemap[child], " weighing: ", supports)
		//I can see that the wrong node is in index 0, and then retrieve it's original value. Then +- the deviation to get the answer
		fmt.Println("fucker: ", nodemap[childs[0]], " val: ", values[childs[0]]) //in this case it is minus 8, so answer is 2310, not 2318
	}

	return supports
}

func isDivergent(n1 []int)bool{
	for r := range n1{
		if n1[r] != n1[0]{
			return true
		}
	}
	return false
}

func exists(s string)bool{
	for r := range nodemap{
		if r == s{
			return true
		}
	}
	return false
}

func sumSlice(n1 []int)int{
	total := 0
	for r := range n1{
		total += n1[r]
	}
	return total
}

func main() {

	b, err := os.Open("input.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	var nodes []string
	var supported []string





	for r.Scan() {
		node := r.Text()

		if strings.Contains(node, "->"){

			node = strings.Replace(node, ",", "", -1)
			split := strings.Split(node, " ")
			nodes = append(nodes, split[0])
			val := strings.Trim(split[1], "()")
			val2, _ := strconv.Atoi(val)
			values[split[0]] = val2

			for i := 3; i < len(split); i++{
				supported = append(supported, split[i])
				nodemap[split[0]] = append(nodemap[split[0]], split[i])
			}
		}else{
			split := strings.Split(node, " ")
			nodes = append(nodes, split[0])

			val := strings.Trim(split[1], "()")
			val2, _ := strconv.Atoi(val)
			values[split[0]] = val2
		}
	}

	nodecpy := nodes
	answer := ""
	for r := range nodes{
		for e := range supported{
			if nodes[r] == supported[e]{
				nodecpy[r] = ""
			}
		}
		if nodecpy[r] != ""{
			answer = nodecpy[r]
		}
	}

	fmt.Println(answer)

	//PART 2

	val := child_values(answer)

	fmt.Println(val)

}
