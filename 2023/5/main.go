package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	sampleInput := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

	input := sampleInput
	//input = realInput

	lines := strings.Split(input, "\n")
	ssMap := mmap{}
	sfMap := mmap{}
	fwMap := mmap{}
	wlMap := mmap{}
	ltMap := mmap{}
	thMap := mmap{}
	hlMap := mmap{}
	var seeds []valrange
	var currentMap *mmap
	for _,l := range lines {
		switch strings.TrimSpace(l) {
		case "seed-to-soil map:": currentMap = &ssMap
		case "soil-to-fertilizer map:": currentMap = &sfMap
		case "fertilizer-to-water map:": currentMap = &fwMap
		case "water-to-light map:": currentMap = &wlMap
		case "light-to-temperature map:": currentMap = &ltMap
		case "temperature-to-humidity map:": currentMap = &thMap
		case "humidity-to-location map:": currentMap = &hlMap
		case "": continue
		default:
			if currentMap != nil {
//				fmt.Println(l)
				vals := strings.Split(l, " ")
				me := mapEntry{}
				me.destStart, _ = strconv.Atoi(vals[0])
				me.sourceStart, _ = strconv.Atoi(vals[1])
				me.length, _ = strconv.Atoi(vals[2])
				*currentMap = append(*currentMap, me)
			} else if len(l) > 6 && l[:6]=="seeds:" {
				seedVals := strings.Split(strings.TrimSpace(l[6:]), " ")
				fmt.Println("num of seed vals", len(seedVals))
				fmt.Println(seedVals)
				for i:=0;i<len(seedVals);i+=2 {
					start,_ := strconv.Atoi(seedVals[i])
					length,_ := strconv.Atoi(seedVals[i+1])
					seeds = append(seeds, valrange{start: start, end: start + length})
				}
				fmt.Println("loaded seeds:", seeds)
			}
		}
	}

	lowest := 100000000000
	for _,s := range seeds {
		soil := ssMap.mapRange(s)
		fert := sfMap.mapRanges(soil)
		water := fwMap.mapRanges(fert)
		light := wlMap.mapRanges(water)
		temp := ltMap.mapRanges(light)
		hum := thMap.mapRanges(temp)
		loc := hlMap.mapRanges(hum)
		fmt.Println(s,soil,fert,water,light,temp,hum,loc)
		for _,r := range loc {
			if r.start < lowest {
				lowest = r.start
				fmt.Println("new low", lowest)
			}
		}
	}
	fmt.Println(lowest)
	
}

type mmap []mapEntry
type mapEntry struct {
	destStart int
	sourceStart int
	length int
}
func (mm mmap) mapValue(i int) int {
	for _,me := range mm {
		val,ok := me.mapValue(i)
		if ok {
			return val
		}
	}
	return i
}
func (mm mmap) mapRange(r valrange) (out []valrange) {
	bits := map[int]int{}
	bits[r.start] = r.start
	for _,me := range mm {
		o, ok := me.findOverlap(r)
		if ok {
			bits[o.start] = o.start + me.destStart - me.sourceStart
			if _,endexists := bits[o.end]; !endexists { 
				bits[o.end] = o.end
			}
		}
	}
	starts := make([]int, 0, len(bits))
	out = make([]valrange, len(bits))
	for k,_ := range bits {
		starts = append(starts, k)
	}
	slices.Sort(starts)
	for i,s := range starts {
		out[i].start = bits[s]
		if i > 0 {
			out[i-1].end = bits[s]
		}
	}
	out[len(out)-1].end = mm.mapValue(r.end)
	return
}
func (mm mmap) mapRanges(rs []valrange) (out []valrange) {
	for _,r := range rs {
		out = append(out, mm.mapRange(r)...)
	}
	return
}
func (me mapEntry) mapValue(i int) (int,bool) {
	index := i - me.sourceStart
	if index >= 0 && index < me.length {
		return me.destStart + index, true
	} else {
		return 0, false
	}
}
func (me mapEntry) findOverlap(r valrange) (valrange,bool) {
	fmt.Println("input range:", r, "map entry:", me)
	meEnd := me.sourceStart + me.length
	if r.start < me.sourceStart {
		r.start = me.sourceStart
	}
	if r.end > meEnd {
		r.end = meEnd
	}
	if r.start < r.end {
		fmt.Println("overlap range:", r, )
		return r, true
	} else {
		fmt.Println("no overlap")
		return valrange{}, false
	}
}

type valrange struct {
	start int
	end int
}
