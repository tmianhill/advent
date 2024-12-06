package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sampleInput := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	input := sampleInput
	input = realInput

	lines := strings.Split(input, "\n")

	chars := make([][]string, len(lines))
	for i, l := range lines {
		chars[i] = strings.Split(l, "")
	}

	regex := regexp.MustCompile(`\d+`)

	gearParts := map[struct{x int;y int}][]int{}

	sum := 0
	for i,l := range lines {
		matchIndexes := regex.FindAllStringIndex(l, -1)
		matchStrings := regex.FindAllString(l, -1)
		for mi,m := range matchStrings {
			val,_ := strconv.Atoi(m)

			ind := matchIndexes[mi]
			startCol := ind[0]
			if startCol > 0 { startCol--}
			endCol := ind[1]
			if endCol < len(l) { endCol++}
			startRow := i-1
			if startRow < 0 { startRow = 0 }
			endRow := i+1
			if endRow >= len(lines) { endRow-- }

			for i:=startRow;i<=endRow;i++ {
				for j:=startCol;j<endCol;j++ {
					c := chars[i][j]
					if c == "*" {
						coords := struct{x,y int}{x:i,y:j}
						gearParts[coords] = append(gearParts[coords], val)
						fmt.Println(m, i, j)
					}
				}
			}
		}
	}

	for k,v := range gearParts {
		if len(v) == 2 {
			prod := v[0] * v[1]
			sum += prod
			fmt.Println(k.x, k.y, v[0], v[1], prod)
		}
	}
	
	fmt.Println(sum)
}

