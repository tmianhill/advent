package main

import (
	"advent/utils"
	"fmt"
	"slices"
	"strings"
)

func main() {
	sampleInput := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	input := sampleInput
	input = realInput

	lines := strings.Split(input, "\n")

	rules := []rule{}

	correctTotal := 0
	wrongTotal := 0
	for _,l := range lines {
		ruleVals := utils.SplitAndParseInts(l, "|")
		if len(ruleVals) == 2 {
			rules = append(rules, rule{before: ruleVals[0], after:ruleVals[1]})
			continue
		}

		pages := utils.SplitAndParseInts(l, ",")
		if len(pages) > 1 {
			pageIndex := indexPages(pages)

			orderCorrect := true
			for _,r := range rules {
				beforeIndex,beforeOK :=pageIndex[r.before]
				afterIndex,afterOK := pageIndex[r.after]
				if beforeOK && afterOK && beforeIndex > afterIndex {
					orderCorrect = false
					break
				}
			}

			if orderCorrect {
				middleVal := pages[(len(pages)-1)/2]
				correctTotal += middleVal
				fmt.Println(l, "correct", middleVal)
			} else {
				sorted := sortPages(pages, rules)
				middleVal := pages[(len(pages)-1)/2]
				wrongTotal += middleVal
				fmt.Println(l, "wrong, corrected to", sorted, middleVal)
			}	
		}
	}

	fmt.Println("correct", correctTotal)
	fmt.Println("wrong", wrongTotal)
}

type rule struct { before int; after int }

func indexPages(pages []int) map[int]int {
	indexes := make(map[int]int, len(pages))
	for i,p:=range pages { indexes[p]=i }
	return indexes
}

func sortPages(pages []int, rules []rule) []int {
	slices.SortFunc(pages, func(a,b int) int {
		for _,r := range rules {
			if r.after == a && r.before == b {
				return 1
			}
			if r.before == a && r.after == b {
				return -1
			}
		}
		return 0
	} )
	return pages
}

