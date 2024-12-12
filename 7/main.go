package main

import (
	"advent/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := sampleInput
	input = realInput

	opers := []operator{PLUS,TIMES,CONCAT}
	numOpers := len(opers)

	lines := strings.Split(input, "\n")
	total := 0
	for _,l := range lines {
		bits := strings.Split(l, ":")
		target, _ := strconv.Atoi(strings.TrimSpace(bits[0]))
		vals := utils.SplitAndParseInts(strings.TrimSpace(bits[1]), " ")

		max := int(math.Pow(float64(len(opers)), float64(len(vals)-1)))
		for i:= range max {
			result := vals[0]
			expr := strconv.Itoa(result)
			for j := 1; j < len(vals); j++ {
				div := i / int(math.Pow(float64(len(opers)), float64(j-1)))
				o := opers[div % numOpers]
				result, expr = o.apply(result, vals[j], expr)
			}
			isCorrect := result == target
			if isCorrect {
				fmt.Println(expr, "=",result,isCorrect)
				total += target
				break
			}
		}
	}

	fmt.Println("total is", total)

	_ = input
}

type operator int

const PLUS = operator(0)
const TIMES = operator(1)
const CONCAT = operator(2)

func (o operator) apply(val1 int, val2 int, expr1 string) (int, string) {
	switch(o) {
	case PLUS: return val1 + val2, expr1 + " + " + strconv.Itoa(val2)
	case TIMES: return val1 * val2, expr1 + " * " + strconv.Itoa(val2)
	case CONCAT: 
		result, _ := strconv.Atoi(fmt.Sprintf("%d%d", val1, val2))
		return result, expr1 + " || " + strconv.Itoa(val2)
	default: panic("unknown operator")
	}
}