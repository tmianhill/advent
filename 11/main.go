package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

func main() {
	input := sampleInput
	input = realInput

	stones := utils.SplitAndParseInts(input, " ")
	fmt.Println(stones)

	stoneCounts := map[int]int{}
	for _, s := range stones {
		stoneCounts[s]++
	}

	for i := range 75 {
		newStoneCounts := map[int]int{}
		total := 0
		for s,c := range stoneCounts {
			for _, ns := range workStone(s) {
				newStoneCounts[ns] += c
				total += c
			}
		}
		stoneCounts = newStoneCounts
		fmt.Println(i+1, "blinks:", total, "stones")
	}




	_ = input
}

func workStone(s int) []int {
	if s == 0 {
		return []int{1}
	}
	ss := strconv.Itoa(s)
	if len(ss) % 2 == 0 {
		n1,_ := strconv.Atoi(ss[:len(ss)/2])
		n2,_ := strconv.Atoi(ss[len(ss)/2:])
		return []int{n1, n2}
	}

	return []int{s*2024}
}

