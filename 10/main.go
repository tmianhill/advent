package main

import (
	"advent/utils"
	"fmt"
	"strings"
)

func main() {
	input := sampleInput
	input = realInput

	m := loadMap(input)

	starts := m.findStarts()

	totalCount := 0
	totalRating := 0
	for _,s := range starts {
		m.reset()
		m.traverse(s)
		c := m.countReachedSummits()
		r := m.calculateRating()
		fmt.Println("reached", c, "summits from", s, ", rating is", r)
		totalCount += c
		totalRating += r
	}

	fmt.Println("total count is", totalCount, ", total rating is", totalRating)

	_ = input
}

func loadMap(input string) aMap {
	lines := strings.Split(input, "\n")
	m := aMap{}
	for _,l := range lines {
		lineValues := utils.SplitAndParseInts(l, "")
		lineSpots := make([]spot, len(lineValues))
		for i,v := range lineValues {
			lineSpots[i] = spot{height:v, reachedCount: 0}
		}
		m.spots = append(m.spots, lineSpots)
	}
	return m
}

func (m aMap) findStarts() []pos {
	starts := []pos{}
	for i := range m.spots {
		for j,s := range m.spots[i] {
			if s.height == 0 {
				starts = append(starts, pos{i,j})
			}
		}
	} 
	return starts
}

func (m aMap) countReachedSummits() int {
	count := 0
	for i := range m.spots {
		for _,s := range m.spots[i] {
			if s.height == 9 && s.reachedCount > 0 {
				count++
			}
		}
	} 
	return count
}

func (m aMap) calculateRating() int {
	rating := 0
	for i := range m.spots {
		for _,s := range m.spots[i] {
			if s.height == 9 {
				rating+= s.reachedCount
			}
		}
	} 
	return rating
}

func (m aMap) reset() {
	for i := range m.spots {
		for j := range m.spots[i] {
			m.spots[i][j].reachedCount = 0
		}
	}
}

type aMap struct {
	spots [][]spot
}

type spot struct {
	height int
	reachedCount int
}

func (m aMap) traverse(start pos) {

	spotsToConsider := []pos{start}

	for i := 0; i < len(spotsToConsider); i++ {
		p := spotsToConsider[i]
		x := p.x
		y := p.y
		h := m.spots[x][y].height

		for _,p2 := range []pos{{x-1,y},{x+1,y},{x,y-1},{x,y+1}} {
			if p2.x >= 0 && p2.x < len(m.spots) && p2.y >= 0 && p2.y < len(m.spots[0]) {
				newSpot := m.spots[p2.x][p2.y]
				if newSpot.height == h+1 {
					m.spots[p2.x][p2.y].reachedCount++
					if newSpot.height < 9 {
						spotsToConsider = append(spotsToConsider, p2)
					}
				}
			}
		}
	}

}

type pos struct {x int;y int }