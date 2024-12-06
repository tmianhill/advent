package utils

import (
	"strconv"
	"strings"
)


func SplitAndParseInts(s string, sep string) []int {
	strparts := strings.Split(s, sep)
	out := make([]int, len(strparts))
	for i,p:=range strparts {
		out[i],_ = strconv.Atoi(p)
	}
	return out
}