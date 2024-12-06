package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleInput := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	input := sampleInput
	//input = realInput

	lines := strings.Split(input, "\n")

	chars := make([][]string, len(lines))
	for i, l := range lines {
		chars[i] = strings.Split(l, "")
	}

	count := 0
	for i := 1; i < len(chars)-1; i++ {
		for j := 1; j < len(chars[0])-1; j++ {
			if chars[i][j] == "A" {
				pos := false
				//diag pos
				if chars[i-1][j-1] == "M" && chars[i+1][j+1] == "S" {
					pos = true
				}
				if chars[i-1][j-1] == "S" && chars[i+1][j+1] == "M" {
					pos = true
				}
				//diag neg
				neg := false
				if chars[i-1][j+1] == "M" && chars[i+1][j-1] == "S" {
					neg = true
				}
				if chars[i-1][j+1] == "S" && chars[i+1][j-1] == "M" {
					neg = true
				}
				if pos && neg {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
