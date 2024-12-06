package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	sampleInput := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	input := sampleInput
	input = realInput

	lines := strings.Split(input, "\n")
	copies := make([]int, len(lines))
	total := 0
	for i, line := range lines {
		copies[i]++
		parts := strings.Split(line, ":")
		parts = strings.Split(parts[1], "|")
		winNums := strings.Split(strings.TrimSpace(parts[0]), " ")
		myNums := strings.Split(strings.TrimSpace(parts[1]), " ")
		matchCount := 0
		for _, n := range winNums {
			if n != "" && slices.Contains(myNums, n) {
				matchCount++
			}
		}
		for j:=0;j<matchCount;j++ {
			copies[i+j+1]+=copies[i]
		}
	fmt.Println("after",i,":",copies)
	}

	for k,v := range copies {
		fmt.Println(k,v)
		total += v
	}

	fmt.Println(total)
}
